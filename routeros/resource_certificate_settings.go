package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
  "builtin-trust-store": "default",
  "crl-download": "false",
  "crl-store": "ram",
  "crl-use": "false",
  "current-defaults": "fetch,mqtt,email,netwatch,container,lora,wiliot,dns,www,reverse-proxy"
}
*/

// The list of services that can use the built-in trust store, as reported by the
// console completion of `/certificate/settings/set builtin-trust-store=` on RouterOS 7.23.
var certificateSettingsTrustStoreServices = []string{
	"ipsec", "wpa-eap", "capsman", "fetch", "sstp", "ovpn", "mqtt", "email", "netwatch", "radius",
	"container", "userman", "lora", "wiliot", "openflow", "tr069", "dot1x", "dns", "www", "api",
	"reverse-proxy", "logging",
}

// ResourceCertificateSettings https://help.mikrotik.com/docs/spaces/ROS/pages/2555969/Certificates
func ResourceCertificateSettings() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/certificate/settings"),
		MetaId:           PropId(Id),

		"builtin_trust_store": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Specifies which services may use the built-in trust store authorities for certificate " +
				"verification. Accepts a comma-separated list of service names, or one of the shortcuts `all`, " +
				"`default` (the set reported by `current_defaults`) and `untrusted` (no service).",
			ValidateDiagFunc: ValidationMultiValInSlice(
				append([]string{"all", "default", "untrusted"}, certificateSettingsTrustStoreServices...), false, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"crl_download": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Whether to automatically download and update the CRLs.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"crl_store": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Where the downloaded CRL information is stored." +
				"\n  * ram - keep the CRLs in RAM only;" +
				"\n  * system - store the CRLs on the system disk.",
			ValidateFunc:     validation.StringInSlice([]string{"ram", "system"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"crl_use": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Whether every certificate in a chain is verified against its CRL.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"current_defaults": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "A comma-separated list of the services that the `default` value of " +
				"`builtin_trust_store` expands to.",
			ValidateDiagFunc: ValidationMultiValInSlice(certificateSettingsTrustStoreServices, false, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
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
