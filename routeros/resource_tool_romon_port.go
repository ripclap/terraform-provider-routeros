package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
	{
	  ".id": "*1",
	  "cost": "100",
	  "default": "true",
	  "disabled": "false",
	  "dynamic": "false",
	  "forbid": "false",
	  "interface": "all",
	  "secrets": ""
	}
*/

// ResourceToolRomonPort Per port RoMON participation rules.
// https://help.mikrotik.com/docs/spaces/ROS/pages/8978569/RoMON
func ResourceToolRomonPort() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/tool/romon/port"),
		MetaId:           PropId(Id),

		KeyComment: PropCommentRw,
		"cost": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Port cost used when RoMON builds the overlay topology (0..4294967295). A lower cost " +
				"makes the port more preferred. The default wildcard entry uses `100`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyDefault:  PropDefaultRo,
		KeyDisabled: PropDisabledRw,
		KeyDynamic:  PropDynamicRo,
		"forbid": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Whether the matched interface is forbidden to participate in the RoMON network.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyInterface: {
			Type:     schema.TypeString,
			Required: true,
			Description: "Name of the interface or of the interface list this entry matches. The special value " +
				"`all` matches every interface.",
		},
		"secrets": {
			Type:      schema.TypeString,
			Optional:  true,
			Sensitive: true,
			Description: "Comma separated list of secrets used for RoMON message authentication on this port. " +
				"Overrides the global `/tool/romon` secrets for the matched interfaces.",
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
