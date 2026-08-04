package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  {
    ".id": "*3",
    "actual-mtu": "1480",
    "clamp-tcp-mss": "true",
    "disabled": "false",
    "dont-fragment": "no",
    "dscp": "inherit",
    "local-address": "0.0.0.0",
    "mtu": "auto",
    "name": "6to4-tunnel1",
    "remote-address": "unspecified",
    "running": "true"
  }
*/

// https://help.mikrotik.com/docs/spaces/ROS/pages/135004174/6to4
func ResourceInterface6to4() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/6to4"),
		MetaId:           PropId(Id),

		KeyActualMtu:    PropActualMtuRo,
		KeyClampTcpMss:  PropClampTcpMssRw,
		KeyComment:      PropCommentRw,
		KeyDisabled:     PropDisabledRw,
		KeyDontFragment: PropDontFragmentRw,
		KeyDscp:         PropDscpRw,
		KeyIpsecSecret:  PropIpsecSecretRw,
		KeyKeepalive:    PropKeepaliveRw,
		KeyLocalAddress: PropLocalAddressRw,
		KeyMtu:          PropMtuRw(),
		KeyName:         PropName("Interface name."),
		// Not PropRemoteAddressRw: this menu reports `unspecified` when no remote
		// end is set, and that is also how the value is cleared. An attribute
		// validated as an IP address alone cannot express it, so a tunnel that
		// once had a remote could not be returned to having none.
		KeyRemoteAddress: {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "IP address of the remote end of the tunnel, or `unspecified`.",
			ValidateFunc: validation.Any(
				validation.IsIPAddress,
				validation.StringInSlice([]string{"unspecified"}, false),
			),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyRunning: PropRunningRo,
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
