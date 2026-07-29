package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
{
  "auto-send-supout": "false",
  "automatic-supout": "true",
  "ping-start-after-boot": "5m",
  "ping-timeout": "1m",
  "watch-address": "none",
  "watchdog-timer": "true"
}
*/

// ResourceSystemWatchdog https://help.mikrotik.com/docs/spaces/ROS/pages/8978694/Watchdog
func ResourceSystemWatchdog() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/system/watchdog"),
		MetaId:           PropId(Id),

		"auto_send_supout": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "After the supout file is automatically generated, it can be sent by e-mail to the address " +
				"configured in `send_email_to`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"automatic_supout": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "When a software failure happens, a file named `autosupout.rif` is generated automatically. " +
				"The previous `autosupout.rif` file is renamed to `autosupout.old.rif`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"ping_start_after_boot": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Specifies how long the system waits after boot before it starts to reach the " +
				"`watch_address` (known as `no-ping-delay` in older RouterOS versions).",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEqual,
		},
		"ping_timeout": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Specifies the time interval in which the device will be pinged 6 times (after " +
				"`ping_start_after_boot`).",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEqual,
		},
		"send_email_from": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The e-mail address the supout file is sent from. If not set, the value from `/tool/e-mail` " +
				"is used.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"send_email_to": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The e-mail address to send the supout file to.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"send_smtp_server": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The SMTP server used to send the supout file. If not set, the value from `/tool/e-mail` " +
				"is used.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"watch_address": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The system will reboot in case 6 sequential pings to the given IP address fail. " +
				"`none` disables the ping watchdog.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"watchdog_timer": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Whether to reboot the device if the system is unresponsive for a minute.",
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
