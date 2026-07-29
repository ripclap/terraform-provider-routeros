package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
{
  ".id": "*1",
  "active": "true",
  "fallback-to": "next",
  "name": "part0",
  "running": "true",
  "size": "128",
  "version": "RouterOS v7.23.2 2026-07-03 09:08:08"
}
*/

// ResourcePartitions A RouterOS partition of the internal storage. Partitions are created by
// `/partitions/repartition` and cannot be added or removed, so this resource looks one up by `name`
// and only manages its attributes.
// https://help.mikrotik.com/docs/spaces/ROS/pages/328103/Partitions
func ResourcePartitions() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/partitions"),
		MetaId:           PropId(Id),
		MetaSkipFields:   PropSkipFields("name"),

		"active": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether the partition is the one that is booted by default.",
		},
		KeyComment: PropCommentRw,
		"fallback_to": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The partition that is booted if this one fails to boot: the name of another partition, " +
				"`next` for the next partition in order, or `etherboot` to fall back to network boot.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyName: PropName("The name of the partition, for example `part0`. The partition has to exist already; " +
			"it is looked up by this name."),
		KeyRunning: PropRunningRo,
		"size": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The size of the partition, in MiB.",
		},
		"version": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The RouterOS version installed on the partition.",
		},
	}

	return &schema.Resource{
		CreateContext: DefaultCreateUpdate(resSchema),
		ReadContext:   DefaultRead(resSchema),
		UpdateContext: DefaultCreateUpdate(resSchema),
		DeleteContext: DefaultSystemDelete(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: ImportStateCustomContext(resSchema),
		},

		Schema: resSchema,
	}
}
