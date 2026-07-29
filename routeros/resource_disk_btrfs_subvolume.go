package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
{
  ".id": "*1",
  "creation-time": "2026-07-01 12:00:00",
  "dead": "false",
  "default": "false",
  "fs": "BtrfsDisk",
  "fullname": "BtrfsDisk/data",
  "generation": "42",
  "mount": "true",
  "mounted": "true",
  "mountpoint": "data",
  "name": "data",
  "read-only": "false",
  "snapshot": "false",
  "subvolume-id": "256",
  "top-level": "5",
  "uuid": "00000000-0000-0000-0000-000000000000"
}
*/

// ResourceDiskBtrfsSubvolume A Btrfs subvolume or snapshot on a Btrfs file system created with
// `/disk/format-drive file-system=btrfs`.
// https://help.mikrotik.com/docs/spaces/ROS/pages/295239711/Btrfs
func ResourceDiskBtrfsSubvolume() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/disk/btrfs/subvolume"),
		MetaId:           PropId(Id),

		"creation_time": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The moment the subvolume was created.",
		},
		"dead": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether the subvolume has been deleted and is only waiting to be cleaned up.",
		},
		KeyDefault: PropDefaultRo,
		"fs": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The label of the Btrfs file system that holds this subvolume.",
		},
		"fullname": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The full path of the subvolume inside the Btrfs file system.",
		},
		"generation": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The Btrfs generation the subvolume was last modified in.",
		},
		"mount": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Whether the subvolume is mounted automatically.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"mounted": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether the subvolume is currently mounted.",
		},
		"mountpoint": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The path the subvolume is made available at when `mount` is enabled.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyName: PropName("Name of the subvolume."),
		"parent": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The parent subvolume. Setting it creates a snapshot of that subvolume instead of an " +
				"empty subvolume.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"path": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The path of the subvolume relative to the top level of the Btrfs file system.",
		},
		"read_only": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Makes the subvolume immutable. Has to be set to `yes` on the source subvolume when " +
				"creating snapshots that are going to be sent to another device.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"received_uuid": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The UUID of the source subvolume when this subvolume was created by a receive transfer.",
		},
		"recv_time": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The moment the subvolume was received from another device.",
		},
		"recv_trans_id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The Btrfs transaction id of the receive operation.",
		},
		"send_time": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The moment the subvolume was last sent to another device.",
		},
		"send_trans_id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The Btrfs transaction id of the send operation.",
		},
		"snapshot": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether this subvolume is a snapshot of another subvolume.",
		},
		"snapshots": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The snapshots taken from this subvolume.",
		},
		"subvolume_id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The Btrfs internal subvolume id.",
		},
		"top_level": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The id of the subvolume this subvolume is nested in.",
		},
		"uuid": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The UUID of the subvolume.",
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
