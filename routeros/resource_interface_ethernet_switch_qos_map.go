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
	  "name": "default"
	}

	Writable fields reported by the RouterOS 7.23:
	/console/inspect request=child path="interface,ethernet,switch,qos,map,add"
	  comment  copy-from  disabled  name

	The mapping entries themselves live in the `/interface/ethernet/switch/qos/map/vlan` and
	`/interface/ethernet/switch/qos/map/ip` sub-menus.
*/

// ResourceInterfaceEthernetSwitchQosMap Switch chip QoS priority-to-profile mapping tables.
// https://help.mikrotik.com/docs/spaces/ROS/pages/189497483/Quality+of+Service
func ResourceInterfaceEthernetSwitchQosMap() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/ethernet/switch/qos/map"),
		MetaId:           PropId(Id),

		KeyComment:  PropCommentRw,
		KeyDefault:  PropDefaultRo,
		KeyDisabled: PropDisabledRw,
		"hw_id": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "The index the mapping table occupies in the switch chip.",
		},
		KeyHwOffloaded: {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether the mapping table is programmed into the switch chip.",
		},
		KeyInactive: PropInactiveRo,
		KeyName:     PropName("Name of the mapping table."),
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
