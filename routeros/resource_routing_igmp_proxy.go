package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
{
  "query-interval": "2m5s",
  "query-response-interval": "10s",
  "quick-leave": "false"
}
*/

// ResourceRoutingIgmpProxy The global settings of the IGMP proxy.
// https://help.mikrotik.com/docs/spaces/ROS/pages/128221386/IGMP+Proxy
func ResourceRoutingIgmpProxy() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/routing/igmp-proxy"),
		MetaId:           PropId(Id),

		"query_interval": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The interval between the IGMP general membership queries sent on the downstream interfaces.",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEqual,
		},
		"query_response_interval": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The maximum response time the router advertises in the IGMP membership queries. A " +
				"membership is considered gone when no report is received within this time.",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEqual,
		},
		"quick_leave": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "If enabled, the router stops forwarding a group as soon as an IGMP leave message is " +
				"received, without sending the group specific queries first. Only use it when a single receiver is " +
				"attached to the downstream interface.",
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
