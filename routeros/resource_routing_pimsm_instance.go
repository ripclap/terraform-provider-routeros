package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  {
    ".id": "*1",
    "afi": "",
    "bsm-forward-back": "false",
    "crp-advertise-contained": "false",
    "disabled": "false",
    "inactive": "false",
    "name": "",
    "rp-hash-mask-length": "",
    "rp-static-override": "false",
    "ssm-range": "",
    "switch-to-spt": "false",
    "switch-to-spt-bytes": "",
    "switch-to-spt-interval": "",
    "vrf": "main"
  }
*/

// ResourceRoutingPimsmInstance https://help.mikrotik.com/docs/spaces/ROS/pages/61767728/PIM-SM
func ResourceRoutingPimsmInstance() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/routing/pimsm/instance"),
		MetaId:           PropId(Id),

		"afi": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "The address family this instance operates on.",
			ValidateFunc: validation.StringInSlice([]string{"ip", "ipv6"}, false),
		},
		"bsm_forward_back": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Allow the bootstrap messages to be forwarded back to the interface they were received " +
				"on.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"crp_advertise_contained": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "When acting as a candidate RP, advertise only the group ranges that are contained in the " +
				"scope zone of the elected bootstrap router.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyDisabled: PropDisabledRw,
		KeyInactive: PropInactiveRo,
		KeyName:     PropName("Name of the PIM-SM instance, referenced by the interface templates and the RP configuration."),
		"rp_hash_mask_length": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "The length of the mask used by the RP hash function that distributes the group ranges " +
				"between the candidate rendezvous points.",
			ValidateFunc:     validation.IntBetween(0, 128),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"rp_static_override": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "If enabled, the statically configured rendezvous points take precedence over the ones " +
				"learned from the bootstrap router.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"ssm_range": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The group range that is treated as Source Specific Multicast. No shared tree is built " +
				"for the groups inside this range.",
		},
		"switch_to_spt": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Allow the last hop router to switch from the RP shared tree to the shortest path tree of " +
				"the source.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"switch_to_spt_bytes": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "The amount of bytes that has to be received within `switch_to_spt_interval` before the " +
				"switch to the shortest path tree is performed.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"switch_to_spt_interval": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The measurement interval used together with `switch_to_spt_bytes`.",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEqual,
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
