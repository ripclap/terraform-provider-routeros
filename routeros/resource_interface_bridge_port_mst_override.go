package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
	The menu was empty on the RouterOS 7.23:
	`GET /rest/interface/bridge/port/mst-override` -> `[]`, so no value sample is available.

	Writable fields reported by the device:
	/console/inspect request=child path="interface,bridge,port,mst-override,add"
	  comment  copy-from  disabled  identifier  interface  internal-path-cost  priority

	Read-only fields reported by the same device: debug-info, dynamic
*/

// ResourceInterfaceBridgePortMstOverride Per-MSTI overrides of the bridge port path cost and priority.
// https://help.mikrotik.com/docs/spaces/ROS/pages/21725254/Spanning+Tree+Protocol
func ResourceInterfaceBridgePortMstOverride() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/bridge/port/mst-override"),
		MetaId:           PropId(Id),
		MetaSkipFields:   PropSkipFields("debug_info"),

		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		KeyDynamic:  PropDynamicRo,
		"identifier": {
			Type:         schema.TypeInt,
			Required:     true,
			Description:  "The MST instance identifier this override applies to.",
			ValidateFunc: validation.IntBetween(1, 31),
		},
		KeyInterface: {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The bridge port the override applies to.",
		},
		"internal_path_cost": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Path cost to the interface inside the MST instance, used on VLANs that are facing towards " +
				"the regional root bridge.",
			ValidateFunc:     validation.IntBetween(1, 200000000),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"priority": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The priority of the interface inside the MST instance, used on VLANs that are facing away " +
				"from the regional root bridge to manipulate path selection. Accepts a decimal or a hexadecimal " +
				"value in the range 0..240, in steps of 16.",
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
