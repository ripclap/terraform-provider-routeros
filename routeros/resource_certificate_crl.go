package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
{
  ".id": "*1",
  "akid": "",
  "cert": "cacert.pem_16",
  "dynamic": "true",
  "fingerprint": "",
  "invalid": "true",
  "revoked": "unknown",
  "signature": "",
  "trust-store": "all",
  "url": "http://www.example.net/crl/root_class_3_ca.crl"
}
*/

// ResourceCertificateCrl A Certificate Revocation List distribution point. Most entries are created dynamically; only a static entry's URL is writable.
// https://help.mikrotik.com/docs/spaces/ROS/pages/2555969/Certificates
func ResourceCertificateCrl() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/certificate/crl"),
		MetaId:           PropId(Id),

		"akid": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The authority key identifier of the CRL issuer.",
		},
		"cert": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The name of the certificate this CRL belongs to.",
		},
		KeyDynamic: PropDynamicRo,
		"expired": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether the downloaded CRL is past its `next_update` time.",
		},
		"fingerprint": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The fingerprint of the downloaded CRL.",
		},
		KeyInvalid: PropInvalidRo,
		"last_update": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The `thisUpdate` time of the downloaded CRL.",
		},
		"next_update": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The `nextUpdate` time of the downloaded CRL.",
		},
		"num": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The number of entries in the downloaded CRL.",
		},
		"revoked": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The revocation status determined from this CRL.",
		},
		"signature": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The signature algorithm of the downloaded CRL.",
		},
		"trust_store": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The trust store the certificate that references this CRL belongs to.",
		},
		"url": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The URL the CRL is downloaded from.",
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
