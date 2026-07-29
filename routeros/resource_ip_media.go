package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
	The menu was empty on the reference device (RouterOS 7.23.2), so no value sample is available.
*/

// ResourceIpMedia DLNA media server share.
// https://help.mikrotik.com/docs/spaces/ROS/pages/237699479/DLNA+Media+server
func ResourceIpMedia() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/media"),
		MetaId:           PropId(Id),

		"allowed_hostname": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Restricts access to this media server to the listed client hostnames.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"allowed_ip": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Restricts access to this media server to the listed client IP addresses.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyDisabled: PropDisabledRw,
		"friendly_name": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Name under which the DLNA server is advertised on the network.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyInterface: {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Network interface the DLNA server is served on.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"path": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Path on the router's file system where the media content is stored, e.g. `usb1/movies`.",
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
