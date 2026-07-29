package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
{
  "thumbnails": ""
}
*/

// ResourceIpMediaSettings Global DLNA media server settings.
// https://help.mikrotik.com/docs/spaces/ROS/pages/237699479/DLNA+Media+server
func ResourceIpMediaSettings() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/media/settings"),
		MetaId:           PropId(Id),

		"thumbnails": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Location used by the media server for the generated thumbnails; empty by default. " +
				"VERIFY: this is the only property `/ip/media/settings/set` accepts on RouterOS 7.23.2, but " +
				"MikroTik does not document it, so the exact accepted values were not established.",
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
