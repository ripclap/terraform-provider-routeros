package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
  "channel": "stable",
  "check-certificate": "yes",
  "installed-version": "7.23",
  "ip-version": "auto",
  "mode": "https"
}
*/

// ResourceSystemPackageUpdate Settings of the RouterOS "check-for-updates" facility.
// https://help.mikrotik.com/docs/spaces/ROS/pages/40992872/Packages
func ResourceSystemPackageUpdate() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/system/package/update"),
		MetaId:           PropId(Id),

		"channel": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The release channel that is queried for the available RouterOS version." +
				"\n  * development;" +
				"\n  * long-term;" +
				"\n  * stable;" +
				"\n  * testing.",
			ValidateFunc:     validation.StringInSlice([]string{"development", "long-term", "stable", "testing"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"check_certificate": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Whether the certificate of the upgrade server is verified when `mode` is `https`." +
				"\n  * no;" +
				"\n  * yes;" +
				"\n  * yes-without-crl - verify the certificate but do not check the CRL.",
			ValidateFunc:     validation.StringInSlice([]string{"no", "yes", "yes-without-crl"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"installed_version": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The RouterOS version that is currently installed.",
		},
		"ip_version": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The IP protocol version used to contact the upgrade server." +
				"\n  * auto;" +
				"\n  * ipv4;" +
				"\n  * ipv6.",
			ValidateFunc:     validation.StringInSlice([]string{"auto", "ipv4", "ipv6"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"latest_version": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The latest RouterOS version available in the selected channel, filled in after a check.",
		},
		"mode": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The protocol used to reach the upgrade server." +
				"\n  * http;" +
				"\n  * https.",
			ValidateFunc:     validation.StringInSlice([]string{"http", "https"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"status": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The result of the last update check.",
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
