package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
  ".id": "*1",
  "delay-mode": "auto",
  "disabled": "false",
  "domain": "auto",
  "inactive": "false",
  "name": "ptp1",
  "priority1": "auto",
  "priority2": "auto",
  "profile": "default",
  "transport": "auto"
}
*/

// ResourceSystemPtp A PTP (IEEE 1588, Precision Time Protocol) instance. Requires hardware with PTP
// support; the interfaces are attached to the instance with `routeros_system_ptp_port`.
// https://help.mikrotik.com/docs/spaces/ROS/pages/64127015/Precision+Time+Protocol
func ResourceSystemPtp() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/system/ptp"),
		MetaId:           PropId(Id),

		KeyComment: PropCommentRw,
		"delay_mode": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The delay measurement mechanism." +
				"\n  * auto - selected by the profile;" +
				"\n  * e2e - end-to-end (delay request-response);" +
				"\n  * p2p - peer-to-peer (peer delay).",
			ValidateFunc:     validation.StringInSlice([]string{"auto", "e2e", "p2p"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyDisabled: PropDisabledRw,
		"domain": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The PTP domain number, used to separate different PTP instances in the same network. " +
				"An integer `0..127`, or `auto` to take the value from the profile.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyInactive: PropInactiveRo,
		KeyName:     PropName("Name of the PTP instance."),
		"priority1": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The first grandmaster election parameter, an integer `0..255` (lower wins), or `auto` " +
				"to take the value from the profile.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"priority2": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The second grandmaster election parameter used to break ties, an integer `0..255` " +
				"(lower wins), or `auto` to take the value from the profile.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"profile": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The PTP profile. Each profile comes with its own predefined `auto` values for the PTP " +
				"operating parameters." +
				"\n  * 802.1as;" +
				"\n  * aes67;" +
				"\n  * default;" +
				"\n  * g8275.1;" +
				"\n  * smpte-2059.",
			ValidateFunc: validation.StringInSlice([]string{"802.1as", "aes67", "default", "g8275.1",
				"smpte-2059"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"transport": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The transport used to carry the PTP messages." +
				"\n  * auto - selected by the profile;" +
				"\n  * ipv4 - UDP over IPv4;" +
				"\n  * l2-forwardable - Ethernet, forwardable multicast address;" +
				"\n  * l2-non-forwardable - Ethernet, link-local multicast address.",
			ValidateFunc: validation.StringInSlice([]string{"auto", "ipv4", "l2-forwardable",
				"l2-non-forwardable"}, false),
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
