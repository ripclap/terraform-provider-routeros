package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
	{
	  ".id": "*1",
	  "ciphers": "aes-gcm-128",
	  "default": "true",
	  "default-name": "default",
	  "name": "default",
	  "server-priority": "10"
	}

	Writable fields reported by the reference device (RouterOS 7.23.2):
	/console/inspect request=child path="interface,macsec,profile,add"
	  ciphers  copy-from  name  server-priority
*/

// ResourceInterfaceMacsecProfile MACsec profiles.
// https://help.mikrotik.com/docs/spaces/ROS/pages/201523202/MACsec
func ResourceInterfaceMacsecProfile() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/macsec/profile"),
		MetaId:           PropId(Id),

		"ciphers": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The cipher suite used to encrypt the MACsec traffic.",
			ValidateFunc:     validation.StringInSlice([]string{"aes-gcm-128", "aes-gcm-xpn-128"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyDefault:     PropDefaultRo,
		KeyDefaultName: PropDefaultNameRo("The default name of the profile."),
		KeyName:        PropName("Name of the MACsec profile."),
		"server_priority": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Key server priority. A lower value means a higher priority; the MAC address is used as " +
				"a tiebreaker when both sides use the same priority.",
			ValidateFunc:     validation.IntBetween(0, 255),
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
