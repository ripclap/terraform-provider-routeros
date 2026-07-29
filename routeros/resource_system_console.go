package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
{
  ".id": "*1",
  "channel": "0",
  "default": "true",
  "disabled": "false",
  "free": "true",
  "port": "serial0",
  "term": "vt102"
}
*/

// ResourceSystemConsole https://help.mikrotik.com/docs/spaces/ROS/pages/328139/Serial+Console
func ResourceSystemConsole() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/system/console"),
		MetaId:           PropId(Id),

		"channel": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "The console channel number that is attached to the port.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyDefault:  PropDefaultRo,
		KeyDisabled: PropDisabledRw,
		"free": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether the port is currently free (no session is attached to the console).",
		},
		"port": {
			Type:     schema.TypeString,
			Required: true,
			Description: "The name of the serial or USB port the console is bound to, as listed by `/port/print` " +
				"(for example `serial0`).",
		},
		"term": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The terminal type reported to the attached terminal, for example `vt102`, `vt100`, " +
				"`linux` or `ansi`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"used": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether the console is currently in use.",
		},
		"vcno": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The number of the virtual console that is attached to the port.",
		},
		"wedged": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether the console is stuck and does not accept input.",
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
