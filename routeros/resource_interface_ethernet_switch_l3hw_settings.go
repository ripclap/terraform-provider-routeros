package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
	{
	  "autorestart": "true",
	  "fasttrack-hw": "true",
	  "hw-supports-fasttrack": "true",
	  "icmp-reply-on-error": "true",
	  "ipv6-hw": "false"
	}
*/

// ResourceInterfaceEthernetSwitchL3HwSettings Global settings of the L3 hardware offloading engine.
// https://help.mikrotik.com/docs/spaces/ROS/pages/62390319/L3+Hardware+Offloading
func ResourceInterfaceEthernetSwitchL3HwSettings() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/ethernet/switch/l3hw-settings"),
		MetaId:           PropId(Id),

		"autorestart": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Automatically restarts the l3hw driver in case of an error.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"fasttrack_hw": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Enables or disables FastTrack hardware offloading.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"hw_supports_fasttrack": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Indicates if the hardware (switch chip) supports FastTrack hardware offloading.",
		},
		"icmp_reply_on_error": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Since the hardware cannot send ICMP messages, the packet must be redirected to the CPU. " +
				"Disabling this option stops the router from generating ICMP errors for hardware offloaded traffic.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"ipv6_hw": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Enables or disables IPv6 hardware offloading.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
	}

	return &schema.Resource{
		CreateContext: DefaultSystemCreate(resSchema),
		ReadContext:   DefaultSystemRead(resSchema),
		UpdateContext: DefaultSystemUpdate(resSchema),
		DeleteContext: DefaultSystemDelete(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resSchema,
	}
}
