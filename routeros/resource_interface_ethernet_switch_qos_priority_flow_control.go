package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
	{
	  ".id": "*0",
	  "default": "true",
	  "disabled": "false",
	  "hw-id": "0",
	  "hw-offloaded": "true",
	  "inactive": "false",
	  "name": "disabled",
	  "rx": "false",
	  "tx": "false"
	}

	Writable fields reported by the RouterOS 7.23:
	/console/inspect request=child path="interface,ethernet,switch,qos,priority-flow-control,add"
	  comment  copy-from  disabled  name  pause-threshold  resume-threshold  rx  traffic-class  tx
*/

// ResourceInterfaceEthernetSwitchQosPriorityFlowControl Switch chip Priority-based Flow Control (PFC) profiles.
// https://help.mikrotik.com/docs/spaces/ROS/pages/189497483/Quality+of+Service
func ResourceInterfaceEthernetSwitchQosPriorityFlowControl() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/ethernet/switch/qos/priority-flow-control"),
		MetaId:           PropId(Id),

		KeyComment:  PropCommentRw,
		KeyDefault:  PropDefaultRo,
		KeyDisabled: PropDisabledRw,
		"hw_id": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "The index the PFC profile occupies in the switch chip.",
		},
		KeyHwOffloaded: {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether the PFC profile is programmed into the switch chip.",
		},
		KeyInactive: PropInactiveRo,
		KeyName:     PropName("Name of the PFC profile."),
		"pause_threshold": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Transmit a pause frame (XOFF) when the amount of enqueued packets reaches this threshold. " +
				"Accepts a percentage, an amount of bytes, or `auto`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"resume_threshold": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Transmit a resume frame (XON) when the amount of enqueued packets drops to this threshold. " +
				"Accepts a percentage, an amount of bytes, or `auto`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"rx": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Enables receiving of PFC frames.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"traffic_class": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The list of PFC-enabled traffic classes. Accepts a comma separated list of traffic " +
				"classes (0..7).",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"tx": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Enables transmission of PFC frames.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
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
