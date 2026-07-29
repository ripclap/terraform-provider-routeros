package routeros

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
  {
    ".id": "*1",
    "chain": "",
    "comment": "",
    "disabled": "false",
    "do-group-num": "",
    "do-group-prfx": "",
    "do-jump": "",
    "do-select-num": "",
    "do-select-prfx": "",
    "do-take": "",
    "do-where": "",
    "invalid": "false"
  }
*/

// ResourceRoutingFilterSelectRule A rule of a route select chain. Select chains are referenced from the
// `output.filter-select` / `out-filter-select` properties of BGP, OSPF, RIP and IS-IS and decide which of the
// candidate routes are advertised.
// https://help.mikrotik.com/docs/spaces/ROS/pages/74678285/Route+Selection+and+Filters
func ResourceRoutingFilterSelectRule() *schema.Resource {
	// The value grammar of the do-select-*/do-group-* properties was taken from the console completion engine of
	// RouterOS 7.23: a property token, optionally followed by `>` and one of the ordering selectors.
	const selectSyntax = "\n\nThe value is a property token, optionally followed by `>` and an ordering selector " +
		"(`largest-none-best`, `largest-none-worst`, `smallest-none-best`, `smallest-none-worst`)."

	const numProps = "\n\nNumeric property tokens reported by the device: `bgp-input-local-as`, " +
		"`bgp-input-remote-as`, `bgp-local-pref`, `bgp-med`, `bgp-out-med`, `bgp-output-local-as`, " +
		"`bgp-output-remote-as`, `bgp-path-len`, `bgp-path-peer-prepend`, `bgp-path-prepend`, `bgp-weight`, " +
		"`distance`, `dst-len`, `ospf-ext-metric`, `ospf-ext-tag`, `ospf-metric`, `ospf-tag`, `rip-ext-metric`, " +
		"`rip-ext-tag`, `rip-metric`, `rip-tag`, `scope`, `target-scope`."

	const prfxProps = "\n\nPrefix property tokens reported by the device: `bgp-input-local-addr`, " +
		"`bgp-input-remote-addr`, `bgp-output-local-addr`, `bgp-output-remote-addr`, `dst`, `gw`, `ospf-fwd`, " +
		"`pref-src`."

	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/routing/filter/select-rule"),
		MetaId:           PropId(Id),
		// 'about' is a console-only pseudo-property, it is never present in the REST output and must never be
		// written. The field is used as a base for the temporary 'place_before' exclusion during the update.
		MetaSkipFields: PropSkipFields("about"),

		"chain": {
			Type:     schema.TypeString,
			Required: true,
			Description: "Name of the select chain this rule belongs to. If the name does not match an already " +
				"defined chain, a new chain is created.",
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"do_group_num": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Group the candidate routes by the value of a numeric route property." + selectSyntax + numProps,
		},
		"do_group_prfx": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Group the candidate routes by the value of a prefix/address route property." + selectSyntax + prfxProps,
		},
		"do_jump": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Name of the select chain to jump to.",
		},
		"do_select_num": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Select the candidate routes by the value of a numeric route property." + selectSyntax + numProps,
		},
		"do_select_prfx": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Select the candidate routes by the value of a prefix/address route property." + selectSyntax + prfxProps,
		},
		"do_take": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "VERIFY: limits how many of the remaining candidate routes are taken. The console exposes " +
				"no value domain for this property and it is not covered by the MikroTik documentation - supply the " +
				"raw RouterOS value.",
		},
		"do_where": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Name of a `/routing/filter/rule` chain that is called to decide which of the candidate " +
				"routes stay in the selection.",
		},
		KeyInvalid:     PropInvalidRo,
		KeyPlaceBefore: PropPlaceBefore,
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
