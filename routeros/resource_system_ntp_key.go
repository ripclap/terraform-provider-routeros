package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
{
  ".id": "*1",
  "comment": "",
  "key-id": "1",
  "key-val": "secret"
}
*/

// ResourceSystemNtpKey https://help.mikrotik.com/docs/spaces/ROS/pages/40992869/NTP
func ResourceSystemNtpKey() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/system/ntp/key"),
		MetaId:           PropId(Id),

		KeyComment: PropCommentRw,
		"key_id": {
			Type:     schema.TypeInt,
			Required: true,
			Description: "The key identifier - an integer identifying the cryptographic key used to generate the " +
				"message authentication code.",
		},
		"key_val": {
			Type:        schema.TypeString,
			Required:    true,
			Sensitive:   true,
			Description: "The value of the symmetric key shared with the NTP peer.",
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
