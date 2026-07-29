package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
  ".id": "*1",
  "ca-identity": "ca",
  "disabled": "false",
  "fingerprint-algorithm": "sha256",
  "name": "ra1",
  "on-smart-card": "false",
  "ra-path": "/scep/ra",
  "ra-transaction-lifetime": "1d",
  "server-url": "http://ca.example.com/scep",
  "status": "idle"
}
*/

// ResourceCertificateScepServerRa A SCEP Registration Authority that forwards approved enrollment requests to the upstream SCEP CA at `server_url`.
// https://help.mikrotik.com/docs/spaces/ROS/pages/2555969/Certificates
func ResourceCertificateScepServerRa() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/certificate/scep-server/ra"),
		MetaId:           PropId(Id),

		"ca_fingerprint": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The fingerprint of the CA certificate received from the upstream SCEP server.",
		},
		"ca_identity": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The CA identity string that is sent to the upstream SCEP server in the `GetCACert` " +
				"request.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"challenge_password": {
			Type:      schema.TypeString,
			Optional:  true,
			Sensitive: true,
			Description: "The challenge password that the enrolling clients have to supply and that is forwarded " +
				"to the upstream SCEP server.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyDisabled: PropDisabledRw,
		"fingerprint_algorithm": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The digest algorithm used to compute the CA certificate fingerprint." +
				"\n  * md5;" +
				"\n  * sha1;" +
				"\n  * sha256.",
			ValidateFunc:     validation.StringInSlice([]string{"md5", "sha1", "sha256"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyName: PropName("Name of the registration authority."),
		"on_smart_card": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Whether the RA private key is kept on a smart card instead of the file system.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"ra_path": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The local HTTP path that the registration authority listens on, for example " +
				"`/scep/ra`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"ra_transaction_lifetime": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "How long an enrollment transaction that is waiting to be granted is kept before it is " +
				"discarded.",
			DiffSuppressFunc: TimeEqual,
		},
		"req_fingerprint": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The fingerprint of the pending enrollment request.",
		},
		"server_url": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The URL of the upstream SCEP server that the requests are forwarded to.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"smart_card_key": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The smart card key that holds the RA private key when `on_smart_card` is enabled.",
		},
		"status": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The current state of the registration authority.",
		},
		"template": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "VERIFY: the name of the template used for the certificates issued through this " +
				"registration authority. The console completion of this argument was empty on the reference device, " +
				"so the menu the value is taken from could not be established.",
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
