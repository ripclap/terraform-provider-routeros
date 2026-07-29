package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
{
  "check-interval": "1d",
  "enabled": "false",
  "primary-server": "::",
  "secondary-server": "::",
  "software-id": "XXXX-XXXX",
  "user": ""
}
*/

// ResourceSystemPackageLocalUpdateMirror Mirrors RouterOS packages for all architectures from a local
// package server so that other devices in the network can be upgraded from this device.
// https://help.mikrotik.com/docs/spaces/ROS/pages/40992872/Packages
func ResourceSystemPackageLocalUpdateMirror() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/system/package/local-update/mirror"),
		MetaId:           PropId(Id),

		"check_interval": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The time interval at which the device checks the local package server.",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEqual,
		},
		KeyEnabled: PropEnabled("Whether the periodic check of the local package server is enabled."),
		"password": {
			Type:             schema.TypeString,
			Optional:         true,
			Sensitive:        true,
			Description:      "The password that is used to access the local package server.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"primary_server": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The IPv4 or IPv6 address of the primary local package server.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"secondary_server": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The IPv4 or IPv6 address of the secondary local package server.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"software_id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The software ID of this device.",
		},
		"user": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The user name that is used to access the local package server.",
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
