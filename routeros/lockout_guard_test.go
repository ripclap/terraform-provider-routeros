package routeros

import "testing"

func TestCheckFirewallLockout(t *testing.T) {
	for _, tc := range []struct {
		name    string
		path    string
		item    MikrotikItem
		refused bool
	}{
		{"unconditional input drop", "/ip/firewall/filter",
			MikrotikItem{"chain": "input", "action": "drop"}, true},
		{"unconditional input reject", "/ip/firewall/filter",
			MikrotikItem{"chain": "input", "action": "reject"}, true},
		{"unconditional forward drop", "/ip/firewall/filter",
			MikrotikItem{"chain": "forward", "action": "drop"}, true},
		{"unconditional v6 input drop", "/ipv6/firewall/filter",
			MikrotikItem{"chain": "input", "action": "drop"}, true},
		{"case and space tolerant", "/ip/firewall/filter",
			MikrotikItem{"chain": " Input ", "action": " DROP "}, true},

		{"narrowed by source", "/ip/firewall/filter",
			MikrotikItem{"chain": "input", "action": "drop", "src-address": "192.0.2.0/24"}, false},
		{"narrowed by interface", "/ip/firewall/filter",
			MikrotikItem{"chain": "input", "action": "drop", "in-interface": "ether1"}, false},
		{"narrowed by connection state", "/ip/firewall/filter",
			MikrotikItem{"chain": "input", "action": "drop", "connection-state": "invalid"}, false},
		{"accept is never refused", "/ip/firewall/filter",
			MikrotikItem{"chain": "input", "action": "accept"}, false},
		{"other chains are not guarded", "/ip/firewall/filter",
			MikrotikItem{"chain": "output", "action": "drop"}, false},
		{"empty match value does not count", "/ip/firewall/filter",
			MikrotikItem{"chain": "input", "action": "drop", "src-address": "  "}, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			err := checkFirewallLockout(tc.path, tc.item)
			if tc.refused && err == nil {
				t.Fatal("expected the rule to be refused")
			}
			if !tc.refused && err != nil {
				t.Fatalf("expected the rule to be allowed, got: %v", err)
			}
		})
	}
}

func TestCheckMacServerLockout(t *testing.T) {
	for _, tc := range []struct {
		name    string
		item    MikrotikItem
		refused bool
	}{
		{"cleared", MikrotikItem{"allowed-interface-list": ""}, true},
		{"set to none", MikrotikItem{"allowed-interface-list": "none"}, true},
		{"a list is named", MikrotikItem{"allowed-interface-list": "mgmt"}, false},
		{"attribute absent", MikrotikItem{"enabled": "true"}, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			err := checkMacServerLockout(tc.item)
			if tc.refused != (err != nil) {
				t.Fatalf("refused=%v, err=%v", tc.refused, err)
			}
		})
	}
}

// Every key the guard treats as narrowing must be one the firewall resources
// actually accept, or the guard would refuse a rule the user cannot express.
func TestFirewallMatchKeysExistInSchema(t *testing.T) {
	schemas := []map[string]interface{}{}
	for _, r := range []string{"routeros_ip_firewall_filter", "routeros_ipv6_firewall_filter"} {
		res, ok := NewProvider().ResourcesMap[r]
		if !ok {
			t.Fatalf("%s is not registered", r)
		}
		m := map[string]interface{}{}
		for k := range res.Schema {
			m[k] = nil
		}
		schemas = append(schemas, m)
	}

	var orphan []string
	for _, k := range firewallMatchKeys {
		snake := KebabToSnake(k)
		found := false
		for _, m := range schemas {
			if _, ok := m[snake]; ok {
				found = true
				break
			}
		}
		if !found {
			orphan = append(orphan, k)
		}
	}
	if len(orphan) > 0 {
		t.Errorf("match keys absent from both firewall schemas: %v", orphan)
	}
}
