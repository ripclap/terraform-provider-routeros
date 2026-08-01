package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
	The menu was empty on the RouterOS 7.23, so no value sample is available.
*/

// https://help.mikrotik.com/docs/spaces/ROS/pages/295239888/File+share
func ResourceIpCloudBackToHomeFile() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/cloud/back-to-home-file"),
		MetaId:           PropId(Id),
		// Read-only share statistics and secrets, dropped rather than exposed.
		MetaSkipFields: PropSkipFields("downloads", "key"),

		"allow_uploads": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Enables the option for anyone holding the share link to upload files to the router. " +
				"Applicable to shared directories.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"direct_url": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Direct download link of the share (the `url` with the `?dl` parameter appended).",
		},
		"expires": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Date and time when the share expires, in the ISO 8601 `YYYY-MM-DD HH:MM:SS` form " +
				"(for example `2026-01-25 00:00:00`).",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"path": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Path of the file or of the directory on the router that is being shared.",
		},
		"url": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Generated HTTPS link of the share.",
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
