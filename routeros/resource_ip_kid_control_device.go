package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
	The menu was empty on the reference device (RouterOS 7.23.2), so no value sample is available.
*/

// ResourceIpKidControlDevice Device attached to a Kid Control profile.
// https://help.mikrotik.com/docs/spaces/ROS/pages/129531911/Kid+Control
func ResourceIpKidControlDevice() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/kid-control/device"),
		MetaId:           PropId(Id),
		// Runtime traffic statistics and status flags, dropped on read.
		MetaSkipFields: PropSkipFields("blocked", "bytes_down", "bytes_up", "idle_time", "limited",
			"rate_down", "rate_up"),

		KeyDisabled: PropDisabledRw,
		KeyDynamic:  PropDynamicRo,
		KeyMacAddress: PropMacAddressRw("MAC address of the device that is being controlled.",
			true),
		KeyName: PropName("Name of the device."),
		"user": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Name of the Kid Control profile (`/ip/kid-control`) the device is appended to.",
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
