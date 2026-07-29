package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
{
  "certificate": "none",
  "enabled": "false"
}
*/

/*
	`certificate` and `enabled` are reported but not accepted by `set`, so they are exposed read-only.
*/

// https://help.mikrotik.com/docs/spaces/ROS/pages/295239888/File+share
func ResourceIpCloudBackToHomeFileSettings() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/cloud/back-to-home-file/settings"),
		MetaId:           PropId(Id),

		"certificate": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Certificate used by the File share service. Maintained by the service, it cannot be set.",
		},
		KeyEnabled: {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether the File share service is currently running. Turned on by the service itself as soon as a share exists, it cannot be set.",
		},
		"prefer_relay_code": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Relay code that will be preferred for the File share connection. If it is not set, the " +
				"relay with the smallest RTT is chosen. The available relay codes are listed by `/ip/cloud/print`.",
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
