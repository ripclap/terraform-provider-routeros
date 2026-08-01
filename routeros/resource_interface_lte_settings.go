package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
	{
	  "esim-channel": "auto",
	  "firmware-path": "firmware",
	  "link-recovery-timer": "120",
	  "mode": "auto"
	}

	The hardware dependent properties `external-antenna`, `external-antenna-selected` and
	`sim-slot` that MikroTik documents for some LTE products are not exposed by the reference
	device (RouterOS 7.23) and are therefore not part of the schema.
*/

// ResourceInterfaceLteSettings Global LTE modem settings.
// https://help.mikrotik.com/docs/spaces/ROS/pages/30146563/LTE+5G
func ResourceInterfaceLteSettings() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/lte/settings"),
		MetaId:           PropId(Id),

		"esim_channel": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The channel used to communicate with the eSIM applet of the modem. `at` forces the AT " +
				"command channel, `auto` lets RouterOS pick the channel.",
			ValidateFunc:     validation.StringInSlice([]string{"at", "auto"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"firmware_path": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Path in the router file system where the modem firmware images are stored.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"link_recovery_timer": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Time without a working data link after which RouterOS restarts the modem link." +
				"\nThis property is exposed by RouterOS 7.23 but is not covered by the MikroTik " +
				"documentation; only the name and the default value (120) are known.",
			DiffSuppressFunc: TimeEqual,
		},
		"mode": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Operation mode of the modem interface." +
				"\n  * auto - RouterOS selects the mode automatically." +
				"\n  * mbim - switch the modem to MBIM mode if it is supported." +
				"\n  * serial - expose the modem serial ports only, so that it can be driven by `/interface/ppp-client`." +
				"\n  * user - do not switch the modem mode automatically.",
			ValidateFunc:     validation.StringInSlice([]string{"auto", "mbim", "serial", "user"}, false),
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
