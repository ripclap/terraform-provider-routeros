package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
  {
    ".id": "*1",
    "address": "",
    "disabled": "false",
    "instance": ""
  }
*/

// ResourceRoutingRipStaticNeighbor https://help.mikrotik.com/docs/spaces/ROS/pages/328211/RIP
func ResourceRoutingRipStaticNeighbor() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/routing/rip/static-neighbor"),
		MetaId:           PropId(Id),

		"address": {
			Type:     schema.TypeString,
			Required: true,
			Description: "The address of the RIP neighbor. For RIPng the link-local address together with the " +
				"interface it is reachable through is used, for example `fe80::1%ether1`.",
		},
		KeyDisabled: PropDisabledRw,
		"instance": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of the RIP instance this neighbor belongs to.",
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
