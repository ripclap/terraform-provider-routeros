package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
	{
	  ".id": "*0",
	  "automap": "true",
	  "default": "true",
	  "disabled": "false",
	  "dscp": "0",
	  "dynamic": "true",
	  "hw-id": "0",
	  "hw-offloaded": "true",
	  "inactive": "false",
	  "name": "default",
	  "pcp": "0",
	  "traffic-class": "1"
	}

	Writable fields reported by the reference device (RouterOS 7.23.2):
	/console/inspect request=child path="interface,ethernet,switch,qos,profile,add"
	  automap  comment  copy-from  disabled  dscp  name  pcp  traffic-class

	The `color` property documented by MikroTik is not exposed by RouterOS 7.23.2 on this switch
	chip, so it is not part of the schema.
*/

// ResourceInterfaceEthernetSwitchQosProfile Switch chip QoS profiles.
// https://help.mikrotik.com/docs/spaces/ROS/pages/189497483/Quality+of+Service
func ResourceInterfaceEthernetSwitchQosProfile() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/ethernet/switch/qos/profile"),
		MetaId:           PropId(Id),

		"automap": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Automatically map the packets whose PCP or DSCP value matches the one of this profile to " +
				"this profile.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyComment:  PropCommentRw,
		KeyDefault:  PropDefaultRo,
		KeyDisabled: PropDisabledRw,
		"dscp": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "IPv4/IPv6 DSCP field value assigned to the egress packets of this QoS profile.",
			ValidateFunc:     validation.IntBetween(0, 63),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyDynamic: PropDynamicRo,
		"hw_id": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "The index the profile occupies in the switch chip.",
		},
		KeyHwOffloaded: {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether the profile is programmed into the switch chip.",
		},
		KeyInactive: PropInactiveRo,
		KeyName:     PropName("Name of the QoS profile."),
		"pcp": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "VLAN priority (IEEE 802.1Q PCP - Priority Code Point) value of this QoS profile.",
			ValidateFunc:     validation.IntBetween(0, 7),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"traffic_class": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Traffic class of this QoS profile. It determines the packet priority and the egress queue " +
				"the packet is placed into.",
			ValidateFunc:     validation.IntBetween(0, 7),
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
