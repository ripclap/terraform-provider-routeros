package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
	The menu was empty on the reference device (RouterOS 7.23.2):
	`GET /rest/interface/pppoe-server/server` -> `[]`, so no value sample is available.

	Writable fields reported by the device:
	/console/inspect request=child path="interface,pppoe-server,server,add"
	  accept-empty-service  accept-untagged  authentication  comment  copy-from  default-profile
	  disabled  interface  keepalive-timeout  max-mru  max-mtu  max-sessions  mrru
	  one-session-per-host  pado-delay  pppoe-over-vlan-range  service-name

	Read-only fields reported by the same device: invalid

	`max-mru` and `max-mtu` accept the value `auto` in addition to an integer, so they are modelled
	as strings.
*/

// ResourceInterfacePppoeServerService A PPPoE server service running on an interface.
// The static server binding interfaces are managed by `routeros_interface_pppoe_server`.
// https://help.mikrotik.com/docs/spaces/ROS/pages/2031625/PPPoE
func ResourceInterfacePppoeServerService() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/pppoe-server/server"),
		MetaId:           PropId(Id),

		"accept_empty_service": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Whether the server accepts clients that send a PADI message with an empty service-name " +
				"field.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"accept_untagged": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Whether the PPPoE server accepts untagged (non-VLAN) PPPoE packets on its interface when " +
				"`pppoe_over_vlan_range` is specified.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
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
		KeyComment: PropCommentRw,
		"default_profile": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The PPP profile used by default for the connecting clients.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyDisabled: PropDisabledRw,
		KeyInterface: {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Interface that the clients are connected to.",
		},
		KeyInvalid: PropInvalidRo,
		"keepalive_timeout": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Defines the time period (in seconds) after which the router starts sending keepalive " +
				"packets every second. If there is no traffic and no keepalive responses arrive for that period of " +
				"time, the non responding client is proclaimed disconnected. Set to `disabled` to turn the keepalive " +
				"off.",
			DiffSuppressFunc: TimeEqual,
		},
		"max_mru": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Maximum Receive Unit. The optimal value is the MTU of the interface the tunnel is working " +
				"over reduced by 20. Accepts an integer or `auto`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"max_mtu": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Maximum Transmission Unit. The optimal value is the MTU of the interface the tunnel is " +
				"working over reduced by 20. Accepts an integer or `auto`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"max_sessions": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Maximum number of clients that the access concentrator can serve. Accepts an integer or " +
				"`unlimited`.",
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
		"one_session_per_host": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Allow only one session per host (determined by MAC address). If a host tries to establish " +
				"a new session, the old one will be closed.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"pado_delay": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Delay in milliseconds before the PADO packet is sent. Can be used to prefer another access " +
				"concentrator on the same broadcast domain.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"pppoe_over_vlan_range": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Allows the PPPoE server to operate over 802.1Q VLANs. Accepts a range of VLAN IDs as well " +
				"as individual VLANs specified as comma-separated values, e.g. `100-115,120,122,128-130`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"service_name": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The PPPoE service name. The server will accept clients that send a PADI message with a " +
				"service-name that matches this setting, or with an unset service-name field.",
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
