package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
 {
	".id": "*1",
	"comment": "",
	"disabled": "false",
	"dscp": "0",
	"map": "0",
	"profile": "profile1"
  }
*/

// ResourceInterfaceEthernetSwitchQosMapIp maps a DSCP value onto an internal
// traffic class for a switch QoS profile.
// https://help.mikrotik.com/docs/display/ROS/Switch+Chip+Features
func ResourceInterfaceEthernetSwitchQosMapIp() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/ethernet/switch/qos/map/ip"),
		MetaId:           PropId(Id),

		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"dscp": {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "The DSCP value of the received packet that this entry matches.",
			ValidateFunc: validation.IntBetween(0, 63),
		},
		"map": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The internal traffic class the matched packets are mapped to.",
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
