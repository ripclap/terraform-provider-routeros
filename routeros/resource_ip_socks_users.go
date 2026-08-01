package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
	Menu is empty on the RouterOS 7.23, so no GET sample is available.
*/

// ResourceIpSocksUsers SOCKS proxy users, used when `/ip/socks` runs with auth-method=password.
// https://help.mikrotik.com/docs/spaces/ROS/pages/73826308/SOCKS
func ResourceIpSocksUsers() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/socks/users"),
		MetaId:           PropId(Id),

		KeyDisabled: PropDisabledRw,
		KeyName:     PropName("Name of the SOCKS user."),
		"only_one": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Allows only one simultaneous connection per user.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"password": {
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   true,
			Description: "Password used by the user to access the SOCKS server.",
		},
		"rate_limit": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Rate limit applied to this user, in bits per second.",
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
