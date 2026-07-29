package routeros

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
  {
    ".id": "*1",
    "bcast.l1.csnp-interval": "",
    "bcast.l1.hello-interval": "",
    "bcast.l1.hello-interval-dr": "",
    "bcast.l1.hello-multiplier": "",
    "bcast.l1.metric": "",
    "bcast.l1.priority": "",
    "bcast.l1.psnp-interval": "",
    "bcast.l2.csnp-interval": "",
    "bcast.l2.hello-interval": "",
    "bcast.l2.hello-interval-dr": "",
    "bcast.l2.hello-multiplier": "",
    "bcast.l2.metric": "",
    "bcast.l2.priority": "",
    "bcast.l2.psnp-interval": "",
    "comment": "",
    "disabled": "false",
    "inactive": "false",
    "instance": "isis-instance-1",
    "interfaces": "ether1",
    "levels": "l1,l2",
    "passive": "",
    "ptp": "",
    "ptp.hello-3way": "",
    "ptp.hello-interval": "",
    "ptp.hello-multiplier": "",
    "ptp.l1.csnp-interval": "",
    "ptp.l1.metric": "",
    "ptp.l1.psnp-interval": "",
    "ptp.l2.csnp-interval": "",
    "ptp.l2.metric": "",
    "ptp.l2.psnp-interval": ""
  }
*/

// ResourceRoutingIsisInterfaceTemplate https://help.mikrotik.com/docs/spaces/ROS/pages/201523543/IS-IS
func ResourceRoutingIsisInterfaceTemplate() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/routing/isis/interface-template"),
		MetaId:           PropId(Id),
		MetaTransformSet: PropTransformSet(
			"bcast_l1_csnp_interval: bcast.l1.csnp-interval",
			"bcast_l1_hello_interval: bcast.l1.hello-interval",
			"bcast_l1_hello_interval_dr: bcast.l1.hello-interval-dr",
			"bcast_l1_hello_multiplier: bcast.l1.hello-multiplier",
			"bcast_l1_metric: bcast.l1.metric",
			"bcast_l1_priority: bcast.l1.priority",
			"bcast_l1_psnp_interval: bcast.l1.psnp-interval",
			"bcast_l2_csnp_interval: bcast.l2.csnp-interval",
			"bcast_l2_hello_interval: bcast.l2.hello-interval",
			"bcast_l2_hello_interval_dr: bcast.l2.hello-interval-dr",
			"bcast_l2_hello_multiplier: bcast.l2.hello-multiplier",
			"bcast_l2_metric: bcast.l2.metric",
			"bcast_l2_priority: bcast.l2.priority",
			"bcast_l2_psnp_interval: bcast.l2.psnp-interval",
			"ptp_hello_3way: ptp.hello-3way",
			"ptp_hello_interval: ptp.hello-interval",
			"ptp_hello_multiplier: ptp.hello-multiplier",
			"ptp_l1_csnp_interval: ptp.l1.csnp-interval",
			"ptp_l1_metric: ptp.l1.metric",
			"ptp_l1_psnp_interval: ptp.l1.psnp-interval",
			"ptp_l2_csnp_interval: ptp.l2.csnp-interval",
			"ptp_l2_metric: ptp.l2.metric",
			"ptp_l2_psnp_interval: ptp.l2.psnp-interval",
		),
		// 'about' is console-only, never in REST output and must not be written; also the base for the temporary 'place_before' update exclusion.
		MetaSkipFields: PropSkipFields("about"),
		// Flag properties: RouterOS reports an empty value when set and omits them when unset; accepted without a value.
		MetaSetUnsetFields: PropSetUnsetFields("passive", "ptp", "ptp_hello_3way"),

		"bcast_l1_csnp_interval": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The interval between the complete sequence number PDUs sent by the designated " +
				"intermediate system on a level 1 broadcast circuit.",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEqual,
		},
		"bcast_l1_hello_interval": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The interval between the IIH PDUs sent on a level 1 broadcast circuit.",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEqual,
		},
		"bcast_l1_hello_interval_dr": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The interval between the IIH PDUs sent on a level 1 broadcast circuit while this router " +
				"is the designated intermediate system. It is normally a third of `bcast_l1_hello_interval`.",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEqual,
		},
		"bcast_l1_hello_multiplier": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "The number of missed hellos after which a level 1 broadcast adjacency is declared down. " +
				"The advertised holding time is the hello interval multiplied by this value.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"bcast_l1_metric": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "The level 1 metric of a broadcast circuit.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"bcast_l1_priority": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "The priority used in the election of the level 1 designated intermediate system on a " +
				"broadcast circuit. The highest priority wins.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"bcast_l1_psnp_interval": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The minimum interval between the partial sequence number PDUs sent on a level 1 " +
				"broadcast circuit.",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEqual,
		},
		"bcast_l2_csnp_interval": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The interval between the complete sequence number PDUs sent by the designated " +
				"intermediate system on a level 2 broadcast circuit.",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEqual,
		},
		"bcast_l2_hello_interval": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The interval between the IIH PDUs sent on a level 2 broadcast circuit.",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEqual,
		},
		"bcast_l2_hello_interval_dr": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The interval between the IIH PDUs sent on a level 2 broadcast circuit while this router " +
				"is the designated intermediate system. It is normally a third of `bcast_l2_hello_interval`.",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEqual,
		},
		"bcast_l2_hello_multiplier": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "The number of missed hellos after which a level 2 broadcast adjacency is declared down. " +
				"The advertised holding time is the hello interval multiplied by this value.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"bcast_l2_metric": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "The level 2 metric of a broadcast circuit.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"bcast_l2_priority": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "The priority used in the election of the level 2 designated intermediate system on a " +
				"broadcast circuit. The highest priority wins.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"bcast_l2_psnp_interval": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The minimum interval between the partial sequence number PDUs sent on a level 2 " +
				"broadcast circuit.",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEqual,
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		KeyInactive: PropInactiveRo,
		"instance": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of the IS-IS instance the matching interfaces are attached to.",
		},
		"interfaces": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "Interfaces to match. Both interface names and interface list names are accepted.",
		},
		"levels": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type:             schema.TypeString,
				ValidateDiagFunc: ValidationMultiValInSlice([]string{"l1", "l2"}, false, false),
			},
			Description: "The IS-IS levels the matching interfaces form adjacencies on.",
		},
		"passive": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
			Description: "If enabled, then do not send or receive IS-IS traffic on the matching interfaces, but " +
				"still advertise the attached networks." +
				"\n<em>This property is a flag: the correct value may not be displayed in Winbox, check it in the console.</em>",
		},
		KeyPlaceBefore: PropPlaceBefore,
		"ptp": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
			Description: "Treat the matching interfaces as point to point circuits instead of broadcast circuits. " +
				"The `ptp_*` properties are the ones that apply then." +
				"\n<em>This property is a flag: the correct value may not be displayed in Winbox, check it in the console.</em>",
		},
		"ptp_hello_3way": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
			Description: "Use the three way handshake of RFC 5303 when bringing up a point to point adjacency." +
				"\n<em>This property is a flag: the correct value may not be displayed in Winbox, check it in the console.</em>",
		},
		"ptp_hello_interval": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The interval between the IIH PDUs sent on a point to point circuit.",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEqual,
		},
		"ptp_hello_multiplier": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "The number of missed hellos after which a point to point adjacency is declared down. " +
				"The advertised holding time is the hello interval multiplied by this value.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"ptp_l1_csnp_interval": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The interval between the complete sequence number PDUs sent on a level 1 point to point " +
				"circuit.",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEqual,
		},
		"ptp_l1_metric": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "The level 1 metric of a point to point circuit.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"ptp_l1_psnp_interval": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The minimum interval between the partial sequence number PDUs sent on a level 1 point to " +
				"point circuit.",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEqual,
		},
		"ptp_l2_csnp_interval": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The interval between the complete sequence number PDUs sent on a level 2 point to point " +
				"circuit.",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEqual,
		},
		"ptp_l2_metric": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "The level 2 metric of a point to point circuit.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"ptp_l2_psnp_interval": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The minimum interval between the partial sequence number PDUs sent on a level 2 point to " +
				"point circuit.",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEqual,
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
