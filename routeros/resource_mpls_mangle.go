package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  The menu is empty on RouterOS 7.23 (ROS 7.23, MPLS not configured),
  the field set below is taken from `/console/inspect request=syntax path="mpls,mangle,add"` and
  from `print proplist=` completion on that device.

  {
    ".id": "*1",
    "chain": "forward",
    "comment": "",
    "disabled": "false",
    "exp": "0",
    "packets": "0",
    "set-exp": "3",
    "set-mark": "m0"
  }
*/

// ResourceMplsMangle https://help.mikrotik.com/docs/spaces/ROS/pages/122388503/EXP+bit+and+MPLS+Queuing
func ResourceMplsMangle() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/mpls/mangle"),
		MetaId:           PropId(Id),
		MetaSkipFields:   PropSkipFields("packets"),

		"chain": {
			Type:     schema.TypeString,
			Required: true,
			Description: "Chain the rule belongs to. `forward` processes the transit labeled packets, `output` " +
				"processes the labeled packets originated by this router.",
			ValidateFunc: validation.StringInSlice([]string{"forward", "output"}, false),
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"exp": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Matches the EXP (traffic class) bits of the topmost label of the packet.",
			ValidateFunc:     validation.IntBetween(0, 7),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"set_exp": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "New EXP (traffic class) bit value written into the topmost label of the matched packet.",
			ValidateFunc:     validation.IntBetween(0, 7),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"set_mark": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Packet mark assigned to the matched packet so that it can be picked up by the queue " +
				"tree. The value `no-mark` removes the mark.",
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
