package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
  The menu is empty on RouterOS 7.23 (ROS 7.23, MPLS not configured),
  the field set below is taken from `/console/inspect request=syntax path="mpls,ldp,local-mapping,add"`
  and from `print proplist=` completion on that device.

  {
    ".id": "*1",
    "adv-path": "",
    "comment": "",
    "disabled": "false",
    "dst-address": "10.10.0.0/24",
    "dynamic": "false",
    "egress": "false",
    "gateway": "false",
    "inactive": "false",
    "label": "impl-null",
    "local": "true",
    "peers": "",
    "pw-fec": "",
    "vpls": "false",
    "vrf": "main"
  }
*/

// ResourceMplsLdpLocalMapping https://help.mikrotik.com/docs/spaces/ROS/pages/121995275/LDP
func ResourceMplsLdpLocalMapping() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/mpls/ldp/local-mapping"),
		MetaId:           PropId(Id),

		"adv_path": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Advertised path, filled in by the router.",
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"dst_address": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Prefix the local label is bound to, written as `address/prefix-length`. " +
				"Both IPv4 and IPv6 prefixes are accepted.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyDynamic: PropDynamicRo,
		"egress": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether this router is the egress LSR for the prefix.",
		},
		"gateway": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether the mapping is installed towards a gateway.",
		},
		KeyInactive: PropInactiveRo,
		"label": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Label advertised for the prefix. Either a label number (the console grammar accepts " +
				"16..1048576) or one of the reserved names `alert`, `expl-null`, `expl-null6`, `impl-null`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"local": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether the mapping is originated locally.",
		},
		"peers": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Peers the mapping has been advertised to, in the `address:label-space` notation.",
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
