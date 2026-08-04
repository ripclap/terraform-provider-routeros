package routeros

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// A configuration that severs management access cannot be rolled back by the
// tool that applied it: the next plan cannot reach the device. These guards
// refuse the few changes that do that unconditionally, and are deliberately
// narrow - each one fires only where there is no reading under which the
// change is safe. Set `lockout_ack = true` on the resource to proceed anyway.

const KeyLockoutAck = "lockout_ack"

// PropLockoutAck is added to the resources the guards cover.
var PropLockoutAck = &schema.Schema{
	Type:     schema.TypeBool,
	Optional: true,
	Default:  false,
	Description: "Proceed with a change this provider would otherwise refuse for severing " +
		"management access. Applies only to the guard on this resource; it does not disable " +
		"validation.",
}

// firewallMatchKeys are the arguments that narrow a firewall rule. A rule with
// none of them set matches every packet in its chain.
var firewallMatchKeys = []string{
	"src-address", "dst-address", "src-address-list", "dst-address-list",
	"src-address-type", "dst-address-type",
	"in-interface", "out-interface", "in-interface-list", "out-interface-list",
	"in-bridge-port", "out-bridge-port", "in-bridge-port-list", "out-bridge-port-list",
	"protocol", "src-port", "dst-port", "port",
	"connection-state", "connection-nat-state", "connection-mark", "connection-type",
	"connection-bytes", "connection-rate", "connection-limit",
	"packet-mark", "routing-mark", "routing-table", "packet-size",
	"icmp-options", "ipv4-options", "tcp-flags", "tcp-mss", "tls-host",
	"layer7-protocol", "content", "dscp", "priority",
	"limit", "time", "random", "nth", "hotspot", "ipsec-policy", "psd",
	"src-mac-address", "fragment", "per-connection-classifier",
}

// Services that are a way in to a device. Disabling the last one, or the one
// currently in use, ends the session with no way back except a console.
var managementServices = map[string]struct{}{
	"api": {}, "api-ssl": {}, "www": {}, "www-ssl": {},
	"ssh": {}, "winbox": {},
}

func lockoutAcknowledged(s map[string]*schema.Schema, d *schema.ResourceData) bool {
	if _, ok := s[KeyLockoutAck]; !ok {
		return false
	}
	ack, _ := d.Get(KeyLockoutAck).(bool)
	return ack
}

// CheckLockout is called from the shared create and update paths with the item
// about to be written. A non-nil error stops the write.
func CheckLockout(path string, item MikrotikItem, s map[string]*schema.Schema, d *schema.ResourceData) error {
	if lockoutAcknowledged(s, d) {
		return nil
	}

	switch path {
	case "/ip/firewall/filter", "/ipv6/firewall/filter":
		return checkFirewallLockout(path, item)
	case "/ip/service":
		return checkServiceLockout(item, d)
	case "/tool/mac-server":
		return checkMacServerLockout(item)
	}

	return nil
}

// checkFirewallLockout refuses an unconditional drop in a chain that carries
// management traffic. A rule that narrows on anything at all is allowed: the
// aim is to catch the rule that matches everything, not to review firewalls.
func checkFirewallLockout(path string, item MikrotikItem) error {
	chain := strings.ToLower(strings.TrimSpace(item["chain"]))
	if chain != "input" && chain != "forward" {
		return nil
	}

	switch strings.ToLower(strings.TrimSpace(item["action"])) {
	case "drop", "reject", "tarpit":
	default:
		return nil
	}

	for _, k := range firewallMatchKeys {
		if strings.TrimSpace(item[k]) != "" {
			return nil
		}
	}

	return fmt.Errorf(
		"refusing to add a %s rule to the %q chain of %s with no match conditions: "+
			"it would drop every packet in that chain, including this session. "+
			"Narrow it with src_address, in_interface, protocol, connection_state or "+
			"any other match, or set %s = true to proceed",
		strings.ToLower(item["action"]), chain, path, KeyLockoutAck)
}

// checkServiceLockout refuses disabling a service that is a way in to the
// device. Only a change to disabled is checked, so editing the port or the
// address list of an already-disabled service is unaffected.
func checkServiceLockout(item MikrotikItem, d *schema.ResourceData) error {
	name := strings.ToLower(strings.TrimSpace(item["name"]))
	if name == "" {
		name = strings.ToLower(strings.TrimSpace(d.Get(KeyName).(string)))
	}
	if _, ok := managementServices[name]; !ok {
		return nil
	}
	if strings.ToLower(strings.TrimSpace(item["disabled"])) != "true" {
		return nil
	}

	return fmt.Errorf(
		"refusing to disable the %q service: it is one of the ways in to this device, "+
			"and this provider reaches it over one of them. Disable it from a session you "+
			"are not using, or set %s = true to proceed",
		name, KeyLockoutAck)
}

// checkMacServerLockout refuses emptying the interface list that MAC-Winbox
// and MAC-telnet answer on, which is the recovery path when IP is unreachable.
func checkMacServerLockout(item MikrotikItem) error {
	v, ok := item["allowed-interface-list"]
	if !ok {
		return nil
	}
	if l := strings.ToLower(strings.TrimSpace(v)); l != "" && l != "none" {
		return nil
	}

	return fmt.Errorf(
		"refusing to clear allowed-interface-list on the MAC server: MAC-Winbox and "+
			"MAC-telnet are the recovery path when the device is unreachable over IP. "+
			"Name an interface list, or set %s = true to proceed", KeyLockoutAck)
}
