package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
	{
	  "authentication": "mschap1,mschap2",
	  "default-profile": "default-encryption",
	  "enabled": "false",
	  "keepalive-timeout": "30",
	  "max-mru": "1450",
	  "max-mtu": "1450",
	  "mrru": "disabled"
	}
*/

// ResourceInterfacePptpServer The PPTP server service.
// The static server binding interfaces are managed by `routeros_interface_pptp_server_binding`.
// https://help.mikrotik.com/docs/spaces/ROS/pages/2031638/PPTP
func ResourceInterfacePptpServer() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/pptp-server/server"),
		MetaId:           PropId(Id),

		"authentication": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "Authentication methods that the server will accept.",
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"mschap2", "mschap1", "chap", "pap"}, false),
			},
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"default_profile": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The PPP profile used by default for the connecting clients.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyEnabled: PropEnabled("Defines whether the PPTP server is enabled or not."),
		"keepalive_timeout": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Defines the time period (in seconds) after which the router starts sending keepalive " +
				"packets every second. Set to `disabled` to turn the keepalive off.",
			DiffSuppressFunc: TimeEqual,
		},
		"max_mru": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Maximum Receive Unit. Maximum packet size that can be received without fragmentation.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"max_mtu": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Maximum Transmission Unit. Maximum packet size that can be sent without fragmentation.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"mrru": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Maximum packet size that can be received on the link. If a packet is bigger than tunnel MTU, " +
				"it will be split into multiple packets, allowing full size IP or Ethernet packets to be sent over the " +
				"tunnel. Set to `disabled` to turn Multilink PPP off.",
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
