package routeros

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
  {
    ".id": "*4",
    "comment": "Private 16-bit RFC6996",
    "list": "bogon-asns",
    "range": "64512-65534"
  }
*/

// ResourceRoutingFilterNumList A named list of numbers or number ranges (AS numbers, metrics, distances, ...)
// that can be referenced from the routing filter rules and from the AS-path regular expressions with the
// `[[:list_name:]]` syntax.
// https://help.mikrotik.com/docs/spaces/ROS/pages/74678285/Route+Selection+and+Filters
func ResourceRoutingFilterNumList() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/routing/filter/num-list"),
		MetaId:           PropId(Id),
		// 'about' is a console-only pseudo-property, it is never present in the REST output and must never be
		// written. The field is used as a base for the temporary 'place_before' exclusion during the update.
		MetaSkipFields: PropSkipFields("about"),

		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"list": {
			Type:     schema.TypeString,
			Required: true,
			Description: "Name of the number list this entry belongs to. All the entries sharing the same name form " +
				"a single list that can be referenced from the routing filter rules.",
		},
		KeyPlaceBefore: PropPlaceBefore,
		"range": {
			Type:     schema.TypeString,
			Required: true,
			Description: "A single number or an inclusive range of numbers written as `first-last`, " +
				"for example `65535` or `64512-65534`.",
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
