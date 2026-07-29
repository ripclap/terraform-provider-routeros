package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  {
    ".id": "*1",
    "address": "",
    "area": "",
    "auth": "",
    "auth-id": "",
    "auth-key": "",
    "bdr": "",
    "cost": "",
    "dead-interval": "",
    "disabled": "false",
    "dr": "",
    "dynamic": "false",
    "hello-interval": "",
    "instance-id": "",
    "interface": "",
    "passive": "",
    "priority": "",
    "retransmit-interval": "",
    "state": "",
    "transmit-delay": "",
    "type": "",
    "use-bfd": "",
    "vlink-neighbor-id": "",
    "vlink-remote-address": "",
    "vlink-transit-area": ""
  }
*/

// This menu has no `comment` property; RouterOS answers `unknown parameter comment` on `add` and `set`.

// ResourceRoutingOspfInterface https://help.mikrotik.com/docs/spaces/ROS/pages/9863229/OSPF
func ResourceRoutingOspfInterface() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath:   PropResourcePath("/routing/ospf/interface"),
		MetaId:             PropId(Id),
		MetaSetUnsetFields: PropSetUnsetFields("passive"),

		"address": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The local address OSPF is run on. It selects one of the addresses of the interface when " +
				"more than one is configured.",
		},
		"area": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The OSPF area this interface belongs to.",
		},
		"auth": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "Specifies the authentication method for the OSPF protocol messages.",
			ValidateFunc: validation.StringInSlice([]string{"simple", "md5", "sha1", "sha256", "sha384", "sha512"}, true),
		},
		"auth_id": {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "The key id used to calculate the message digest when MD5 or SHA authentication is enabled.",
			ValidateFunc: validation.IntBetween(0, 255),
		},
		"auth_key": {
			Type:      schema.TypeString,
			Optional:  true,
			Sensitive: true,
			Description: "The authentication key to be used, it should match on all the neighbors of the network " +
				"segment.",
		},
		"bdr": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The router ID of the current backup designated router on this segment.",
		},
		"cost": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Interface cost expressed as the link state metric.",
			ValidateFunc:     Validation64k,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"dead_interval": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Specifies the interval after which a neighbor is declared dead.",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEqual,
		},
		KeyDisabled: PropDisabledRw,
		"dr": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The router ID of the current designated router on this segment.",
		},
		KeyDynamic: PropDynamicRo,
		"hello_interval": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The interval between the HELLO packets that the router sends out this interface.",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEqual,
		},
		"instance_id": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "The OSPF instance id carried in the packets sent out this interface.",
			ValidateFunc:     validation.IntBetween(0, 255),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyInterface: PropInterfaceRw,
		"passive": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
			Description: "If enabled, then do not send or receive OSPF traffic on this interface, but still " +
				"advertise the attached network." +
				"\n<em>The correct value of this attribute may not be displayed in Winbox. " +
				"Please check the parameters in the console!</em>",
		},
		"priority": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Router's priority. Used to determine the designated router in a broadcast network.",
			ValidateFunc:     validation.IntBetween(0, 255),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"retransmit_interval": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Time interval after which a lost link state advertisement will be resent.",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEqual,
		},
		"state": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The current OSPF interface state.",
		},
		"transmit_delay": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Link-state transmit delay is the estimated time it takes to transmit a link-state update " +
				"packet on the interface.",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEqual,
		},
		"type": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The OSPF network type on this interface.",
			ValidateFunc: validation.StringInSlice([]string{"broadcast", "nbma", "ptmp", "ptmp-broadcast", "ptp",
				"ptp-unnumbered", "virtual-link"}, true),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"use_bfd": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Whether to use the BFD protocol for faster connection state detection.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"vlink_neighbor_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Specifies the router id of the neighbor which should be connected over the virtual link.",
		},
		"vlink_remote_address": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The address of the remote end of the established virtual link.",
		},
		"vlink_transit_area": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "A non-backbone area the two routers have in common over which the virtual link will be " +
				"established.",
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
