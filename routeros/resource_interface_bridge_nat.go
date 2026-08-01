package routeros

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
	The menu was empty on the RouterOS 7.23:
	`GET /rest/interface/bridge/nat` -> `[]`, so no value sample is available.

	The writable field set is taken from the device itself:

	/console/inspect request=child path="interface,bridge,nat,add"
	  802.3-sap  802.3-type  action  arp-dst-address  arp-dst-mac-address  arp-gratuitous
	  arp-hardware-type  arp-opcode  arp-packet-type  arp-src-address  arp-src-mac-address  chain
	  comment  copy-from  disabled  dst-address  dst-address6  dst-mac-address  dst-port  in-bridge
	  in-bridge-list  in-interface  in-interface-list  ingress-priority  ip-protocol  jump-target
	  limit  log  log-prefix  mac-protocol  new-packet-mark  new-priority  out-bridge  out-bridge-list
	  out-interface  out-interface-list  packet-mark  packet-type  passthrough  place-before
	  src-address  src-address6  src-mac-address  src-port  stp-flags  stp-forward-delay
	  stp-hello-time  stp-max-age  stp-msg-age  stp-port  stp-root-address  stp-root-cost
	  stp-root-priority  stp-sender-address  stp-sender-priority  stp-type  tls-host
	  to-arp-reply-mac-address  to-dst-mac-address  to-src-mac-address  vlan-encap  vlan-id
	  vlan-priority

	The read-only fields reported by the same device are: bytes, dynamic, invalid, packets.

	`802.3-sap` and `802.3-type` are accepted by the console, but a dot in a RouterOS field name is
	used by this provider to address composite (nested) attributes, so those two LLC matchers cannot
	be expressed in the schema. They are omitted here exactly as they are omitted from the existing
	`/interface/bridge/filter` resource.
*/

// ResourceInterfaceBridgeNat Bridge NAT rules.
// https://help.mikrotik.com/docs/spaces/ROS/pages/328068/Bridging+and+Switching#BridgingandSwitching-BridgeFirewall
func ResourceInterfaceBridgeNat() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/bridge/nat"),
		MetaId:           PropId(Id),
		MetaSkipFields:   PropSkipFields("bytes", "packets"),
		MetaSetUnsetFields: PropSetUnsetFields("arp_dst_address", "arp_dst_mac_address", "arp_gratuitous",
			"arp_hardware_type", "arp_opcode", "arp_packet_type", "arp_src_address", "arp_src_mac_address",
			"dst_address", "dst_address6", "dst_mac_address", "dst_port", "in_bridge", "in_bridge_list",
			"in_interface", "in_interface_list", "ingress_priority", "ip_protocol", "limit", "mac_protocol",
			"new_packet_mark", "new_priority", "out_bridge", "out_bridge_list", "out_interface",
			"out_interface_list", "packet_mark", "packet_type", "src_address", "src_address6", "src_mac_address",
			"src_port", "stp_flags", "stp_forward_delay", "stp_hello_time", "stp_max_age", "stp_msg_age",
			"stp_port", "stp_root_address", "stp_root_cost", "stp_root_priority", "stp_sender_address",
			"stp_sender_priority", "stp_type", "tls_host", "to_arp_reply_mac_address", "to_dst_mac_address",
			"to_src_mac_address", "vlan_encap", "vlan_id", "vlan_priority"),

		"action": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Action to take if a packet is matched by the rule.",
			ValidateFunc: validation.StringInSlice([]string{
				"accept", "arp-reply", "drop", "dst-nat", "jump", "log", "mark-packet", "passthrough",
				"redirect", "return", "set-priority", "src-nat",
			}, false),
		},
		"arp_dst_address": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "ARP destination IP address.",
			ValidateFunc: ValidationIpAddress,
		},
		"arp_dst_mac_address": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "ARP destination MAC address.",
			ValidateFunc: ValidationMacAddress,
		},
		"arp_gratuitous": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Matches ARP gratuitous packets.",
		},
		"arp_hardware_type": {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "ARP hardware type. This is normally Ethernet (Type 1).",
			ValidateFunc: Validation64k,
		},
		"arp_opcode": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "ARP opcode (packet type).",
			ValidateDiagFunc: ValidationValInSlice([]string{
				"arp-nak", "drarp-error", "drarp-reply", "drarp-request", "inarp-reply", "inarp-request",
				"reply", "reply-reverse", "request", "request-reverse",
			}, false, true),
		},
		"arp_packet_type": {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "ARP packet type.",
			ValidateFunc: Validation64k,
		},
		"arp_src_address": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "ARP source IP address.",
			ValidateFunc: ValidationIpAddress,
		},
		"arp_src_mac_address": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "ARP source MAC address.",
			ValidateFunc: ValidationMacAddress,
		},
		"chain": {
			Type:     schema.TypeString,
			Required: true,
			Description: "Specifies to which chain rule will be added. If the input does not match the name of an " +
				"already defined chain, a new chain will be created. Built-in chains: `dstnat`, `srcnat`.",
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"dst_address": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "Destination IPv4 address (only if MAC protocol is set to IPv4).",
			ValidateFunc: ValidationIpAddress,
		},
		"dst_address6": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Destination IPv6 address (only if MAC protocol is set to IPv6). A `!` prefix negates the " +
				"match.",
		},
		"dst_mac_address": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "Destination MAC address.",
			ValidateFunc: ValidationMacAddressWithMask,
		},
		"dst_port": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "List of destination port numbers or port number ranges.",
		},
		KeyDynamic: PropDynamicRo,
		"in_bridge": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Bridge interface through which the packet is coming in.",
		},
		"in_bridge_list": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Set of bridge interfaces defined in interface list. Works the same as in-bridge.",
		},
		"in_interface": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Physical interface (i.e., bridge port) through which the packet is coming in.",
		},
		"in_interface_list": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Set of interfaces defined in interface list. Works the same as in-interface.",
		},
		"ingress_priority": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Matches the priority of an ingress packet. Priority may be derived from VLAN, WMM, DSCP, " +
				"or MPLS EXP bit.",
			ValidateFunc: Validation64k,
		},
		KeyInvalid: PropInvalidRo,
		"ip_protocol": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "IP protocol (only if MAC protocol is set to IPv4 or IPv6).",
			ValidateDiagFunc: ValidationValInSlice([]string{
				"dccp", "ddp", "egp", "encap", "etherip", "ggp", "gre", "hmp", "icmp", "icmpv6", "idpr-cmtp",
				"igmp", "ipencap", "ipip", "ipsec-ah", "ipsec-esp", "ipv6-encap", "ipv6-frag", "ipv6-nonxt",
				"ipv6-opts", "ipv6-route", "iso-tp4", "l2tp", "ospf", "pim", "pup", "rdp", "rspf", "rsvp",
				"sctp", "st", "tcp", "udp", "udp-lite", "vmtp", "vrrp", "xns-idp", "xtp",
			}, false, true),
		},
		"jump_target": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Name of the target chain to jump to. Applicable only if action=jump.",
		},
		"limit": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Matches packets up to a limited rate (packet rate or bit rate). A rule using this matcher " +
				"will match until this limit is reached. Parameters are written in the following format: " +
				"rate[/time],burst:mode.",
		},
		"log": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Add a message to the system log.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"log_prefix": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Adds specified text at the beginning of every log message. Applicable if action=log or " +
				"log=yes configured.",
		},
		"mac_protocol": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Ethernet payload type (MAC-level protocol). To match the protocol type of VLAN encapsulated " +
				"frames (0x8100 or 0x88a8), the vlan-encap property should be used. Accepted values are an integer " +
				"0..65535 or one of: `802.2`, `arp`, `capsman`, `dot1x`, `homeplug-av`, `ip`, `ipv6`, `ipx`, `lacp`, " +
				"`length`, `lldp`, `loop-protect`, `macsec`, `mpls-multicast`, `mpls-unicast`, `mvrp`, " +
				"`packing-compr`, `packing-simple`, `pppoe`, `pppoe-discovery`, `rarp`, `romon`, `service-vlan`, " +
				"`vlan`. A `!` prefix negates the match.",
		},
		"new_packet_mark": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Sets a new packet-mark value. Applicable only if action=mark-packet.",
		},
		"new_priority": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Sets a new priority for a packet. This can be the VLAN, WMM or MPLS EXP priority. " +
				"Accepted values are an integer 0..63 or `from-ingress`. Applicable only if action=set-priority.",
		},
		"out_bridge": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Bridge interface through which the packet is going out.",
		},
		"out_bridge_list": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Set of bridge interfaces defined in interface list. Works the same as out-bridge.",
		},
		"out_interface": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Interface the packet is leaving the router.",
		},
		"out_interface_list": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Set of interfaces defined in interface list. Works the same as out-interface.",
		},
		"packet_mark": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Match packets with a certain packet mark.",
		},
		"packet_type": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Matches the MAC frame type.",
			ValidateDiagFunc: ValidationValInSlice([]string{"broadcast", "host", "multicast", "other-host"}, false, true),
		},
		"passthrough": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Whether to let the packet to pass further (like action passthrough) into the " +
				"filter or not (property only valid for some actions).",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyPlaceBefore: PropPlaceBefore,
		"src_address": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "Source IPv4 address (only if MAC protocol is set to IPv4).",
			ValidateFunc: ValidationIpAddress,
		},
		"src_address6": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Source IPv6 address (only if MAC protocol is set to IPv6). A `!` prefix negates the match.",
		},
		"src_mac_address": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "Source MAC address.",
			ValidateFunc: ValidationMacAddress,
		},
		"src_port": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "List of source port numbers or port number ranges.",
		},
		"stp_flags": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The BPDU flags.",
			ValidateDiagFunc: ValidationValInSlice([]string{"topology-change", "topology-change-ack"}, false, true),
		},
		"stp_forward_delay": {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "Forward delay timer.",
			ValidateFunc: Validation64k,
		},
		"stp_hello_time": {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "STP hello packets time.",
			ValidateFunc: Validation64k,
		},
		"stp_max_age": {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "Maximal STP message age.",
			ValidateFunc: Validation64k,
		},
		"stp_msg_age": {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "STP message age.",
			ValidateFunc: Validation64k,
		},
		"stp_port": {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "STP port identifier.",
			ValidateFunc: Validation64k,
		},
		"stp_root_address": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "Root bridge MAC address.",
			ValidateFunc: ValidationMacAddress,
		},
		"stp_root_cost": {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "Root bridge cost.",
			ValidateFunc: Validation64k,
		},
		"stp_root_priority": {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "Root bridge priority.",
			ValidateFunc: Validation64k,
		},
		"stp_sender_address": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "STP message sender MAC address.",
			ValidateFunc: ValidationMacAddress,
		},
		"stp_sender_priority": {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "STP sender priority.",
			ValidateFunc: Validation64k,
		},
		"stp_type": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The BPDU type: config - configuration BPDU OR tcn - topology change notification.",
			ValidateDiagFunc: ValidationValInSlice([]string{"config", "tcn"}, false, true),
		},
		"tls_host": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Allows matching HTTPS traffic based on TLS SNI hostname. Accepts GLOB syntax for wildcard " +
				"matching.",
		},
		"to_arp_reply_mac_address": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "Source MAC address to put in the ARP reply. Applicable only if action=arp-reply.",
			ValidateFunc: ValidationMacAddress,
		},
		"to_dst_mac_address": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "The new destination MAC address. Applicable only if action=dst-nat.",
			ValidateFunc: ValidationMacAddress,
		},
		"to_src_mac_address": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "The new source MAC address. Applicable only if action=src-nat.",
			ValidateFunc: ValidationMacAddress,
		},
		"vlan_encap": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Matches the MAC protocol type encapsulated in the VLAN frame. Accepts the same values as " +
				"`mac_protocol`: an integer 0..65535 or a protocol name. A `!` prefix negates the match.",
		},
		KeyVlanId: PropVlanIdRw("Matches the VLAN identifier field.", false),
		"vlan_priority": {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "Matches the VLAN priority (PCP) field.",
			ValidateFunc: validation.IntBetween(0, 7),
		},
	}
	return &schema.Resource{
		CreateContext: DefaultCreate(resSchema),
		ReadContext:   DefaultRead(resSchema),
		UpdateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			skip := resSchema[MetaSkipFields].Default.(string)
			resSchema[MetaSkipFields].Default = skip + `,"place_before"`
			defer func() {
				resSchema[MetaSkipFields].Default = skip
			}()

			return ResourceUpdate(ctx, resSchema, d, m)
		},
		DeleteContext: DefaultDelete(resSchema),
		Importer: &schema.ResourceImporter{
			StateContext: ImportStateCustomContext(resSchema),
		},

		Schema: resSchema,
	}
}
