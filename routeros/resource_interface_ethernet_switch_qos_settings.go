package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
	{
	  "lossless-buffers": "auto",
	  "lossless-traffic-class": "auto",
	  "mirror-buffers": "auto",
	  "mirror-profile": "default",
	  "multicast-buffers": "auto",
	  "shared-buffers": "auto",
	  "wred-threshold": "medium"
	}
*/

// ResourceInterfaceEthernetSwitchQosSettings Global switch chip QoS settings.
// https://help.mikrotik.com/docs/spaces/ROS/pages/189497483/Quality+of+Service
func ResourceInterfaceEthernetSwitchQosSettings() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/ethernet/switch/qos/settings"),
		MetaId:           PropId(Id),

		"lossless_buffers": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Size of the lossless pool as a percentage of the shared buffer memory (0..100), an amount " +
				"of bytes, or `auto`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"lossless_traffic_class": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The list of lossless traffic classes. Accepts `auto` or a comma separated list of traffic " +
				"classes (0..7).",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"mirror_buffers": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Maximum amount of packet buffers used for mirrored traffic as a percentage of the total " +
				"buffer memory (1..90), an amount of bytes, or `auto`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"mirror_profile": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The QoS profile assigned to mirrored packets.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"multicast_buffers": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Maximum amount of packet buffers used for multicast and broadcast traffic as a percentage " +
				"of the total buffer memory (1..90), an amount of bytes, or `auto`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"shared_buffers": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Maximum amount of packet buffers shared between ports as a percentage of the total buffer " +
				"memory (0..90), an amount of bytes, or `auto`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"wred_threshold": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Relative amount of packets above the shared queue cap at which random drops start to " +
				"occur.",
			ValidateFunc:     validation.StringInSlice([]string{"high", "low", "medium"}, false),
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
