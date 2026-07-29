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

// ResourceRoutingFilterCommunityLargeList A named set of BGP large communities (RFC 8092) that can be referenced
// from the routing filter rules with the `large-community-list` matcher.
// https://help.mikrotik.com/docs/spaces/ROS/pages/74678285/Route+Selection+and+Filters
func ResourceRoutingFilterCommunityLargeList() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/routing/filter/community-large-list"),
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
			Description: "A set of large communities that belong to this list. Every element is written in the " +
				"RFC 8092 three-part notation `global_administrator:local_data_part_1:local_data_part_2`.",
		},
		KeyDisabled: PropDisabledRw,
		"list": {
			Type:     schema.TypeString,
			Required: true,
			Description: "Name of the large community list this entry belongs to. All the entries sharing the same " +
				"name form a single list that can be referenced from the routing filter rules.",
		},
		KeyPlaceBefore: PropPlaceBefore,
		"regexp": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "A regular expression that is matched against the large community values instead of the " +
				"exact `communities` set.",
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
