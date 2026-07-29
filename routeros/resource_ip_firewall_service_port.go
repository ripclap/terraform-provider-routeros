package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
{
  ".id": "*5",
  "disabled": "false",
  "invalid": "false",
  "name": "sip",
  "ports": "5060",
  "sip-direct-media": "false",
  "sip-timeout": "1h"
}
*/

/*
	The connection tracking helpers are a fixed list that can be modified but not added or removed.
	`ports`, `sip-direct-media` and `sip-timeout` are only reported for the helpers that support them.
*/

// ResourceIPFirewallServicePort Connection tracking helpers (NAT helpers).
// https://help.mikrotik.com/docs/spaces/ROS/pages/250708066/Firewall
func ResourceIPFirewallServicePort() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/firewall/service-port"),
		MetaId:           PropId(Id),
		// The helper is looked up by name, the name itself cannot be changed.
		MetaSkipFields: PropSkipFields("name"),

		KeyDisabled: PropDisabledRw,
		KeyInvalid:  PropInvalidRo,
		KeyName: PropName("Name of the connection tracking helper: `ftp`, `tftp`, `irc`, `h323`, `sip`, `pptp`, " +
			"`rtsp`, `udplite`, `dccp` or `sctp`."),
		"ports": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			Description: "Comma separated list of ports the helper listens on. Only the helpers that are port based " +
				"(`ftp`, `tftp`, `irc`, `rtsp`, `sip`) report and accept this property.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"sip_direct_media": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Allows the RTP/RTCP media stream to flow directly between the endpoints. `sip` helper only.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"sip_timeout": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Maximum duration of a SIP call kept by the helper. `sip` helper only.",
			DiffSuppressFunc: TimeEqual,
		},
	}

	return &schema.Resource{
		CreateContext: DefaultCreateUpdate(resSchema),
		ReadContext:   DefaultRead(resSchema),
		UpdateContext: DefaultCreateUpdate(resSchema),
		DeleteContext: DefaultSystemDelete(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: ImportStateCustomContext(resSchema),
		},

		Schema: resSchema,
	}
}
