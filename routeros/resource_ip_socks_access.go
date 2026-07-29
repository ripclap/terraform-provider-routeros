package routeros

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
	Menu is empty on the reference device (RouterOS 7.23.2), so no GET sample is available.
*/

// ResourceIpSocksAccess SOCKS proxy access list.
// https://help.mikrotik.com/docs/spaces/ROS/pages/73826308/SOCKS
func ResourceIpSocksAccess() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/socks/access"),
		MetaId:           PropId(Id),
		// 'about' is a console-only pseudo-property, it is never present in the REST output and must never be
		// written. The field is used as a base for the temporary 'place_before' exclusion during the update.
		MetaSkipFields: PropSkipFields("about"),

		"action": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Action taken for the packets matched by the rule." +
				"\n  * allow - allow the matched packets to be forwarded for further processing." +
				"\n  * deny - deny access to the matched packets.",
			ValidateFunc:     validation.StringInSlice([]string{"allow", "deny"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"dst_address": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Destination (server) address of the packet, optionally with a netmask.",
		},
		"dst_port": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Destination TCP port of the packet.",
		},
		KeyPlaceBefore: PropPlaceBefore,
		"src_address": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Source (client) address of the packet, optionally with a netmask.",
		},
		"src_port": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Source TCP port of the packet.",
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
