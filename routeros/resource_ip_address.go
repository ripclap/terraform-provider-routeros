package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  {
    ".id": "*2",
    "actual-interface": "vlan2",
    "address": "10.0.0.1/24",
    "disabled": "false",
    "dynamic": "false",
    "interface": "vlan15",
    "invalid": "false",
    "network": "10.0.0.0"
  },
*/

// ResourceIPAddress https://wiki.mikrotik.com/wiki/Manual:IP/Address
func ResourceIPAddress() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/address"),
		MetaId:           PropId(Id),

		"address": {
			Type:         schema.TypeString,
			Required:     true,
			Description:  "IP address.",
			ValidateFunc: ValidationIpAddress,
		},
		"actual_interface": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Name of the actual interface the logical one is bound to.",
		},
		"broadcast": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			Description: "Broadcast address of the network. It is derived from the address and the netmask, and only " +
				"needs to be set when a non-standard broadcast address is required.",
			ValidateFunc: validation.IsIPAddress,
		},
		KeyComment:   PropCommentRw,
		KeyDisabled:  PropDisabledRw,
		KeyDynamic:   PropDynamicRo,
		KeyInterface: PropInterfaceRw,
		KeyInvalid:   PropInvalidRo,
		"netmask": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Delimits the network address part of the IP address from the host part. It is derived from " +
				"the prefix length of the address when not set.",
			ValidateFunc: validation.IntBetween(0, 32),
		},
		"network": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			Description: "IP address for the network. For point-to-point links it should be the address of the " +
				"remote end. Starting from v5RC6 this parameter is configurable only for addresses with /32 netmask " +
				"(point to point links)",
			ValidateFunc: validation.IsIPAddress,
		},
		"slave": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether address belongs to an interface which is a slave port to some other master interface",
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
