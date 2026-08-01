package routeros

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
  The menu is empty on RouterOS 7.23 (ROS 7.23, MPLS not configured),
  the field set below is taken from `/console/inspect request=syntax path="mpls,ldp,advertise-filter,add"`
  and from `print proplist=` completion on that device.

  {
    ".id": "*1",
    "advertise": "true",
    "comment": "",
    "disabled": "false",
    "neighbor": "10.0.0.2",
    "prefix": "10.10.0.0/24",
    "vrf": "main"
  }
*/

// ResourceMplsLdpAdvertiseFilter https://help.mikrotik.com/docs/spaces/ROS/pages/121995275/LDP
func ResourceMplsLdpAdvertiseFilter() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/mpls/ldp/advertise-filter"),
		MetaId:           PropId(Id),
		// 'about' is a console-only pseudo column, it is never a part of the configuration.
		MetaSkipFields: PropSkipFields("about"),

		"advertise": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Whether to advertise the label mappings matched by this rule.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"neighbor": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Neighbor the label mapping would be advertised to, written as an IPv4/IPv6 address or " +
				"as an `address/prefix` range. An empty value matches every neighbor.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyPlaceBefore: PropPlaceBefore,
		"prefix": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Prefix of the label mapping to match, written as `address/prefix-length`. An empty " +
				"value matches every prefix.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"vrf": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Name of the VRF this rule applies to. The special value `any` applies the rule to " +
				"every VRF.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
	}

	return &schema.Resource{
		CreateContext: DefaultCreate(resSchema),
		ReadContext:   DefaultRead(resSchema),
		UpdateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			skip := resSchema[MetaSkipFields].Default.(string)
			resSchema[MetaSkipFields].Default = skip + `,"place_before"`
			defer func() {
				resSchema[MetaSkipFields].Default = skip
			}()

			return ResourceUpdate(ctx, resSchema, d, m)
		},
		DeleteContext: DefaultDelete(resSchema),
		Importer: &schema.ResourceImporter{
			StateContext: ImportStateCustomContext(resSchema),
		},

		Schema: resSchema,
	}
}
