package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
	{
	  ".id": "*6C6",
	  "bridge": "bridge1",
	  "disabled": "false",
	  "dynamic": "true",
	  "external": "false",
	  "interface": "ether1",
	  "invalid": "false",
	  "local": "true",
	  "mac-address": "00:00:5E:00:53:01",
	  "on-interface": "ether1"
	}

	Writable fields reported by the reference device (RouterOS 7.23.2):
	/console/inspect request=child path="interface,bridge,host,add"
	  bridge  comment  copy-from  disabled  interface  mac-address  vid

	Read-only fields reported by the same device:
	  aged  aged-peer  dynamic  external  invalid  local  on-interface  remote-ip
*/

// ResourceInterfaceBridgeHost Static entries of the bridge host (MAC) table.
// https://help.mikrotik.com/docs/spaces/ROS/pages/328068/Bridging+and+Switching#BridgingandSwitching-HostTable
func ResourceInterfaceBridgeHost() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/bridge/host"),
		MetaId:           PropId(Id),

		"aged": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Time since the host entry was last refreshed.",
		},
		"aged_peer": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Time since the host entry was last refreshed on the MLAG peer device.",
		},
		"bridge": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The bridge interface the host entry belongs to.",
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		KeyDynamic:  PropDynamicRo,
		"external": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether the host entry is present in the switch chip (hardware) host table.",
		},
		KeyInterface: {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The bridge port the MAC address is reachable through.",
		},
		KeyInvalid: PropInvalidRo,
		"local": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether the host entry is a MAC address of the bridge itself or of one of its ports.",
		},
		KeyMacAddress: PropMacAddressRw("Host's MAC address.", true),
		"on_interface": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The interface the host is currently reachable through.",
		},
		"remote_ip": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The remote IP address the host entry was learned from (MLAG and VXLAN host entries).",
		},
		"vid": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "VLAN ID of the host entry. Only relevant when VLAN filtering is enabled on the bridge.",
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
