package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  {
    ".id": "*1",
    "cost": "",
    "disabled": "false",
    "instance": "",
    "interfaces": "",
    "key-chain": "",
    "mode": "",
    "password": "",
    "poison-reverse": "false",
    "source-addresses": "",
    "split-horizon": "false",
    "use-bfd": "false"
  }
*/

// ResourceRoutingRipInterfaceTemplate https://help.mikrotik.com/docs/spaces/ROS/pages/328211/RIP
func ResourceRoutingRipInterfaceTemplate() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/routing/rip/interface-template"),
		MetaId:           PropId(Id),

		"cost": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "The metric that is added to the routes received on the matching interfaces. In RIP the " +
				"metric 16 means infinity, i.e. an unreachable destination.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyDisabled: PropDisabledRw,
		"instance": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of the RIP instance the matching interfaces are attached to.",
		},
		"interfaces": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "Interfaces to match. Both interface names and interface list names are accepted.",
		},
		"key_chain": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Name of the `/routing/rip/keys` chain used to authenticate the RIP messages on the " +
				"matching interfaces.",
		},
		"mode": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The operation mode of the matching interfaces." +
				"\n  * passive - do not send the RIP updates, only receive and process them" +
				"\n  * strict - Not covered by the MikroTik documentation, the value is reported by the " +
				"RouterOS console completion",
			ValidateFunc: validation.StringInSlice([]string{"passive", "strict"}, false),
		},
		"password": {
			Type:      schema.TypeString,
			Optional:  true,
			Sensitive: true,
			Description: "The plain text authentication password used on the matching interfaces. Use `key_chain` " +
				"for the cryptographic authentication instead.",
		},
		"poison_reverse": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Advertise the routes learned on an interface back through the same interface with an " +
				"infinite metric instead of omitting them.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"source_addresses": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "The local addresses used as the source of the RIP messages sent on the matching " +
				"interfaces.",
		},
		"split_horizon": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Do not advertise the routes learned on an interface back through the same interface.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"use_bfd": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Whether to use the BFD protocol for faster connection state detection.",
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
