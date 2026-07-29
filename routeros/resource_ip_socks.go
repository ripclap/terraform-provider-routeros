package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
  "auth-method": "none",
  "connection-idle-timeout": "2m",
  "enabled": "false",
  "max-connections": "200",
  "port": "1080",
  "version": "4",
  "vrf": "main"
}
*/

// ResourceIpSocks SOCKS proxy server settings.
// https://help.mikrotik.com/docs/spaces/ROS/pages/73826308/SOCKS
func ResourceIpSocks() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/socks"),
		MetaId:           PropId(Id),

		"auth_method": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Authentication method required from the clients." +
				"\n  * none - no authentication." +
				"\n  * password - the client must present the credentials of a `/ip/socks/users` entry (SOCKS5 only).",
			ValidateFunc:     validation.StringInSlice([]string{"none", "password"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"connection_idle_timeout": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Time after which an idle connection is terminated.",
			DiffSuppressFunc: TimeEqual,
		},
		KeyEnabled: {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Enables the SOCKS proxy.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"max_connections": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Maximum number of simultaneous connections.",
			ValidateFunc:     validation.IntBetween(1, 500),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"port": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "TCP port on which the SOCKS server listens for connections.",
			ValidateFunc:     validation.IntBetween(1, 65535),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"version": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "SOCKS protocol version served: `4` or `5`.",
			ValidateFunc:     validation.StringInSlice([]string{"4", "5"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyVrf: PropVrfRw,
	}

	return &schema.Resource{
		CreateContext: DefaultSystemCreate(resSchema),
		ReadContext:   DefaultSystemRead(resSchema),
		UpdateContext: DefaultSystemUpdate(resSchema),
		DeleteContext: DefaultSystemDelete(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resSchema,
	}
}
