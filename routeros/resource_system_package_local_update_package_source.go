package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
{
  ".id": "*1",
  "address": "192.0.2.1",
  "user": "admin"
}
*/

// ResourceSystemPackageLocalUpdatePackageSource A local package server that this device fetches
// RouterOS packages from during a local update.
// https://help.mikrotik.com/docs/spaces/ROS/pages/40992872/Packages
func ResourceSystemPackageLocalUpdatePackageSource() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/system/package/local-update/update-package-source"),
		MetaId:           PropId(Id),

		"address": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The IPv4 or IPv6 address of the local package server.",
		},
		"password": {
			Type:             schema.TypeString,
			Optional:         true,
			Sensitive:        true,
			Description:      "The password that is used to access the local package server.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"user": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The user name that is used to access the local package server.",
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
