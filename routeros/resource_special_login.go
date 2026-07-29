package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
	The menu was empty on the reference device (RouterOS 7.23.2):
	`GET /rest/special-login` -> `[]`, so no value sample is available.

	Settable arguments, `/console/inspect request=syntax path="special-login,add"`:
	  channel  copy-from  disabled  port  user

	Readable properties, `/console/inspect request=completion
	input="/special-login print proplist="`:
	  about  channel  disabled  port  user
*/

// ResourceSpecialLogin https://help.mikrotik.com/docs/spaces/ROS/pages/8978525/Ports
func ResourceSpecialLogin() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/special-login"),
		MetaId:           PropId(Id),

		"channel": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Channel of the serial port the login is bound to.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyDisabled: PropDisabledRw,
		"port": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of the `/port` entry the user is logged in on, for example `serial0`.",
		},
		"user": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of the `/user` account that is logged in on the port.",
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
