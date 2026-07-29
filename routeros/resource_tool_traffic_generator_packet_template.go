package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
	Header field properties accept a comma separated list of up to 16 values; when the same header
	appears in the stack more than once the n-th list element is used for the n-th header.
*/

// ResourceToolTrafficGeneratorPacketTemplate Structured packet templates for the traffic generator.
// https://help.mikrotik.com/docs/spaces/ROS/pages/128221376/Traffic+Generator
func ResourceToolTrafficGeneratorPacketTemplate() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/tool/traffic-generator/packet-template"),
		MetaId:           PropId(Id),
		MetaSkipFields: PropSkipFields(
			"assumed_interface", "assumed_ip_dscp", "assumed_ip_dst", "assumed_ip_frag_off", "assumed_ip_id",
			"assumed_ip_protocol", "assumed_ip_src", "assumed_ip_ttl", "assumed_ipv6_dst", "assumed_ipv6_flow_label",
			"assumed_ipv6_hop_limit", "assumed_ipv6_next_header", "assumed_ipv6_src", "assumed_ipv6_traffic_class",
			"assumed_mac_dst", "assumed_mac_protocol", "assumed_mac_src", "assumed_port", "assumed_raw_header",
			"assumed_tcp_ack", "assumed_tcp_data_offset", "assumed_tcp_dst_port", "assumed_tcp_flags",
			"assumed_tcp_src_port", "assumed_tcp_syn", "assumed_tcp_urgent_pointer", "assumed_tcp_window_size",
			"assumed_udp_checksum", "assumed_udp_dst_port", "assumed_udp_src_port", "assumed_vlan_id",
			"assumed_vlan_priority", "assumed_vlan_protocol"),

		KeyComment: PropCommentRw,
		"compute_checksum_from_offset": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Byte offset at which the generator starts computing the two byte checksum that is placed " +
				"into the packet, or `no-checksum` to leave the checksum untouched.",
		},
		"data": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "How the packet payload following the headers is filled.",
			ValidateFunc: validation.StringInSlice([]string{"incrementing", "random", "specific-byte", "uninitialized"}, false),
		},
		"data_byte": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Byte value in hexadecimal notation (`00`..`FF`) used to fill the payload when " +
				"`data=specific-byte`.",
		},
		"header_stack": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Comma separated sequence of headers the generated packet consists of, for example " +
				"`mac,ip,udp`.",
			ValidateDiagFunc: ValidationMultiValInSlice([]string{"ip", "ipv6", "mac", "raw", "tcp", "udp", "vlan"},
				false, false),
		},
		KeyInterface: {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Interface the packets built from this template are sent out of. Mutually exclusive " +
				"with `port`.",
		},
		"ip_dscp": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Comma separated list of DS field values placed into the IP headers.",
		},
		"ip_dst": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Comma separated list of destination addresses (`address` or `address/netmask`) for the IP headers.",
		},
		"ip_frag_off": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Comma separated list of fragment offset values placed into the IP headers.",
		},
		"ip_gateway": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Next hop address the generator resolves to obtain the destination MAC address when " +
				"`mac-dst` is not given.",
		},
		"ip_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Comma separated list of identification field values placed into the IP headers.",
		},
		"ip_protocol": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Comma separated list of protocol values placed into the IP headers, specified by " +
				"protocol name (`tcp`, `udp`, `icmp`, `gre`, …) or by number.",
		},
		"ip_src": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Comma separated list of source addresses (`address` or `address/netmask`) for the IP headers.",
		},
		"ip_ttl": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Comma separated list of time to live values placed into the IP headers.",
		},
		"ipv6_dst": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Comma separated list of destination addresses for the IPv6 headers.",
		},
		"ipv6_flow_label": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Comma separated list of flow label values placed into the IPv6 headers.",
		},
		"ipv6_gateway": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Next hop IPv6 address the generator resolves to obtain the destination MAC address when " +
				"`mac-dst` is not given.",
		},
		"ipv6_hop_limit": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Comma separated list of hop limit values placed into the IPv6 headers.",
		},
		"ipv6_next_header": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Comma separated list of next header values placed into the IPv6 headers, specified by " +
				"protocol name (`tcp`, `udp`, `icmpv6`, …) or by number.",
		},
		"ipv6_src": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Comma separated list of source addresses for the IPv6 headers.",
		},
		"ipv6_traffic_class": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Comma separated list of traffic class values placed into the IPv6 headers.",
		},
		"mac_dst": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Comma separated list of destination MAC addresses (`MAC` or `MAC/MASK`) for the MAC headers.",
		},
		"mac_protocol": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Comma separated list of ethertypes placed into the MAC headers, specified by name " +
				"(`ip`, `ipv6`, `arp`, `vlan`, `service-vlan`, `mpls-unicast`, …) or by number.",
		},
		"mac_src": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Comma separated list of source MAC addresses (`MAC` or `MAC/MASK`) for the MAC headers.",
		},
		KeyName: PropName("Name of the packet template, referenced by `tx_template` of a stream."),
		"port": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Name of the `/tool/traffic-generator/port` entry the packets built from this template " +
				"are sent out of. Mutually exclusive with `interface`.",
		},
		"random_byte_offsets_and_masks": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "List of `offset:mask` pairs describing which packet bytes are randomized and which bits " +
				"of them may change.",
		},
		"random_ranges": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Ranges the randomized header field values are taken from.",
		},
		"raw_header": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Content of a `raw` header as a hexadecimal string.",
		},
		"special_footer": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Whether a special footer with a sequence number and a timestamp is appended to every " +
				"packet. It is required for latency and out of order measurements.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"tcp_ack": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Comma separated list of acknowledgment number values placed into the TCP headers.",
		},
		"tcp_data_offset": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Comma separated list of data offset values placed into the TCP headers.",
		},
		"tcp_dst_port": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Comma separated list of destination ports (`port` or `port/mask`) for the TCP headers.",
		},
		"tcp_flags": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Comma separated list of TCP flags set in the generated TCP headers.",
			ValidateDiagFunc: ValidationMultiValInSlice([]string{"fin", "syn", "rst", "psh", "ack", "urg", "ece",
				"cwr", "ns", "res0", "res1", "res2"}, false, false),
		},
		"tcp_src_port": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Comma separated list of source ports (`port` or `port/mask`) for the TCP headers.",
		},
		"tcp_syn": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Comma separated list of sequence number values placed into the TCP headers.",
		},
		"tcp_urgent_pointer": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Comma separated list of urgent pointer values placed into the TCP headers.",
		},
		"tcp_window_size": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Comma separated list of window size values placed into the TCP headers.",
		},
		"udp_checksum": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Comma separated list of checksum values placed into the UDP headers, or `compute` to let " +
				"the generator calculate them.",
		},
		"udp_dst_port": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Comma separated list of destination ports (`port` or `port/mask`) for the UDP headers.",
		},
		"udp_src_port": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Comma separated list of source ports (`port` or `port/mask`) for the UDP headers.",
		},
		"vlan_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Comma separated list of VLAN identifiers placed into the VLAN headers.",
		},
		"vlan_priority": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Comma separated list of PCP values placed into the VLAN headers.",
		},
		"vlan_protocol": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Comma separated list of ethertypes carried by the VLAN headers, specified by name " +
				"(`ip`, `ipv6`, `vlan`, `service-vlan`, …) or by number.",
		},
	}

	return &schema.Resource{
		CreateContext: DefaultCreate(resSchema),
		ReadContext:   DefaultRead(resSchema),
		UpdateContext: DefaultUpdate(resSchema),
		DeleteContext: DefaultDelete(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: ImportStateCustomContext(resSchema),
		},

		Schema: resSchema,
	}
}
