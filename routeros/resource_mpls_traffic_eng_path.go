package routeros

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  The menu is empty on RouterOS 7.23 (ROS 7.23, MPLS not configured),
  the field set below is taken from `/console/inspect request=syntax path="mpls,traffic-eng,path,add"`
  and from `print proplist=` completion on that device.

  {
    ".id": "*1",
    "affinity-exclude": "0x00000000",
    "affinity-include-all": "0x00000000",
    "affinity-include-any": "0x00000000",
    "comment": "",
    "disabled": "false",
    "holding-priority": "7",
    "hops": "10.3.1.1/strict,10.1.1.2/strict",
    "name": "tun-1-link",
    "record-route": "true",
    "reoptimize-interval": "",
    "setup-priority": "7",
    "use-cspf": "false"
  }
*/

// ResourceMplsTrafficEngPath https://help.mikrotik.com/docs/spaces/ROS/pages/40992796/Traffic+Eng
func ResourceMplsTrafficEngPath() *schema.Resource {
	hexMask := validation.StringMatch(regexp.MustCompile(`^(0[xX][0-9a-fA-F]{1,8}|\d+)$`),
		"value must be a decimal number or a '0x'-prefixed hexadecimal number")

	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/mpls/traffic-eng/path"),
		MetaId:           PropId(Id),

		"affinity_exclude": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Do not use an interface if its `resource_class` matches any of the bits specified " +
				"here. Written as a decimal number or as a `0x`-prefixed hexadecimal number.",
			ValidateFunc:     hexMask,
			DiffSuppressFunc: HexEqual,
		},
		"affinity_include_all": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Use an interface only if its `resource_class` matches all of the bits specified here. " +
				"Written as a decimal number or as a `0x`-prefixed hexadecimal number.",
			ValidateFunc:     hexMask,
			DiffSuppressFunc: HexEqual,
		},
		"affinity_include_any": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Use an interface if its `resource_class` matches any of the bits specified here. " +
				"Written as a decimal number or as a `0x`-prefixed hexadecimal number.",
			ValidateFunc:     hexMask,
			DiffSuppressFunc: HexEqual,
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"holding_priority": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Priority used to decide whether the session set up over this path can be preempted by " +
				"another session. `0` is the highest priority.",
			ValidateFunc:     validation.IntBetween(0, 7),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"hops": {
			Type:     schema.TypeList,
			Optional: true,
			Description: "Explicit route of the path. Every element is an IPv4/IPv6 address optionally followed " +
				"by `/strict` or `/loose`, for example `10.3.1.1/strict`. The order is significant.",
			Elem: &schema.Schema{Type: schema.TypeString},
		},
		KeyName: PropName("Name of the tunnel path, referenced by the `primary_path` and `secondary_paths` " +
			"properties of the traffic engineering tunnels."),
		"record_route": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Whether the sender node asks to record the actual route the LSP traverses, which is " +
				"also used for the loop detection.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"reoptimize_interval": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Interval after which the path is re-optimized.",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEqual,
		},
		"setup_priority": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Priority used to decide whether the session set up over this path can preempt another " +
				"session. `0` is the highest priority.",
			ValidateFunc:     validation.IntBetween(0, 7),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"use_cspf": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Whether to use the CSPF (Constrained Shortest Path First) calculation to complete the " +
				"path instead of following the explicit route only.",
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
