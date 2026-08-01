package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
	The menu was empty on the RouterOS 7.23:
	`GET /rest/interface/eoipv6` -> `[]`, so no value sample is available.

	Writable fields reported by the device:
	/console/inspect request=child path="interface,eoipv6,add"
	  arp  arp-timeout  clamp-tcp-mss  comment  copy-from  disabled  dont-fragment  dscp
	  ipsec-secret  keepalive  local-address  loop-protect  loop-protect-disable-time
	  loop-protect-send-interval  mac-address  mtu  name  remote-address  tunnel-id

	Read-only fields reported by the same device:
	  actual-mtu  current-remote-address  l2mtu  loop-protect-status  running

	Unlike `/interface/eoip`, this menu has no `allow-fast-path` property.
*/

// ResourceInterfaceEoipv6 EoIP tunnel over IPv6.
// https://help.mikrotik.com/docs/spaces/ROS/pages/24805521/EoIP
func ResourceInterfaceEoipv6() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/eoipv6"),
		MetaId:           PropId(Id),

		KeyActualMtu:   PropActualMtuRo,
		KeyArp:         PropArpRw,
		KeyArpTimeout:  PropArpTimeoutRw,
		KeyClampTcpMss: PropClampTcpMssRw,
		KeyComment:     PropCommentRw,
		"current_remote_address": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The IPv6 address of the remote end of the tunnel the interface is currently bound to.",
		},
		KeyDisabled:                PropDisabledRw,
		KeyDontFragment:            PropDontFragmentRw,
		KeyDscp:                    PropDscpRw,
		KeyIpsecSecret:             PropIpsecSecretRw,
		KeyKeepalive:               PropKeepaliveRw,
		KeyL2Mtu:                   PropL2MtuRo,
		KeyLocalAddress:            PropLocalAddressRw,
		KeyLoopProtect:             PropLoopProtectRw,
		KeyLoopProtectDisableTime:  PropLoopProtectDisableTimeRw,
		KeyLoopProtectSendInterval: PropLoopProtectSendIntervalRw,
		KeyLoopProtectStatus:       PropLoopProtectStatusRo,
		KeyMacAddress:              PropMacAddressRw("MAC address of the tunnel interface.", false),
		KeyMtu:                     PropMtuRw(),
		KeyName:                    PropName("Name of the EoIPv6 interface."),
		KeyRemoteAddress:           PropRemoteAddressRw,
		KeyRunning:                 PropRunningRo,
		"tunnel_id": {
			Type:        schema.TypeInt,
			Optional:    true,
			Default:     0,
			Description: "Unique tunnel identifier, which must match the other side of the tunnel.",
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
