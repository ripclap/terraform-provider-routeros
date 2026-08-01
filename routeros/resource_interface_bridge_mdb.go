package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
	{
	  ".id": "*4394",
	  "bridge": "bridge1",
	  "disabled": "false",
	  "dynamic": "true",
	  "group": "ff02::2",
	  "interface": "",
	  "invalid": "false",
	  "on-interface": "sfp-sfpplus1",
	  "vid": "10"
	}

	Writable fields reported by the RouterOS 7.23:
	/console/inspect request=child path="interface,bridge,mdb,add"
	  bridge  comment  copy-from  disabled  group  interface  vid

	Read-only fields reported by the same device:
	  dynamic  invalid  on-interface

	`interface` accepts more than one bridge port: the console completes another interface name
	after a trailing comma, so it is modeled as a list.
*/

// ResourceInterfaceBridgeMdb Static entries of the bridge multicast database (MDB).
// https://help.mikrotik.com/docs/spaces/ROS/pages/328068/Bridging+and+Switching#BridgingandSwitching-MulticastDatabase
func ResourceInterfaceBridgeMdb() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/bridge/mdb"),
		MetaId:           PropId(Id),

		"bridge": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The bridge interface the multicast group entry belongs to.",
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		KeyDynamic:  PropDynamicRo,
		"group": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The IPv4 or IPv6 multicast group address of the entry.",
		},
		KeyInterface: {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Bridge ports that are going to receive the traffic of this multicast group.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		KeyInvalid: PropInvalidRo,
		"on_interface": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The interface the multicast group was learned on.",
		},
		"vid": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "VLAN ID of the entry. Only relevant when VLAN filtering is enabled on the bridge.",
			ValidateFunc:     validation.IntBetween(0, 4094),
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
