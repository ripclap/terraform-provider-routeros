package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
	The menu was empty on the RouterOS 7.23:
	`GET /rest/tool/traffic-generator/port` -> `[]`, so no value sample is available.

	Settable arguments, `/console/inspect request=syntax path="tool,traffic-generator,port,add"`:
	  copy-from  disabled  interface  name

	Readable properties, `/console/inspect request=completion
	input="/tool traffic-generator port print proplist="`:
	  about  disabled  dynamic  first-header  interface  invalid  name
*/

// ResourceToolTrafficGeneratorPort Traffic generator transmit/receive ports.
// https://help.mikrotik.com/docs/spaces/ROS/pages/128221376/Traffic+Generator
func ResourceToolTrafficGeneratorPort() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/tool/traffic-generator/port"),
		MetaId:           PropId(Id),

		KeyDisabled: PropDisabledRw,
		KeyDynamic:  PropDynamicRo,
		"first_header": {
			Type:     schema.TypeString,
			Computed: true,
			Description: "Header type that the traffic generator assumes for packets received on this port. " +
				"Derived by RouterOS from the interface type.",
		},
		KeyInterface: PropInterfaceRw,
		KeyInvalid:   PropInvalidRo,
		KeyName:      PropName("Name of the traffic generator port, referenced by packet templates and streams."),
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
