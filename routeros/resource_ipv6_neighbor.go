package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
  A static entry:
  {
    ".id": "*1C",
    "address": "fe80::1",
    "disabled": "false",
    "dynamic": "false",
    "interface": "ether1",
    "mac-address": "00:00:5E:00:53:02",
    "status": "permanent",
    "vrf": "main"
  }

  A discovered (dynamic) entry, showing the extra read-only "router" flag:
  {
    ".id": "*2D4",
    "address": "fe80::2",
    "disabled": "false",
    "dynamic": "true",
    "interface": "vlan10",
    "mac-address": "00:00:5E:00:53:03",
    "router": "true",
    "status": "reachable",
    "vrf": "main"
  }
*/

// ResourceIPv6Neighbor https://help.mikrotik.com/docs/spaces/ROS/pages/40992815/IPv6+Neighbor+Discovery
func ResourceIPv6Neighbor() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ipv6/neighbor"),
		MetaId:           PropId(Id),

		"address": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "IPv6 address of the neighbor.",
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		KeyDynamic:  PropDynamicRo,
		KeyInterface: {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Interface the neighbor is reachable through.",
		},
		KeyMacAddress: PropMacAddressRw("Link-layer address of the neighbor.", false),
		"router": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether the discovered node announces itself as a router.",
		},
		"status": {
			Type:     schema.TypeString,
			Computed: true,
			Description: "Neighbor cache state of the entry. Values observed on the device: `permanent` (manually " +
				"added), `reachable`, `stale`, `failed`; RouterOS also documents `noarp`, `incomplete`, `delay` " +
				"and `probe`.",
		},
		KeyVrf: {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "VRF the neighbor entry belongs to. Derived from the interface, it is not settable.",
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
