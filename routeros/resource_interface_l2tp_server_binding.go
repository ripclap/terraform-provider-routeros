package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
	The menu was empty on the RouterOS 7.23:
	`GET /rest/interface/l2tp-server` -> `[]`, so no value sample is available.

	Writable fields reported by the device:
	/console/inspect request=child path="interface,l2tp-server,add"
	  comment  copy-from  disabled  name  user

	Read-only fields reported by the same device:
	  client-address  dynamic  encoding  mru  mtu  running  uptime
*/

// ResourceInterfaceL2tpServerBinding A static L2TP server binding interface.
// The L2TP server itself is configured with `routeros_interface_l2tp_server`.
// https://help.mikrotik.com/docs/spaces/ROS/pages/2031631/L2TP
func ResourceInterfaceL2tpServerBinding() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/l2tp-server"),
		MetaId:           PropId(Id),

		"client_address": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The IP address of the currently connected client.",
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		KeyDynamic:  PropDynamicRo,
		"encoding": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Encryption and encoding used by the active session.",
		},
		"mru": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "Maximum Receive Unit negotiated for the active session.",
		},
		"mtu": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "Maximum Transmission Unit negotiated for the active session.",
		},
		KeyName:    PropName("Name of the interface."),
		KeyRunning: PropRunningRo,
		"uptime": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Duration of the active session.",
		},
		"user": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The name of the PPP user the interface is bound to.",
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
