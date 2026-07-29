package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
{
  ".id": "*1",
  "default": "true",
  "directory": "/pub",
  "disabled": "true",
  "dynamic": "false",
  "invalid-users": "",
  "name": "pub",
  "read-only": "false",
  "require-encryption": "false",
  "valid-users": ""
}
*/

// ResourceIpSMBShares
// https://help.mikrotik.com/docs/spaces/ROS/pages/117145608/SMB
func ResourceIpSMBShares() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/smb/shares"),
		MetaId:           PropId(Id),

		KeyComment: PropCommentRw,
		KeyDefault: PropDefaultRo,
		"directory": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Directory on the router assigned to the SMB share. If it is left empty, the value of the " +
				"`name` attribute is used, relative to the root folder.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyDisabled: PropDisabledRw,
		KeyDynamic:  PropDynamicRo,
		"invalid_users": {
			Type: schema.TypeSet,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional:         true,
			Description:      "Users (`/ip/smb/users`) that are explicitly denied access to this share.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyName: PropName("Name of the SMB share as it is advertised to the clients."),
		"read_only": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Makes the share read-only for every user.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"require_encryption": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Enforces the use of encryption for all the connections to this share.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"valid_users": {
			Type: schema.TypeSet,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
			Description: "Users (`/ip/smb/users`) that are allowed to access this share. If it is left empty, every " +
				"user can access the share.",
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
