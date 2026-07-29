package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
{
  ".id": "*7",
  "disabled": "true",
  "name": "ether1"
}
*/

// ResourceSystemResourceIrqRps Receive Packet Steering of a single interface.
// https://help.mikrotik.com/docs/spaces/ROS/pages/40992875/Resource
func ResourceSystemResourceIrqRps() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/system/resource/irq/rps"),
		MetaId:           PropId(Id),
		MetaSkipFields:   PropSkipFields("name"),

		KeyDisabled: PropDisabledRw,
		KeyName: PropName("The name of the interface whose Receive Packet Steering settings are managed. " +
			"The entry has to exist already; it is looked up by this name."),
	}

	return &schema.Resource{
		CreateContext: DefaultCreateUpdate(resSchema),
		ReadContext:   DefaultRead(resSchema),
		UpdateContext: DefaultCreateUpdate(resSchema),
		DeleteContext: DefaultSystemDelete(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: ImportStateCustomContext(resSchema),
		},

		Schema: resSchema,
	}
}
