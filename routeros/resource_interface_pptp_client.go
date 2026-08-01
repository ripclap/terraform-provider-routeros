package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
	The menu was empty on the RouterOS 7.23:
	`GET /rest/interface/pptp-client` -> `[]`, so no value sample is available.

	Writable fields reported by the device:
	/console/inspect request=child path="interface,pptp-client,add"
	  add-default-route  allow  comment  connect-to  copy-from  default-route-distance
	  dial-on-demand  disabled  keepalive-timeout  max-mru  max-mtu  mrru  name  password  profile
	  use-peer-dns  user

	Read-only fields reported by the same device: running
*/

// ResourceInterfacePptpClient PPTP client interface.
// https://help.mikrotik.com/docs/spaces/ROS/pages/2031638/PPTP
func ResourceInterfacePptpClient() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/pptp-client"),
		MetaId:           PropId(Id),

		"add_default_route": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Whether to add the PPTP remote address as a default route.",
		},
		"allow": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "Allowed authentication methods.",
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"mschap2", "mschap1", "chap", "pap"}, false),
			},
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyComment: PropCommentRw,
		"connect_to": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Remote address of the PPTP server.",
		},
		"default_route_distance": {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "Distance value applied to the auto created default route, if add-default-route is selected.",
			ValidateFunc: validation.IntBetween(0, 255),
		},
		"dial_on_demand": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Connects to the PPTP server only when outbound traffic is generated.",
		},
		KeyDisabled: PropDisabledRw,
		"keepalive_timeout": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Tunnel keepalive timeout in seconds. Set to `disabled` to keep the tunnel running even " +
				"when the remote end stops responding.",
			DiffSuppressFunc: TimeEqual,
		},
		"max_mru": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Maximum Receive Unit. Maximum packet size that the PPTP interface will be able to receive " +
				"without packet fragmentation.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"max_mtu": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Maximum Transmission Unit. Maximum packet size that the PPTP interface will be able to " +
				"send without packet fragmentation.",
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
		KeyName: PropName("Descriptive name of the interface."),
		"password": {
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   true,
			Description: "Password used for authentication.",
		},
		"profile": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Specifies which PPP profile configuration will be used when establishing the tunnel.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyRunning: PropRunningRo,
		"use_peer_dns": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Whether to use the DNS server settings advertised by the remote server.",
			ValidateFunc:     validation.StringInSlice([]string{"yes", "no", "exclusively"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"user": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "User name used for authentication.",
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
