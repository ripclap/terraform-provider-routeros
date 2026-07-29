package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
  {
    ".id": "*1",
    "active": "false",
    "active-downstream-interfaces": "",
    "bytes": "0",
    "disabled": "false",
    "downstream-interfaces": "",
    "dynamic": "false",
    "group": "",
    "packets": "0",
    "source": "",
    "upstream-interface": "",
    "wrong-packets": "0"
  }
*/

// ResourceRoutingIgmpProxyMfc A static entry of the multicast forwarding cache of the IGMP proxy. It pins the
// forwarding of a (source, group) pair from one upstream interface to a set of downstream interfaces.
// https://help.mikrotik.com/docs/spaces/ROS/pages/128221386/IGMP+Proxy
func ResourceRoutingIgmpProxyMfc() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/routing/igmp-proxy/mfc"),
		MetaId:           PropId(Id),
		MetaSkipFields:   PropSkipFields("bytes", "packets", "wrong_packets"),

		"active": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether the entry is currently installed in the multicast forwarding cache.",
		},
		"active_downstream_interfaces": {
			Type:     schema.TypeSet,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "The downstream interfaces the traffic is currently replicated to.",
		},
		KeyDisabled: PropDisabledRw,
		"downstream_interfaces": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "Interfaces the multicast traffic of this entry is forwarded to.",
		},
		KeyDynamic: PropDynamicRo,
		"group": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The multicast group address of the forwarding entry.",
		},
		"source": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The source address of the multicast traffic.",
		},
		"upstream_interface": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The interface the multicast traffic of this entry is received on.",
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
