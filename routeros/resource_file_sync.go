package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
	The menu was empty on the reference device (RouterOS 7.23.2):
	`GET /rest/file/sync` -> `[]`, so no value sample is available.

	Settable arguments, `/console/inspect request=syntax path="file,sync,add"`:
	  comment  copy-from  disabled  local-path  mode  password  remote-address  remote-path  user

	Readable properties, `/console/inspect request=completion input="/file sync print proplist="`,
	additionally report the read-only `dynamic`, `invalid` and `status` properties.
*/

// ResourceFileSync rsync based synchronisation jobs between the router file system and a remote host.
// https://help.mikrotik.com/docs/spaces/ROS/pages/2555971/File+Systems
func ResourceFileSync() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/file/sync"),
		MetaId:           PropId(Id),
		MetaSkipFields:   PropSkipFields("status"),

		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		KeyDynamic:  PropDynamicRo,
		KeyInvalid:  PropInvalidRo,
		"local_path": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Path on the router that takes part in the synchronisation.",
		},
		"mode": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Direction of the synchronisation:" +
				"\n  * upload - the local path is copied to the remote host," +
				"\n  * download - the remote path is copied to the router.",
			ValidateFunc:     validation.StringInSlice([]string{"download", "upload"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"password": {
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   true,
			Description: "Password of the rsync user on the remote host.",
		},
		"remote_address": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Address of the remote rsync host.",
		},
		"remote_path": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Path on the remote host that takes part in the synchronisation.",
		},
		"user": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "User name used to authenticate on the remote rsync host.",
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
