package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
	The menu was empty on the reference device (RouterOS 7.23.2):
	`GET /rest/interface/l2tp-ether` -> `[]`, so no value sample is available.

	Writable fields reported by the device:
	/console/inspect request=child path="interface,l2tp-ether,add"
	  allow-fast-path  circuit-id  comment  connect-to  cookie-length  copy-from  digest-hash
	  disabled  ipsec-secret  l2tp-proto-version  local-address  local-session-id  local-tunnel-id
	  mac-address  mtu  name  peer-cookie  remote-session-id  remote-tunnel-id  send-cookie
	  unmanaged-mode  use-ipsec  use-l2-specific-sublayer

	Read-only fields reported by the same device:
	  actual-mtu  dynamic  running  unmanaged
*/

// ResourceInterfaceL2tpEther L2TPv3 Ethernet pseudowire interface.
// https://help.mikrotik.com/docs/spaces/ROS/pages/2031631/L2TP
func ResourceInterfaceL2tpEther() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/l2tp-ether"),
		MetaId:           PropId(Id),

		KeyActualMtu:     PropActualMtuRo,
		KeyAllowFastPath: PropAllowFastPathRw,
		"circuit_id": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The virtual circuit identifier used to bind one end of the L2TPv3 control channel.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyComment: PropCommentRw,
		"connect_to": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Remote address of the L2TP server.",
			ValidateFunc:     validation.IsIPAddress,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"cookie_length": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Length of the L2TPv3 pseudowire static session cookie.",
			ValidateFunc:     validation.StringInSlice([]string{"0", "4-bytes", "8-bytes"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"digest_hash": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The hash function used by the L2TPv3 control channel.",
			ValidateFunc:     validation.StringInSlice([]string{"md5", "none", "sha1"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyDisabled:    PropDisabledRw,
		KeyDynamic:     PropDynamicRo,
		KeyIpsecSecret: PropIpsecSecretRw,
		"l2tp_proto_version": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The transport the L2TPv3 pseudowire is carried over.",
			ValidateFunc:     validation.StringInSlice([]string{"l2tpv3-ip", "l2tpv3-udp"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyLocalAddress: PropLocalAddressRw,
		"local_session_id": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The local session identifier, an integer. Only used in unmanaged mode; RouterOS reports " +
				"`disabled` when it is not set, so the value is handled as a string.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"local_tunnel_id": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The local tunnel identifier, an integer. Only used in unmanaged mode; RouterOS reports " +
				"`disabled` when it is not set, so the value is handled as a string.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyMacAddress: PropMacAddressRw("MAC address of the tunnel interface.", false),
		KeyMtu:        PropMtuRw(),
		KeyName:       PropName("Name of the L2TP ether interface."),
		"peer_cookie": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Optional peer cookie. To enable the cookie, enter the remote cookie value (an 8 or 16 " +
				"character hex string).",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"remote_session_id": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The remote session identifier, an integer. Only used in unmanaged mode; RouterOS reports " +
				"`disabled` when it is not set, so the value is handled as a string.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"remote_tunnel_id": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The remote tunnel identifier, an integer. Only used in unmanaged mode; RouterOS reports " +
				"`disabled` when it is not set, so the value is handled as a string.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyRunning: PropRunningRo,
		"send_cookie": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Optional local cookie. To enable the cookie, enter the cookie value (an 8 or 16 character " +
				"hex string).",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"unmanaged": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether the interface is currently operating as an unmanaged (static) pseudowire.",
		},
		"unmanaged_mode": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Activates the unmanaged (static) mode, in which the tunnel and session identifiers are " +
				"configured manually instead of being negotiated by the L2TP control channel.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"use_ipsec": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "When enabled, a dynamic IPsec peer configuration and policy are added to encapsulate the " +
				"L2TP connection into an IPsec tunnel.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"use_l2_specific_sublayer": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Enables the default L2-specific sublayer of the L2TPv3 pseudowire.",
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
