package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
  The menu is empty on the reference device (ROS 7.23.2, MPLS not configured),
  the field set below is taken from `/console/inspect request=syntax path="mpls,ldp,remote-mapping,add"`
  and from `print proplist=` completion on that device.

  {
    ".id": "*1",
    "comment": "",
    "disabled": "false",
    "dst-address": "10.10.0.0/24",
    "dynamic": "false",
    "inactive": "false",
    "label": "16",
    "nexthop": "10.0.0.2",
    "path": "",
    "peer": "10.0.0.2:0",
    "pw-fec": "",
    "vpls": "false",
    "vrf": "main"
  }
*/

// ResourceMplsLdpRemoteMapping https://help.mikrotik.com/docs/spaces/ROS/pages/121995275/LDP
func ResourceMplsLdpRemoteMapping() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/mpls/ldp/remote-mapping"),
		MetaId:           PropId(Id),

		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"dst_address": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Prefix the remote label is bound to, written as `address/prefix-length`. " +
				"Both IPv4 and IPv6 prefixes are accepted.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyDynamic:  PropDynamicRo,
		KeyInactive: PropInactiveRo,
		"label": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Label the remote LSR uses for the prefix. Either a label number (the console grammar " +
				"accepts 16..1048576) or one of the reserved names `alert`, `expl-null`, `expl-null6`, `impl-null`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"nexthop": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Next-hop the labeled packets are sent to. The console grammar accepts either an " +
				"IPv4/IPv6 address or an interface name.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"path": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Path vector reported by the LDP loop detection.",
		},
		"peer": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Peer the mapping has been received from, in the `address:label-space` notation.",
		},
		"pw_fec": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Pseudowire FEC of the mapping, set for the VPLS pseudowire bindings.",
		},
		"vpls": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether the mapping belongs to a VPLS pseudowire.",
		},
		KeyVrf: PropVrfRw,
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
