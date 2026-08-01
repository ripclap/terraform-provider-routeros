package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  {
    ".id": "*1",
    "address": "",
    "disabled": "false",
    "hashmask-length": "",
    "inactive": "false",
    "instance": "",
    "priority": "",
    "scope4": "",
    "scope6": "",
    "state": ""
  }
*/

// ResourceRoutingPimsmBsrCandidate https://help.mikrotik.com/docs/spaces/ROS/pages/61767728/PIM-SM
func ResourceRoutingPimsmBsrCandidate() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/routing/pimsm/bsr/candidate"),
		MetaId:           PropId(Id),

		"address": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The address advertised as the address of this bootstrap router candidate. The console " +
				"also offers the interface names here, in which case the address of that interface is used.",
		},
		KeyDisabled: PropDisabledRw,
		"hashmask_length": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "The length of the hash mask advertised in the bootstrap messages. It controls over how " +
				"many consecutive groups a single rendezvous point is used.",
			ValidateFunc:     validation.IntBetween(0, 128),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyInactive: PropInactiveRo,
		"instance": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of the PIM-SM instance this bootstrap router candidate belongs to.",
		},
		"priority": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "The bootstrap router election priority. The candidate with the highest priority wins the " +
				"election.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"scope4": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The IPv4 administrative scope this candidate is a bootstrap router for. The " +
				"console exposes no value domain for this property and it is not covered by the MikroTik " +
				"documentation - supply the raw RouterOS value.",
		},
		"scope6": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The IPv6 administrative scope this candidate is a bootstrap router for. The " +
				"console exposes no value domain for this property and it is not covered by the MikroTik " +
				"documentation - supply the raw RouterOS value.",
		},
		"state": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The current state of the bootstrap router election for this candidate.",
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
