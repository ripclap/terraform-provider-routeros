package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
	The menu was empty on the RouterOS 7.23, so no value sample is available.
	The menu has no `set` command, so every attribute is ForceNew. The RouterOS `id` property is
	reserved by the Terraform SDK, so it is exposed as `psk_id` and remapped by the transform set.
*/

// ResourceIpIpsecKeyPsk Static post-quantum pre-shared keys (PPK, RFC 9867).
// https://help.mikrotik.com/docs/spaces/ROS/pages/11993097/IPsec
func ResourceIpIpsecKeyPsk() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/ipsec/key/psk"),
		MetaId:           PropId(Id),
		MetaTransformSet: PropTransformSet("psk_id: id"),

		"key": {
			Type:        schema.TypeString,
			Optional:    true,
			ForceNew:    true,
			Sensitive:   true,
			Description: "The pre-shared key material itself.",
		},
		"peer": {
			Type:        schema.TypeString,
			Optional:    true,
			ForceNew:    true,
			Description: "Name of the IPsec peer (`/ip/ipsec/peer`) this key belongs to.",
		},
		"psk_id": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
			Description: "Identity of the pre-shared key, sent on the wire as the RouterOS `id` property. It is " +
				"exchanged with the remote side to select the matching key.",
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
