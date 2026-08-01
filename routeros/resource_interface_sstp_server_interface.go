package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
 {
	".id": "*1",
	"comment": "",
	"disabled": "false",
	"name": "sstp-in1",
	"running": "true",
	"user": "user1"
  }
*/

// ResourceInterfaceSstpServerInterface the interface a connected SSTP client is bound to.
// The server itself is /interface/sstp-server/server, modelled by ResourceInterfaceSstpServer.
// https://help.mikrotik.com/docs/display/ROS/SSTP
func ResourceInterfaceSstpServerInterface() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/sstp-server"),
		MetaId:           PropId(Id),

		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		KeyName:     PropName("Interface name (Example: sstp-in1)."),
		KeyRunning:  PropRunningRo,
		"user": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "User name used for authentication.",
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
