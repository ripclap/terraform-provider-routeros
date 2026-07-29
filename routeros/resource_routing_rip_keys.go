package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
  {
    ".id": "*1",
    "chain": "",
    "disabled": "false",
    "key": "",
    "key-id": "",
    "valid-from": "",
    "valid-till": ""
  }
*/

// ResourceRoutingRipKeys https://help.mikrotik.com/docs/spaces/ROS/pages/328211/RIP
func ResourceRoutingRipKeys() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/routing/rip/keys"),
		MetaId:           PropId(Id),

		"chain": {
			Type:     schema.TypeString,
			Required: true,
			Description: "Name of the key chain this key belongs to. The chain is referenced from the `key_chain` " +
				"property of a RIP interface template.",
		},
		KeyDisabled: PropDisabledRw,
		"key": {
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   true,
			Description: "The secret used to authenticate the RIP messages.",
		},
		"key_id": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "The key identifier that is carried in the RIP messages so that the receiver knows which " +
				"key of the chain to use.",
		},
		"valid_from": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The moment the key becomes usable, in the RouterOS date and time notation, for example " +
				"`jan/01/2026 00:00:00`.",
		},
		"valid_till": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The moment the key stops being usable, in the RouterOS date and time notation, for " +
				"example `jan/01/2027 00:00:00`.",
		},
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
