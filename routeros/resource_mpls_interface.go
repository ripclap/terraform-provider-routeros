package routeros

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  The menu is empty on the reference device (ROS 7.23.2, MPLS not configured),
  the field set below is taken from `/console/inspect request=syntax path="mpls,interface,add"`
  and from `print proplist=` completion on that device.

  {
    ".id": "*1",
    "builtin": "false",
    "comment": "",
    "disabled": "false",
    "input": "true",
    "interface": "all",
    "mpls-mtu": "1508"
  }
*/

// ResourceMplsInterface https://help.mikrotik.com/docs/spaces/ROS/pages/128974876/MPLS+MTU%2C+Forwarding+and+Label+Bindings
func ResourceMplsInterface() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/mpls/interface"),
		MetaId:           PropId(Id),
		// 'about' is a console-only pseudo column, it is never a part of the configuration.
		MetaSkipFields: PropSkipFields("about"),

		"builtin": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether this is the built-in (system) entry.",
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"input": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Whether to allow MPLS input (reception of labeled packets) on the matched interfaces.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"interface": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Name of the interface or of the interface list to match. The special value `all` matches " +
				"every interface. Entries are evaluated sequentially, the first matching entry is applied.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"mpls_mtu": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "How big packets can be carried over the interface with the MPLS labels added. It should be " +
				"at least the IP MTU plus 4 bytes for every label in the label stack.",
			ValidateFunc:     validation.IntBetween(512, 65535),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
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
