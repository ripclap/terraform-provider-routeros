package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
  The menu is empty on the RouterOS 7.23, so no live JSON sample can be shown.
  This menu has no `disabled` or `dynamic` property.
*/

// ResourcePPPL2tpSecret holds per-peer IPsec pre-shared keys for the L2TP server. When
// `/interface/l2tp-server/server` has `use-ipsec` enabled, RouterOS picks the key of the entry whose
// `address` matches the remote peer, falling back to the server-wide `ipsec-secret` when no entry
// matches.
// https://help.mikrotik.com/docs/spaces/ROS/pages/2031631/L2TP
func ResourcePPPL2tpSecret() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ppp/l2tp-secret"),
		MetaId:           PropId(Id),

		"address": {
			Type:     schema.TypeString,
			Required: true,
			Description: "Address or network of the remote peer this pre-shared key applies to. " +
				"`0.0.0.0/0` matches any peer.",
		},
		KeyComment: PropCommentRw,
		"secret": {
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   true,
			Description: "IPsec pre-shared key used for the L2TP connections coming from `address`.",
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
