package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
{
  "address": "",
  "cache-size": "2",
  "cache-state": "0",
  "certificate": "none",
  "enabled": "false",
  "key-size": "128",
  "kme-id": "",
  "peer-sae-id": "",
  "total-keys-received": "0"
}
*/

/*
	`cache-state` and `total-keys-received` are runtime counters, dropped on read.
*/

// ResourceIpIpsecKeyQkd Quantum Key Distribution client used as an IPsec post-quantum PSK source.
// https://help.mikrotik.com/docs/spaces/ROS/pages/341770268/QKD
func ResourceIpIpsecKeyQkd() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/ipsec/key/qkd"),
		MetaId:           PropId(Id),
		MetaSkipFields:   PropSkipFields("cache_state", "total_keys_received"),

		"address": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Address of the KME (Key Management Entity) server in the `address:port` form, e.g. `10.2.3.4:8020`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"cache_size": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Number of keys IPsec will prefetch from the QKD server.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"certificate": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Client certificate used to authenticate against the QKD (KME) server.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyEnabled: {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Enables the QKD key source.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"key_size": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Requested size of each key. " +
				"VERIFY: MikroTik documents the unit as bytes, while the device default is `128`; the unit was not " +
				"verified on hardware.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"kme_id": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Identifier of the KME, used for certificate validation. If it is not specified, the KME " +
				"identity is not validated.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"peer_sae_id": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Identifier of the peer SAE (Secure Application Entity).",
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
