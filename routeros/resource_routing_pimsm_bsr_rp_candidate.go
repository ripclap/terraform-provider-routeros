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
    "group": "",
    "holdtime": "",
    "inactive": "false",
    "instance": "",
    "priority": ""
  }
*/

// ResourceRoutingPimsmBsrRpCandidate https://help.mikrotik.com/docs/spaces/ROS/pages/61767728/PIM-SM
func ResourceRoutingPimsmBsrRpCandidate() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/routing/pimsm/bsr/rp-candidate"),
		MetaId:           PropId(Id),

		"address": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The address advertised as the address of this rendezvous point candidate. The console " +
				"also offers the interface names here, in which case the address of that interface is used.",
		},
		KeyDisabled: PropDisabledRw,
		"group": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The multicast group range this router offers to serve, written as a prefix, for example " +
				"`239.0.0.0/8`.",
		},
		"holdtime": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The hold time advertised in the candidate RP advertisements. The bootstrap router drops " +
				"the candidacy when no advertisement is received within this time.",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEqual,
		},
		KeyInactive: PropInactiveRo,
		"instance": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of the PIM-SM instance this rendezvous point candidate belongs to.",
		},
		"priority": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "The rendezvous point election priority for the advertised group range. A numerically " +
				"lower value is preferred.",
			ValidateFunc:     validation.IntBetween(0, 255),
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
