package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
	{
	  "log-script-errors": "true",
	  "sanitize-names": "false",
	  "tab-width": "4"
	}

	Reference device: RouterOS 7.23.2, `GET /rest/console/settings`.
*/

// ResourceConsoleSettings Global console behaviour settings.
// https://help.mikrotik.com/docs/spaces/ROS/pages/47579162/Console
func ResourceConsoleSettings() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/console/settings"),
		MetaId:           PropId(Id),

		"log_script_errors": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Whether errors produced by scripts are written to the system log.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"sanitize_names": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Whether the console replaces characters that need quoting in names with a safe " +
				"equivalent when a configuration is exported.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"tab_width": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Number of spaces a tab character expands to in the console output (1..8).",
			ValidateFunc:     validation.IntBetween(1, 8),
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
