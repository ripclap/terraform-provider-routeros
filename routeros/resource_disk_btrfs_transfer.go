package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
  ".id": "*1",
  "bytes": "0",
  "fs": "BtrfsDisk",
  "send-subvolumes": "snapshot1",
  "ssh-address": "192.0.2.2",
  "ssh-port": "22",
  "ssh-receive-mount": "BackupBtrfsDisk/Snapshots",
  "ssh-user": "btrfs",
  "status": "idle",
  "type": "send"
}
*/

// ResourceDiskBtrfsTransfer A Btrfs send/receive transfer of subvolumes between two devices. The menu has
// no `set`, so every configurable attribute is ForceNew.
// https://help.mikrotik.com/docs/spaces/ROS/pages/295239711/Btrfs
func ResourceDiskBtrfsTransfer() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/disk/btrfs/transfer"),
		MetaId:           PropId(Id),

		"bytes": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The number of bytes that have been transferred so far.",
		},
		"file": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
			Description: "The mount path the received snapshots are written to, or the file the stream is read " +
				"from or written to.",
		},
		"fs": {
			Type:        schema.TypeString,
			Optional:    true,
			ForceNew:    true,
			Description: "The label of the Btrfs file system taking part in the transfer.",
		},
		"send_parent": {
			Type:        schema.TypeString,
			Optional:    true,
			ForceNew:    true,
			Description: "The parent snapshot used to send an incremental stream instead of a full one.",
		},
		"send_subvolumes": {
			Type:        schema.TypeString,
			Optional:    true,
			ForceNew:    true,
			Description: "The subvolumes that are sent to the remote device.",
		},
		"ssh_address": {
			Type:        schema.TypeString,
			Optional:    true,
			ForceNew:    true,
			Description: "The address of the remote device the snapshots are transferred to or from.",
		},
		"ssh_port": {
			Type:         schema.TypeInt,
			Optional:     true,
			ForceNew:     true,
			Description:  "The SSH port of the remote device.",
			ValidateFunc: validation.IntBetween(1, 65535),
		},
		"ssh_receive_mount": {
			Type:        schema.TypeString,
			Optional:    true,
			ForceNew:    true,
			Description: "The destination path on the remote device the received snapshots are stored in.",
		},
		"ssh_user": {
			Type:        schema.TypeString,
			Optional:    true,
			ForceNew:    true,
			Description: "The user on the remote device that is authorized for the transfer operation.",
		},
		"status": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The current state of the transfer.",
		},
		"type": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
			Description: "The direction of the transfer." +
				"\n  * receive - accept a stream from a remote device;" +
				"\n  * send - send subvolumes to a remote device.",
			ValidateFunc: validation.StringInSlice([]string{"receive", "send"}, false),
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
