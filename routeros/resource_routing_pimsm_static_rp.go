package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
  {
    ".id": "*1",
    "address": "",
    "comment": "",
    "disabled": "false",
    "group": "",
    "inactive": "false",
    "instance": ""
  }
*/

// ResourceRoutingPimsmStaticRp https://help.mikrotik.com/docs/spaces/ROS/pages/61767728/PIM-SM
func ResourceRoutingPimsmStaticRp() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/routing/pimsm/static-rp"),
		MetaId:           PropId(Id),

		"address": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The address of the rendezvous point.",
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"group": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The multicast group range this rendezvous point serves, written as a prefix, for example " +
				"`239.0.0.0/8`.",
		},
		KeyInactive: PropInactiveRo,
		"instance": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of the PIM-SM instance this rendezvous point belongs to.",
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
