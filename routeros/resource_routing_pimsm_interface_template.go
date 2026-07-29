package routeros

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
  {
    ".id": "*1",
    "disabled": "false",
    "hello-delay": "",
    "hello-period": "",
    "inactive": "false",
    "instance": "",
    "interfaces": "",
    "join-prune-period": "",
    "join-tracking-support": "false",
    "override-interval": "",
    "priority": "",
    "propagation-delay": "",
    "source-addresses": ""
  }
*/

// ResourceRoutingPimsmInterfaceTemplate https://help.mikrotik.com/docs/spaces/ROS/pages/61767728/PIM-SM
func ResourceRoutingPimsmInterfaceTemplate() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/routing/pimsm/interface-template"),
		MetaId:           PropId(Id),
		// 'about' is console-only, never in REST output and must not be written; also the base for the temporary 'place_before' update exclusion.
		MetaSkipFields: PropSkipFields("about"),

		KeyDisabled: PropDisabledRw,
		"hello_delay": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The maximum randomized delay before the first PIM Hello is sent on a newly enabled " +
				"interface.",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEqual,
		},
		"hello_period": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The interval between the periodic PIM Hello messages.",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEqual,
		},
		KeyInactive: PropInactiveRo,
		"instance": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of the PIM-SM instance the matching interfaces are attached to.",
		},
		"interfaces": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "Interfaces to match. Both interface names and interface list names are accepted.",
		},
		"join_prune_period": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The interval between the periodic PIM Join/Prune messages.",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEqual,
		},
		"join_tracking_support": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Announce the support of the join attribute tracking, which lets the upstream router " +
				"suppress the prune override behaviour.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"override_interval": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The value of the Override Interval advertised in the PIM Hello messages. It bounds the " +
				"randomized delay before a Join is sent to override someone else's Prune.",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEqual,
		},
		"priority": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "The Designated Router priority advertised in the PIM Hello messages.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyPlaceBefore: PropPlaceBefore,
		"propagation_delay": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The value of the LAN Propagation Delay advertised in the PIM Hello messages, the expected " +
				"message propagation delay on the link.",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEqual,
		},
		"source_addresses": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "The local addresses used as the source of the PIM messages sent on the matching " +
				"interfaces.",
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
