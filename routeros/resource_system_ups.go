package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
  ".id": "*1",
  "alarm-setting": "immediate",
  "check-capabilities": "true",
  "disabled": "false",
  "invalid": "false",
  "min-runtime": "never",
  "name": "ups1",
  "offline-time": "0s",
  "on-line": "true",
  "port": "serial0"
}
*/

// ResourceSystemUps Monitoring of an APC Smart Protocol compatible UPS attached to a serial or USB port.
// https://help.mikrotik.com/docs/spaces/ROS/pages/120324130/UPS
func ResourceSystemUps() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/system/ups"),
		MetaId:           PropId(Id),

		"alarm_setting": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The UPS sound alarm setting." +
				"\n  * delayed - alarm after the on-battery event;" +
				"\n  * immediate - alarm immediately after the on-battery event;" +
				"\n  * low-battery - alarm only when the battery is low;" +
				"\n  * none - do not alarm.",
			ValidateFunc: validation.StringInSlice([]string{"delayed", "immediate", "low-battery",
				"none"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"check_capabilities": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Whether to query the UPS capabilities before reading the information from it. " +
				"Disabling this can fix compatibility issues with some UPS models.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		KeyInvalid:  PropInvalidRo,
		"load": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The output load of the UPS, in percent of the rated capacity.",
		},
		"manufacture_date": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The manufacturing date of the UPS, in `mm/dd/yy` format.",
		},
		"min_runtime": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The minimal remaining run time. After a utility failure the router monitors the remaining " +
				"run time and goes into hibernate mode once it drops to this value. `never` uses the 10% battery " +
				"threshold instead, `0s` keeps running until the battery is exhausted.",
			// Not TimeEqual: the property also accepts the literal `never`, which is not a duration.
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"model": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The model name reported by the UPS.",
		},
		KeyName: PropName("Name of the UPS instance."),
		"nominal_battery_voltage": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The nominal battery voltage reported by the UPS.",
		},
		"offline_after": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The calculated moment at which the router goes offline while running on battery.",
		},
		"offline_time": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "How long to keep working on batteries before going into hibernate mode until the UPS " +
				"reports that the utility power is back. `0s` defers the decision to `min_runtime`.",
			DiffSuppressFunc: TimeEqual,
		},
		"on_line": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether the UPS is currently powered from the utility line.",
		},
		"port": {
			Type:     schema.TypeString,
			Required: true,
			Description: "The communication port of the router the UPS is attached to, as listed by `/port/print` " +
				"(for example `serial0` or `usb1`).",
		},
		"serial": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The serial number reported by the UPS.",
		},
		"version": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The SKU, firmware revision and country code reported by the UPS.",
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
