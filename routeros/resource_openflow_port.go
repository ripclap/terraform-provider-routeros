package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
  `GET /rest/openflow/port` returns `[]` on the RouterOS 7.23 - no port is attached to an OpenFlow switch, so no live record could be
  captured. The shape below lists the wire field names reported by the device; the values are
  illustrative.

  {
    ".id": "*1",
    "comment": "",
    "disabled": "false",
    "dynamic": "false",
    "inactive": "false",
    "interface": "ether2",
    "port-id": "1",
    "rx-bytes": "0",
    "rx-packets": "0",
    "switch": "ofswitch1",
    "tx-bytes": "0",
    "tx-packets": "0"
  }
*/

// ResourceOpenflowPort https://help.mikrotik.com/docs/spaces/ROS/pages/295239685/Openflow
func ResourceOpenflowPort() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/openflow/port"),
		MetaId:           PropId(Id),
		MetaSkipFields:   PropSkipFields("rx_bytes", "rx_packets", "tx_bytes", "tx_packets"),

		KeyComment:   PropCommentRw,
		KeyDisabled:  PropDisabledRw,
		KeyDynamic:   PropDynamicRo,
		KeyInactive:  PropInactiveRo,
		KeyInterface: PropInterfaceRw,
		"port_id": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "Port ID used to identify the interface in the OpenFlow flow rules.",
		},
		"switch": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of the OpenFlow switch instance from `/openflow` that will be able to control the port.",
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
