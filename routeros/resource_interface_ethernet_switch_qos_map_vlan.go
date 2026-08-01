package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
 {
	".id": "*1",
	"comment": "",
	"dei-only": "false",
	"disabled": "false",
	"map": "0",
	"pcp": "0",
	"profile": "profile1"
  }
*/

// ResourceInterfaceEthernetSwitchQosMapVlan maps a VLAN priority onto an internal
// traffic class for a switch QoS profile.
// https://help.mikrotik.com/docs/display/ROS/Switch+Chip+Features
func ResourceInterfaceEthernetSwitchQosMapVlan() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/ethernet/switch/qos/map/vlan"),
		MetaId:           PropId(Id),

		KeyComment: PropCommentRw,
		"dei_only": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Match only frames that carry the Drop Eligible Indicator.",
		},
		KeyDisabled: PropDisabledRw,
		"map": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The internal traffic class the matched frames are mapped to.",
		},
		"pcp": {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "The Priority Code Point of the received frame that this entry matches.",
			ValidateFunc: validation.IntBetween(0, 7),
		},
		"profile": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Name of the switch QoS profile this entry belongs to.",
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
