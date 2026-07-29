package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
	{
	  ".id": "*1",
	  "comment": "Use this for permanently disconnected ports",
	  "default": "true",
	  "disabled": "false",
	  "hw-id": "1",
	  "hw-offloaded": "true",
	  "inactive": "false",
	  "name": "offline",
	  "queue-buffers": "auto"
	}

	Writable fields reported by the reference device (RouterOS 7.23.2):
	/console/inspect request=child path="interface,ethernet,switch,qos,tx-manager,add"
	  comment  copy-from  disabled  name  queue-buffers

	The per-queue scheduling settings live in the `/interface/ethernet/switch/qos/tx-manager/queue`
	sub-menu.
*/

// ResourceInterfaceEthernetSwitchQosTxManager Switch chip QoS transmission managers.
// https://help.mikrotik.com/docs/spaces/ROS/pages/189497483/Quality+of+Service
func ResourceInterfaceEthernetSwitchQosTxManager() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/ethernet/switch/qos/tx-manager"),
		MetaId:           PropId(Id),

		KeyComment:  PropCommentRw,
		KeyDefault:  PropDefaultRo,
		KeyDisabled: PropDisabledRw,
		"hw_id": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "The index the transmission manager occupies in the switch chip.",
		},
		KeyHwOffloaded: {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether the transmission manager is programmed into the switch chip.",
		},
		KeyInactive: PropInactiveRo,
		KeyName:     PropName("Name of the transmission manager."),
		"queue_buffers": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Total amount of hardware Tx buffers allocated to all ports linked to this transmission " +
				"manager. Accepts a percentage, an amount of bytes, or `auto`.",
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
