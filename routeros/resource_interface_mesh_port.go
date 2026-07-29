package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
	The menu was empty on the reference device (RouterOS 7.23.2):
	`GET /rest/interface/mesh/port` -> `[]`, so no value sample is available.

	Writable fields reported by the device:
	/console/inspect request=child path="interface,mesh,port,add"
	  comment  copy-from  disabled  hello-interval  interface  mesh  path-cost  port-type

	Read-only fields reported by the same device:
	  active-port-type  dr-address  dynamic  inactive

	VERIFY: HWMP+ mesh has no page on help.mikrotik.com; the property descriptions below come from
	the legacy MikroTik HWMP+ mesh manual. The field names, the value types and the enumerations
	were taken from the device.
*/

// ResourceInterfaceMeshPort Ports of an HWMP+ mesh interface.
func ResourceInterfaceMeshPort() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/mesh/port"),
		MetaId:           PropId(Id),

		"active_port_type": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The port type the mesh has actually selected for this port.",
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"dr_address": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "MAC address of the designated router (mesh portal) reachable through this port.",
		},
		KeyDynamic: PropDynamicRo,
		"hello_interval": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "How often to send out the mesh hello messages on this port.",
			DiffSuppressFunc: TimeEqual,
		},
		KeyInactive: PropInactiveRo,
		KeyInterface: {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The interface that is added to the mesh.",
		},
		"mesh": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The mesh interface the port belongs to.",
		},
		"path_cost": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Path cost of this port, used when the mesh selects the best path to a destination.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"port_type": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "How the mesh treats this port. `auto` derives the type from the underlying interface, " +
				"the other values force it.",
			ValidateFunc:     validation.StringInSlice([]string{"WDS", "auto", "ethernet", "wireless"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
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
