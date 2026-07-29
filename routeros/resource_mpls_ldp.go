package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  The menu is empty on the reference device (ROS 7.23.2, MPLS not configured),
  the field set below is taken from `/console/inspect request=syntax path="mpls,ldp,add"` and
  from `print proplist=` completion on that device.

  {
    ".id": "*1",
    "afi": "ip",
    "comment": "",
    "disabled": "false",
    "distribute-for-default": "false",
    "hop-limit": "255",
    "inactive": "false",
    "loop-detect": "false",
    "lsr-id": "10.0.0.1",
    "path-vector-limit": "255",
    "preferred-afi": "ip",
    "transport-addresses": "10.0.0.1",
    "use-explicit-null": "false",
    "vrf": "main"
  }
*/

// ResourceMplsLdp https://help.mikrotik.com/docs/spaces/ROS/pages/121995275/LDP
func ResourceMplsLdp() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/mpls/ldp"),
		MetaId:           PropId(Id),

		"afi": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "Address families for which this LDP instance distributes labels.",
			Elem: &schema.Schema{
				Type:             schema.TypeString,
				ValidateDiagFunc: ValidationValInSlice([]string{"ip", "ipv6"}, false, false),
			},
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"distribute_for_default": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Whether to advertise a label for the default route (0.0.0.0/0 or ::/0).",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"hop_limit": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Maximum number of hops an LSP may traverse, used by the LDP loop detection.",
			ValidateFunc:     validation.IntBetween(0, 255),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyInactive: PropInactiveRo,
		"loop_detect": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Whether to enable the LDP loop detection (hop count and path vector).",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"lsr_id": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Label Switch Router ID of this LDP instance.",
			ValidateFunc:     validation.IsIPv4Address,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"path_vector_limit": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Maximum number of LSRs allowed in the path vector used by the LDP loop detection.",
			ValidateFunc:     validation.IntBetween(0, 255),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"preferred_afi": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Address family preferred for the transport connection when both are available.",
			ValidateFunc:     validation.StringInSlice([]string{"ip", "ipv6"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"transport_addresses": {
			Type:     schema.TypeList,
			Optional: true,
			Description: "List of the IPv4/IPv6 addresses advertised as the LDP transport addresses of this " +
				"instance. The order is significant, the first usable address of the preferred address family is used.",
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPAddress,
			},
		},
		"use_explicit_null": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Whether to advertise the explicit null label instead of the implicit null label to the " +
				"penultimate hop.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
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
