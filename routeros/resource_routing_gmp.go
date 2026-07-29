package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
  {
    ".id": "*1",
    "disabled": "false",
    "exclude": "",
    "groups": "",
    "interfaces": "",
    "sources": ""
  }
*/

// The menu has no `comment` property: RouterOS answers `unknown parameter comment` both on `add` and `set`.

// ResourceRoutingGmp A static group membership entry. The Group Management Protocol menu makes the router join
// the listed multicast groups on the listed interfaces without waiting for a host to report the membership.
// https://help.mikrotik.com/docs/spaces/ROS/pages/128221386/IGMP+Proxy
func ResourceRoutingGmp() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath:   PropResourcePath("/routing/gmp"),
		MetaId:             PropId(Id),
		MetaSetUnsetFields: PropSetUnsetFields("exclude"),

		KeyDisabled: PropDisabledRw,
		"exclude": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
			Description: "Switches the membership to the source filter EXCLUDE mode: the router accepts the traffic " +
				"of the listed groups from every source except the ones listed in `sources`. When disabled, the " +
				"INCLUDE mode is used and only the listed sources are accepted." +
				"\n<em>This property is a flag: the correct value may not be displayed in Winbox, check it in the console.</em>",
		},
		"groups": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "Multicast group addresses the router should join.",
		},
		"interfaces": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "Interfaces on which the membership is created. Both interface names and interface list " +
				"names are accepted.",
		},
		"sources": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "Source addresses of the multicast traffic. The list is interpreted according to the " +
				"`exclude` property: it is the INCLUDE list when `exclude` is disabled and the EXCLUDE list otherwise.",
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
