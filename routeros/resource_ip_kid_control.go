package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
	The menu was empty on the RouterOS 7.23, so no value sample is available.
*/

// ResourceIpKidControl Kid Control profile (per user time and rate policy).
// https://help.mikrotik.com/docs/spaces/ROS/pages/129531911/Kid+Control
func ResourceIpKidControl() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/kid-control"),
		MetaId:           PropId(Id),

		KeyDisabled: PropDisabledRw,
		"fri": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Time of day when internet access is allowed on Friday, e.g. `11:00:00-22:00:00`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"mon": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Time of day when internet access is allowed on Monday, e.g. `11:00:00-22:00:00`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyName: PropName("Name of the Kid Control profile."),
		"rate_limit": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Maximum data rate available to the devices of this profile, e.g. `3M`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"sat": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Time of day when internet access is allowed on Saturday, e.g. `11:00:00-22:00:00`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"sun": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Time of day when internet access is allowed on Sunday, e.g. `11:00:00-22:00:00`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"thu": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Time of day when internet access is allowed on Thursday, e.g. `11:00:00-22:00:00`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"tue": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Time of day when internet access is allowed on Tuesday, e.g. `11:00:00-22:00:00`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"tur_fri": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Time Unlimited Rate window on Friday: the period during which `rate_limit` is not " +
				"applied. It takes precedence over `rate_limit`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"tur_mon": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Time Unlimited Rate window on Monday: the period during which `rate_limit` is not " +
				"applied. It takes precedence over `rate_limit`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"tur_sat": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Time Unlimited Rate window on Saturday: the period during which `rate_limit` is not " +
				"applied. It takes precedence over `rate_limit`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"tur_sun": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Time Unlimited Rate window on Sunday: the period during which `rate_limit` is not " +
				"applied. It takes precedence over `rate_limit`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"tur_thu": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Time Unlimited Rate window on Thursday: the period during which `rate_limit` is not " +
				"applied. It takes precedence over `rate_limit`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"tur_tue": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Time Unlimited Rate window on Tuesday: the period during which `rate_limit` is not " +
				"applied. It takes precedence over `rate_limit`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"tur_wed": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Time Unlimited Rate window on Wednesday: the period during which `rate_limit` is not " +
				"applied. It takes precedence over `rate_limit`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"wed": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Time of day when internet access is allowed on Wednesday, e.g. `11:00:00-22:00:00`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
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
