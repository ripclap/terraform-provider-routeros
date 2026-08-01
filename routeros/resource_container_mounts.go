package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// https://help.mikrotik.com/docs/display/ROS/Container#Container-Addenvironmentvariablesandmounts(optional)
func ResourceContainerMounts() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/container/mounts"),
		MetaId:           PropId(Name),

		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"dst": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Specifies destination path of the mount, which points to defined location in container",
		},
		"list": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Name of the mount list the mount belongs to.",
		},
		"mode": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Access mode of the mount.",
		},
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of the mount.",
		},
		"src": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Specifies source path of the mount, which points to a RouterOS location",
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
