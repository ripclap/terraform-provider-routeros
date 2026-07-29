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

// ResourceIpProxyCache Web proxy cache access list: decides which objects may be cached.
// https://help.mikrotik.com/docs/spaces/ROS/pages/132350000/Proxy
func ResourceIpProxyCache() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/proxy/cache"),
		MetaId:           PropId(Id),
		MetaSkipFields:   PropSkipFields("hits"),

		"action": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Whether the objects matched by the rule may be stored in the cache." +
				"\n  * allow - cache the matched objects." +
				"\n  * deny - do not cache the matched objects.",
			ValidateFunc:     validation.StringInSlice([]string{"allow", "deny"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"dst_address": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Matches the IP address of the destination (server) the request is sent to.",
		},
		"dst_host": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Matches the destination host name or IP address exactly as it was typed by the user. " +
				"Wildcards `*` and `?` are supported; a value starting with a colon is treated as a regular expression.",
		},
		"dst_port": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Matches a destination port, a list of ports or port ranges.",
		},
		"local_port": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Matches the proxy port the request was received on.",
		},
		"method": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Matches the HTTP method of the request. RouterOS stores and reports the method in upper " +
				"case, so the value has to be written in upper case to avoid a permanent difference.",
			ValidateFunc: validation.StringInSlice([]string{
				"CONNECT", "DELETE", "GET", "HEAD", "OPTIONS", "POST", "PUT", "TRACE",
			}, false),
		},
		"path": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Matches the path of the resource requested from the server. Wildcards `*` and `?` are " +
				"supported; a value starting with a colon is treated as a regular expression.",
		},
		KeyPlaceBefore: PropPlaceBefore,
		"src_address": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Matches the IP address of the client that originated the request.",
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
