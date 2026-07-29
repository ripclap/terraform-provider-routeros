package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
  {
    ".id": "*3",
    "address": "2606:4700:60::2",
    "disabled": "false",
    "expire-interval": "7200",
    "group": "cloudflare",
    "port": "8282",
    "preference": "1",
    "refresh-interval": "3600",
    "retry-interval": "600",
    "vrf": "main"
  }
*/

// ResourceRoutingRpki https://help.mikrotik.com/docs/spaces/ROS/pages/59277471/RPKI
func ResourceRoutingRpki() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/routing/rpki"),
		MetaId:           PropId(Id),

		"address": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The address of the RPKI validator the RTR session is established with.",
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"expire_interval": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The time the cached data stays usable after the connection to the validator is lost. " +
				"When it elapses, the records of this session are dropped.",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEqual,
		},
		"group": {
			Type:     schema.TypeString,
			Required: true,
			Description: "Name of the RPKI group this session belongs to. The group name is what the routing " +
				"filter `rpki-verify` matcher refers to, and several sessions can back a single group.",
		},
		"port": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "The TCP port of the RTR service on the validator.",
			ValidateFunc:     Validation64k,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"preference": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "The preference of this session inside the group, used when several sessions back the same " +
				"group. VERIFY: the direction of the ordering could not be established - the console exposes no help " +
				"text for this property and the MikroTik RPKI page does not document it.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"refresh_interval": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The interval between the serial queries sent to the validator to refresh the data.",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEqual,
		},
		"retry_interval": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The time to wait before retrying a failed connection to the validator.",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEqual,
		},
		KeyVrf: PropVrfRw,
	}

	return &schema.Resource{
		CreateContext: DefaultCreate(resSchema),
		ReadContext:   DefaultRead(resSchema),
		UpdateContext: DefaultUpdate(resSchema),
		DeleteContext: DefaultDelete(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: ImportStateCustomContext(resSchema),
		},

		Schema: resSchema,
	}
}
