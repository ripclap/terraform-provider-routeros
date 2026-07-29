package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
{
  ".id": "*1",
  "inactive": "false",
  "interface": "ether1",
  "ptp": "ptp1"
}
*/

// ResourceSystemPtpPort Attaches an interface to a PTP instance.
// https://help.mikrotik.com/docs/spaces/ROS/pages/64127015/Precision+Time+Protocol
func ResourceSystemPtpPort() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/system/ptp/port"),
		MetaId:           PropId(Id),

		KeyComment:   PropCommentRw,
		KeyInactive:  PropInactiveRo,
		KeyInterface: PropInterfaceRw,
		"ptp": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The name of the PTP instance (`/system/ptp`) this port belongs to.",
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
