package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
	The menu was empty on the reference device (RouterOS 7.23.2):
	`GET /rest/tool/traffic-generator/stream` -> `[]`, so no value sample is available.

	Settable arguments, `/console/inspect request=syntax
	path="tool,traffic-generator,stream,add"`:
	  copy-from  cpu-core  disabled  id  mbps  name  packet-count  packet-size  port  pps
	  tx-template

	Readable properties, `/console/inspect request=completion
	input="/tool traffic-generator stream print proplist="`, additionally report the read-only
	`default-port` and `invalid` properties.

	The RouterOS property is literally called `id`; `id` is reserved by the Terraform SDK, so the
	attribute is exposed as `stream_id` and mapped back onto the wire name by the transform set.
*/

// ResourceToolTrafficGeneratorStream Traffic generator streams.
// https://help.mikrotik.com/docs/spaces/ROS/pages/128221376/Traffic+Generator
func ResourceToolTrafficGeneratorStream() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/tool/traffic-generator/stream"),
		MetaId:           PropId(Id),
		MetaTransformSet: PropTransformSet("stream_id: id"),

		"cpu_core": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "CPU core, or range of CPU cores, that generates this stream. The accepted format is " +
				"`Start[-End]` with both bounds in the range `0..255`, for example `0` or `0-3`.",
		},
		"default_port": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Port RouterOS uses for this stream when `port` is not set explicitly.",
		},
		KeyDisabled: PropDisabledRw,
		KeyInvalid:  PropInvalidRo,
		"mbps": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Target transmit rate of the stream in megabits per second (0..4294967295). Mutually " +
				"exclusive with `pps`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyName: PropName("Name of the stream."),
		"packet_count": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "How many packets are generated before the stream stops, or `unlimited` to keep generating " +
				"packets until the test is stopped." +
				"\n> The property is not purely numeric (`/console/inspect request=completion " +
				"path=tool,traffic-generator,stream input=\"add packet-count=\"` offers `unlimited`), so it is " +
				"declared as a string.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"packet_size": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Size of the generated packets in bytes. A range such as `64-1500` makes the generator " +
				"vary the packet size.",
		},
		"port": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Name of the `/tool/traffic-generator/port` entry this stream is transmitted from.",
		},
		"pps": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Target transmit rate of the stream in packets per second (0..4294967295). Mutually " +
				"exclusive with `mbps`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"stream_id": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Stream number, sent on the wire as the RouterOS `id` property. It is written into the " +
				"generated packets so that the receiving side can attribute statistics to this stream.",
			ValidateFunc:     validation.IntBetween(0, 255),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"tx_template": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Name of the `/tool/traffic-generator/packet-template` or " +
				"`/tool/traffic-generator/raw-packet-template` entry that describes the packets of this stream.",
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
