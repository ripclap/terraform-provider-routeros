package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
	The menu was empty on the RouterOS 7.23:
	`GET /rest/tool/traffic-monitor` -> `[]`, so no value sample is available.

	Settable arguments, `/console/inspect request=syntax path="tool,traffic-monitor,add"`:
	  comment  copy-from  disabled  interface  name  on-event  threshold  traffic  trigger

	Readable properties, `/console/inspect request=completion
	input="/tool traffic-monitor print proplist="`, additionally report the read-only `invalid`
	property.
*/

// ResourceToolTrafficMonitor Runs a script when the traffic on an interface crosses a threshold.
// https://help.mikrotik.com/docs/spaces/ROS/pages/21725255/Tools
func ResourceToolTrafficMonitor() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/tool/traffic-monitor"),
		MetaId:           PropId(Id),

		KeyComment:   PropCommentRw,
		KeyDisabled:  PropDisabledRw,
		KeyInterface: PropInterfaceRw,
		KeyInvalid:   PropInvalidRo,
		KeyName:      PropName("Name of the traffic monitor entry."),
		"on_event": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Name of the script in `/system/script` that is executed when the condition described by " +
				"`traffic`, `trigger` and `threshold` is met.",
		},
		"threshold": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Traffic level, in bits per second (0..4294967295), the interface traffic is compared " +
				"against.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"traffic": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Direction of the interface traffic that is watched.",
			ValidateFunc:     validation.StringInSlice([]string{"received", "transmitted"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"trigger": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "When the script is executed:" +
				"\n  * above - when the traffic rises above the threshold," +
				"\n  * below - when the traffic drops below the threshold," +
				"\n  * always - on both events.",
			ValidateFunc:     validation.StringInSlice([]string{"above", "always", "below"}, false),
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
