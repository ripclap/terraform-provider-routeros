package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
	The menu was empty on the RouterOS 7.23:
	`GET /rest/interface/bridge/msti` -> `[]`, so no value sample is available.

	Writable fields reported by the device:
	/console/inspect request=child path="interface,bridge,msti,add"
	  bridge  comment  copy-from  disabled  identifier  priority  vlan-mapping

	Read-only fields reported by the same device: dynamic
*/

// ResourceInterfaceBridgeMsti MSTP Multiple Spanning Tree Instances.
// https://help.mikrotik.com/docs/spaces/ROS/pages/21725254/Spanning+Tree+Protocol
func ResourceInterfaceBridgeMsti() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/bridge/msti"),
		MetaId:           PropId(Id),

		"bridge": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The bridge interface the MST instance belongs to.",
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		KeyDynamic:  PropDynamicRo,
		"identifier": {
			Type:         schema.TypeInt,
			Required:     true,
			Description:  "The MST instance identifier. MSTI0 always exists and cannot be created.",
			ValidateFunc: validation.IntBetween(1, 31),
		},
		"priority": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The bridge priority inside this MST instance, used to determine the regional root bridge " +
				"of the instance. Accepts a decimal or a hexadecimal value in the range 0..65535, in steps of 4096.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"vlan_mapping": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "VLAN IDs that are mapped to this MST instance. Accepts individual VLAN IDs and VLAN ID " +
				"ranges (1..4094) as a comma separated list, e.g. `100-115,120,128-130`. VLAN IDs that are not " +
				"mapped to any MST instance fall under MSTI0.",
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
