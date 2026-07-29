package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
{
  ".id": "*1",
  "default": "true",
  "disabled": "false",
  "dynamic": "false",
  "name": "guest",
  "password": "",
  "read-only": "true"
}
*/

// ResourceIpSMBUsers
// https://help.mikrotik.com/docs/spaces/ROS/pages/117145608/SMB
func ResourceIpSMBUsers() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/smb/users"),
		MetaId:           PropId(Id),

		KeyComment:  PropCommentRw,
		KeyDefault:  PropDefaultRo,
		KeyDisabled: PropDisabledRw,
		KeyDynamic:  PropDynamicRo,
		KeyName:     PropName("Login name used by the SMB client."),
		"password": {
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   true,
			Description: "Password of the SMB user.",
		},
		"read_only": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Denies write access to the shares for this user.",
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
