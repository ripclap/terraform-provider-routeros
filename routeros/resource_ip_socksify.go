package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
	Menu is empty on the reference device (RouterOS 7.23.2), so no GET sample is available.
*/

// ResourceIpSocksify Redirects intercepted TCP traffic through an upstream SOCKS5 proxy.
// https://help.mikrotik.com/docs/spaces/ROS/pages/343244851/Socksify
func ResourceIpSocksify() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/socksify"),
		MetaId:           PropId(Id),

		KeyComment: PropCommentRw,
		"connection_timeout": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Number of seconds to wait for a response from the SOCKS5 server or from the destination " +
				"before the connection is dropped. `0` disables the timeout.",
			ValidateFunc:     validation.IntBetween(0, 3000),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyDisabled: PropDisabledRw,
		KeyName:     PropName("Name of the Socksify entry."),
		"port": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Local TCP port the Socksify service listens on for the redirected traffic.",
			ValidateFunc:     validation.IntBetween(1, 65535),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"socks5_password": {
			Type:             schema.TypeString,
			Optional:         true,
			Sensitive:        true,
			Description:      "Password used to authenticate against the upstream SOCKS5 server.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"socks5_port": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "TCP port of the upstream SOCKS5 server.",
			ValidateFunc:     validation.IntBetween(1, 65535),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"socks5_server": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "IPv4 address of the upstream SOCKS5 server.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"socks5_user": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "User name used to authenticate against the upstream SOCKS5 server.",
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
