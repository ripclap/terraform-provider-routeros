package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
	{
	  "page-refresh": "300",
	  "store-every": "5min"
	}
*/

// ResourceToolGraphingSettings General graphing settings.
// https://help.mikrotik.com/docs/spaces/ROS/pages/22773810/Graphing
func ResourceToolGraphingSettings() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/tool/graphing"),
		MetaId:           PropId(Id),

		"page_refresh": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Automatic refresh interval of the graphing web pages, in seconds, or `never` to disable " +
				"the automatic refresh." +
				"\n> The property is not purely numeric (`/console/inspect request=completion " +
				"path=tool,graphing input=\"set page-refresh=\"` offers `never`), so it is declared as a string.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"store_every": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "How often to write the collected data to the system drive.",
			ValidateFunc:     validation.StringInSlice([]string{"5min", "hour", "24hours"}, false),
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
