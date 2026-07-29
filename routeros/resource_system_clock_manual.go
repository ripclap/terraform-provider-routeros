package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
{
  "dst-delta": "+00:00",
  "dst-end": "1970-01-01 00:00:00",
  "dst-start": "1970-01-01 00:00:00",
  "time-zone": "+00:00"
}
*/

// ResourceSystemClockManual https://wiki.mikrotik.com/wiki/Manual:System/Time#Manual_time_zone_configuration
func ResourceSystemClockManual() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/system/clock/manual"),
		MetaId:           PropId(Id),

		"dst_delta": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The time offset that is added to the GMT offset while the daylight saving time is active, " +
				"written as `+HH:MM` or `-HH:MM`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"dst_end": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Date and time when the daylight saving time ends, written as `YYYY-MM-DD hh:mm:ss` " +
				"(for example `1970-01-01 00:00:00`).",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"dst_start": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Date and time when the daylight saving time starts, written as `YYYY-MM-DD hh:mm:ss` " +
				"(for example `1970-01-01 00:00:00`).",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"time_zone": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The GMT offset that is used when the time zone is configured manually, written as `+HH:MM` " +
				"or `-HH:MM`.",
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
