package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
	The menu was empty on the reference device (RouterOS 7.23.2):
	`GET /rest/tool/traffic-generator/raw-packet-template` -> `[]`, so no value sample is available.

	Settable arguments, `/console/inspect request=syntax
	path="tool,traffic-generator,raw-packet-template,add"`:
	  comment  compute-checksum-from-offset  copy-from  data  data-byte  header  ip-header-offset
	  ipv6-header-offset  name  port  random-byte-offsets-and-masks  random-ranges
	  special-footer  tcp-header-offset  udp-compute-checksum  udp-header-offset

	Readable properties, `/console/inspect request=completion
	input="/tool traffic-generator raw-packet-template print proplist="`, additionally report the
	read-only `dynamic` and `header-length` properties.
*/

// ResourceToolTrafficGeneratorRawPacketTemplate Raw (hand crafted) packet templates for the traffic generator.
// https://help.mikrotik.com/docs/spaces/ROS/pages/128221376/Traffic+Generator
func ResourceToolTrafficGeneratorRawPacketTemplate() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/tool/traffic-generator/raw-packet-template"),
		MetaId:           PropId(Id),

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
			Description:  "How the packet payload following the raw header is filled.",
			ValidateFunc: validation.StringInSlice([]string{"incrementing", "random", "specific-byte", "uninitialized"}, false),
		},
		"data_byte": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Byte value in hexadecimal notation (`00`..`FF`) used to fill the payload when " +
				"`data=specific-byte`.",
		},
		KeyDynamic: PropDynamicRo,
		"header": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The complete packet header as a hexadecimal string.",
		},
		"header_length": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Length of the configured raw header, calculated by RouterOS.",
		},
		"ip_header_offset": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Byte offset of the IPv4 header inside the raw header, so that the generator can update " +
				"the IPv4 checksum and length fields.",
		},
		"ipv6_header_offset": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Byte offset of the IPv6 header inside the raw header.",
		},
		KeyName: PropName("Name of the raw packet template, referenced by `tx_template` of a stream."),
		"port": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Name of the `/tool/traffic-generator/port` entry the packets built from this template " +
				"are sent out of.",
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
			Description: "Ranges the randomized byte values are taken from.",
		},
		"special_footer": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Whether a special footer with a sequence number and a timestamp is appended to every " +
				"packet. It is required for latency and out of order measurements.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"tcp_header_offset": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Byte offset of the TCP header inside the raw header.",
		},
		"udp_compute_checksum": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Whether the generator recalculates the UDP checksum of the generated packets.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"udp_header_offset": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Byte offset of the UDP header inside the raw header.",
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
