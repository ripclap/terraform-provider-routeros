package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
  ".about": "bind error: cannot bind to port 443: Address in use (12)",
  ".id": "*3",
  "certificate": "apps-*.example.com",
  "disabled": "false",
  "dynamic": "true",
  "ip-address": "192.0.2.108",
  "port": "8000",
  "sni": "app.example.com"
}
*/

/*
	Sample above is real `GET /rest/ip/reverse-proxy` output with the host names anonymized.
*/

// ResourceIpReverseProxy HTTPS reverse proxy: forwards traffic by SNI to a backend instead of a dst-nat rule.
// https://help.mikrotik.com/docs/spaces/ROS/pages/377225232/Reverse+Proxy
func ResourceIpReverseProxy() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/reverse-proxy"),
		MetaId:           PropId(Id),

		"certificate": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Certificate presented for this entry. When it is `none`, the certificate of the " +
				"`www-ssl` service is used.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		KeyDynamic:  PropDynamicRo,
		"ip_address": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "IP address of the backend server that receives the proxied connections.",
			ValidateFunc:     validation.IsIPAddress,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"port": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "TCP port of the backend server that receives the proxied connections.",
			ValidateFunc:     Validation64k,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"sni": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Server Name Indication the incoming TLS connection must carry for this entry to be used. " +
				"The reverse proxy listens on TCP 443, so the `www-ssl` service has to be disabled or moved to " +
				"another port.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyVrf: PropVrfRw,
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
