package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
	{
	  "enabled": "false",
	  "id": "00:00:00:00:00:00",
	  "secrets": ""
	}
*/

// ResourceToolRomon RoMON (Router Management Overlay Network) global settings.
// https://help.mikrotik.com/docs/spaces/ROS/pages/8978569/RoMON
func ResourceToolRomon() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/tool/romon"),
		MetaId:           PropId(Id),
		MetaTransformSet: PropTransformSet("romon_id: id"),

		KeyEnabled: PropEnabled("Whether the RoMON agent is enabled on this device."),
		"romon_id": PropMacAddressRw("Individual RoMON ID of the device, sent on the wire as the RouterOS `id` "+
			"property. When left at `00:00:00:00:00:00` the lowest MAC address of the device is used.", false),
		"secrets": {
			Type:      schema.TypeString,
			Optional:  true,
			Sensitive: true,
			Description: "Comma separated list of secrets used for RoMON message authentication, integrity check " +
				"and replay prevention. Per-port secrets can be set in `/tool/romon/port`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
	}

	return &schema.Resource{
		CreateContext: DefaultSystemCreate(resSchema),
		ReadContext:   DefaultSystemRead(resSchema),
		UpdateContext: DefaultSystemUpdate(resSchema),
		DeleteContext: DefaultSystemDelete(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resSchema,
	}
}
