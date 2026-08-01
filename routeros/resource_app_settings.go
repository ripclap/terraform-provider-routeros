package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
	{
	  "app-store-urls": "",
	  "assumed-download-path": "disk1/media/downloads",
	  "assumed-media-path": "disk1/media",
	  "auto-update": "true",
	  "certificate": "apps-*.example.com",
	  "certificate-status": "ok",
	  "disk": "disk1",
	  "lan-bridge": "bridge1",
	  "registry-mirrors": "",
	  "router-ip": "192.0.2.1",
	  "show-in-webfig": "true"
	}

	Sampled from RouterOS 7.23, `GET /rest/app/settings`
	(the certificate name has been redacted).

	`download-path` and `media-path` can be written but are not reported back; RouterOS returns the
	effective values in the read-only `assumed-download-path` and `assumed-media-path` properties.
*/

// ResourceAppSettings Global settings of the containerised application manager.
// https://help.mikrotik.com/docs/spaces/ROS/pages/343244823/Apps
func ResourceAppSettings() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/app/settings"),
		MetaId:           PropId(Id),

		"app_store_urls": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Comma separated list of custom app store URLs. Every URL must point to a YAML array in " +
				"which each element describes one application.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"assumed_download_path": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Download directory that is actually used, either taken from `download_path` or derived from `disk`.",
		},
		"assumed_media_path": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Media directory that is actually used, either taken from `media_path` or derived from `disk`.",
		},
		"auto_update": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Whether applications are updated automatically when a new container image is published.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"certificate": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Name of the certificate that is used for the HTTPS URLs of the installed applications.",
		},
		"certificate_status": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "State of the application certificate.",
		},
		"disk": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Name of the disk from `/disk` that stores the container images and the application data, " +
				"or `none`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"download_path": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Directory used for downloads by the applications. When it is not set RouterOS derives " +
				"one from `disk` and reports it in `assumed_download_path`." +
				"\n> RouterOS does not report this property back, so its value is kept from the configuration and " +
				"cannot be refreshed from the device.",
		},
		"lan_bridge": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Name of the bridge that represents the local network for applications attached to the " +
				"`lan` network, or `none`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"media_path": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Directory used for media files by the applications. When it is not set RouterOS derives " +
				"one from `disk` and reports it in `assumed_media_path`." +
				"\n> RouterOS does not report this property back, so its value is kept from the configuration and " +
				"cannot be refreshed from the device.",
		},
		"registry_mirrors": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Comma separated list of container registry mirrors used when pulling images.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"router_ip": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Address at which the router itself is reachable by the applications and by the users of " +
				"their web interfaces.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"show_in_webfig": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Whether the links to the enabled applications are shown on the WebFig login page.",
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
