package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/* {"authorization":"false"} */

// ResourceSystemResourceHardwareUsbSettings USB device authorization settings.
// https://help.mikrotik.com/docs/spaces/ROS/pages/40992875/Resource
func ResourceSystemResourceHardwareUsbSettings() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/system/resource/hardware/usb-settings"),
		MetaId:           PropId(Id),

		"authorization": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Whether a connected USB device has to be authorized before it is allowed to operate. " +
				"When enabled, newly attached devices stay unauthorized until they are explicitly allowed.",
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
