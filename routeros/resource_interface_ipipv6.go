package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
	The menu was empty on the reference device (RouterOS 7.23.2):
	`GET /rest/interface/ipipv6` -> `[]`, so no value sample is available.

	Writable fields reported by the device:
	/console/inspect request=child path="interface,ipipv6,add"
	  clamp-tcp-mss  comment  copy-from  disabled  dont-fragment  dscp  ipsec-secret  keepalive
	  local-address  mtu  name  remote-address

	Read-only fields reported by the same device:
	  actual-mtu  current-remote-address  running

	Unlike `/interface/ipip`, this menu has neither an `allow-fast-path` nor an `l2mtu` property.
*/

// ResourceInterfaceIPIPv6 IPIP tunnel over IPv6.
// https://help.mikrotik.com/docs/spaces/ROS/pages/24805500/IPIP
func ResourceInterfaceIPIPv6() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/ipipv6"),
		MetaId:           PropId(Id),

		KeyActualMtu:   PropActualMtuRo,
		KeyClampTcpMss: PropClampTcpMssRw,
		KeyComment:     PropCommentRw,
		"current_remote_address": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The IPv6 address of the remote end of the tunnel the interface is currently bound to.",
		},
		KeyDisabled:      PropDisabledRw,
		KeyDontFragment:  PropDontFragmentRw,
		KeyDscp:          PropDscpRw,
		KeyIpsecSecret:   PropIpsecSecretRw,
		KeyKeepalive:     PropKeepaliveRw,
		KeyLocalAddress:  PropLocalAddressRw,
		KeyMtu:           PropMtuRw(),
		KeyName:          PropName("Name of the IPIPv6 interface."),
		KeyRemoteAddress: PropRemoteAddressRw,
		KeyRunning:       PropRunningRo,
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
