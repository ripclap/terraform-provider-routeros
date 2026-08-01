package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
	The menu was empty on the RouterOS 7.23, so no value sample is available.
*/

// ResourceIpDhcpServerAlert Rogue DHCP server detector.
// https://help.mikrotik.com/docs/spaces/ROS/pages/24805500/DHCP
func ResourceIpDhcpServerAlert() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/dhcp-server/alert"),
		MetaId:           PropId(Id),

		"alert_timeout": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Time after which the alert will be forgotten. If the same server is detected after that " +
				"time, a new alert is generated. `none` means the alert never expires.",
			DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
				if old == new {
					return true
				}

				if old == "none" || new == "none" {
					return false
				}

				return TimeEqual(k, old, new, d)
			},
		},
		KeyComment:   PropCommentRw,
		KeyDisabled:  PropDisabledRw,
		KeyInterface: PropInterfaceRw,
		"on_alert": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Script that is executed when an unknown DHCP server is detected on the interface.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"unknown_server": {
			Type:     schema.TypeString,
			Computed: true,
			Description: "Comma separated list of MAC addresses of the detected unknown DHCP servers. A server is " +
				"removed from this list after `alert_timeout`.",
		},
		"valid_server": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Comma separated list of MAC addresses of the DHCP servers that must not raise an alert.",
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
