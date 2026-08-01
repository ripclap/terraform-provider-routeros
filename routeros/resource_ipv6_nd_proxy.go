package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
  The menu is empty on the RouterOS 7.23, so no live
  JSON sample can be shown. The field set below was taken from the device itself:

    /console/inspect request=syntax path="ipv6,nd,proxy,add"
      -> address, comment, disabled, interface
    /ipv6/nd/proxy/print detail
      -> Flags: X - DISABLED   (no invalid / dynamic flag)
    /ipv6/nd/proxy/print proplist=<name>
      -> "invalid" and "dynamic" do NOT exist in this menu.
*/

// ResourceIPv6NdProxy https://help.mikrotik.com/docs/spaces/ROS/pages/40992815/IPv6+Neighbor+Discovery
func ResourceIPv6NdProxy() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ipv6/nd/proxy"),
		MetaId:           PropId(Id),

		"address": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The IPv6 address the router will assume ownership of and answer Neighbor Solicitations for.",
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		KeyInterface: {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Interface on which the Neighbor Discovery proxy operates.",
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
