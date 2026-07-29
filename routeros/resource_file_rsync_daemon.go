package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
	{
	  "enabled": "false"
	}

	Reference device: RouterOS 7.23.2, `GET /rest/file/rsync-daemon`.
*/

// ResourceFileRsyncDaemon The built in rsync daemon that serves the router file system.
// https://help.mikrotik.com/docs/spaces/ROS/pages/2555971/File+Systems
func ResourceFileRsyncDaemon() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/file/rsync-daemon"),
		MetaId:           PropId(Id),

		KeyEnabled: PropEnabled("Whether the rsync daemon accepts incoming connections."),
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
