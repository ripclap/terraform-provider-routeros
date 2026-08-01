package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
  "cpu-overtemp-check": "false",
  "cpu-overtemp-startup-delay": "1m",
  "cpu-overtemp-threshold": "105",
  "fan-control-interval": "30s",
  "fan-full-speed-temp": "65",
  "fan-min-speed-percent": "12",
  "fan-target-temp": "58"
}
*/

// ResourceSystemHealthSettings https://help.mikrotik.com/docs/spaces/ROS/pages/25690117/Health
func ResourceSystemHealthSettings() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/system/health/settings"),
		MetaId:           PropId(Id),

		"active_fan": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The fan that is currently used by the system.",
		},
		"cpu_overtemp_check": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Enables the CPU overtemperature protection. When the CPU temperature reaches " +
				"`cpu-overtemp-threshold`, the device is rebooted. Available on ARM and ARM64 devices.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"cpu_overtemp_startup_delay": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The time after the system startup during which the CPU overtemperature protection is not " +
				"active yet.",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEqual,
		},
		"cpu_overtemp_threshold": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "The CPU temperature (in degrees Celsius) at which the overtemperature protection triggers.",
			ValidateFunc:     validation.IntBetween(0, 105),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"fan_control_interval": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The interval at which the actual temperature values are read from the CPU, PHY, SWITCH and " +
				"SFP sensors in order to adjust the fan speed.",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEqual,
		},
		"fan_full_speed_temp": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "The temperature (in degrees Celsius) at which the fan speed is increased to the maximum " +
				"possible RPM.",
			ValidateFunc:     validation.IntBetween(-273, 65),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"fan_min_speed_percent": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "The minimum fan speed, in percent, thus not allowing the fans to spin slower than this " +
				"value.",
			ValidateFunc:     validation.IntBetween(0, 100),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"fan_mode": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Fan control mode." +
				"\n  * auto - the fan speed is adjusted by the system according to the measured temperatures;" +
				"\n  * manual - the fan state is controlled by the `fan_switch` property." +
				"\nOnly available on devices with a controllable fan.",
			ValidateFunc:     validation.StringInSlice([]string{"auto", "manual"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"fan_on_threshold": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The threshold at which the fan is switched on. The property is accepted by " +
				"`/system/health/settings/set` but is not reported by a RouterOS 7.23 device, so neither its unit " +
				"(percent or degrees Celsius) nor its valid range could be established on this hardware.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"fan_switch": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The state of the fan when `fan_mode` is `manual`." +
				"\n  * auto - the fan is controlled by the system;" +
				"\n  * off - the fan is always off;" +
				"\n  * on - the fan is always on.",
			ValidateFunc:     validation.StringInSlice([]string{"auto", "off", "on"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"fan_target_temp": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "The target temperature (in degrees Celsius) of the hottest component. If at least one of " +
				"the internally measured temperatures exceeds this value, the fans start to spin; the higher the " +
				"temperature, the faster the fans spin.",
			ValidateFunc:     validation.IntBetween(-273, 65),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"use_fan": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Selects which of the fan connectors is used on devices that have more than one." +
				"\n  * auxiliary;" +
				"\n  * main.",
			ValidateFunc:     validation.StringInSlice([]string{"auxiliary", "main"}, false),
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
