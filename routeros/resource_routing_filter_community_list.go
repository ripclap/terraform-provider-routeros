package routeros

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
  {
    ".id": "*1",
    "comment": "",
    "communities": "",
    "disabled": "false",
    "list": "",
    "regexp": ""
  }
*/

// ResourceRoutingFilterCommunityList A named set of BGP communities that can be referenced from the routing
// filter rules with the `community-list` matcher.
// https://help.mikrotik.com/docs/spaces/ROS/pages/74678285/Route+Selection+and+Filters
func ResourceRoutingFilterCommunityList() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/routing/filter/community-list"),
		MetaId:           PropId(Id),
		// 'about' is a console-only pseudo-property, it is never present in the REST output and must never be
		// written. The field is used as a base for the temporary 'place_before' exclusion during the update.
		MetaSkipFields: PropSkipFields("about"),

		KeyComment: PropCommentRw,
		"communities": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "A set of communities that belong to this list. Every element is a regular community in " +
				"the `asn:value` notation or one of the well-known community names (`accept-own`, `accept-own-nh`, " +
				"`no-advertise`, `no-export`, `no-export-subconfed`, `no-llgr`, `no-peer`).",
		},
		KeyDisabled: PropDisabledRw,
		"list": {
			Type:     schema.TypeString,
			Required: true,
			Description: "Name of the community list this entry belongs to. All the entries sharing the same name " +
				"form a single list that can be referenced from the routing filter rules.",
		},
		KeyPlaceBefore: PropPlaceBefore,
		"regexp": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "A regular expression that is matched against the community values instead of the exact " +
				"`communities` set.",
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
