package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
  ".id": "*2",
  "acquired": "false",
  "block-device": "false",
  "compress": "false",
  "disabled": "false",
  "empty": "false",
  "encrypted": "false",
  "file-offset": "0",
  "file-path": "/disk1/swap",
  "file-size": "17179856896",
  "formatting": "false",
  "free": "17179852800",
  "fs": "-",
  "guid-partition-table": "false",
  "iscsi-export": "false",
  "media-interface": "none",
  "media-sharing": "false",
  "model": "/disk1/swap",
  "mount-filesystem": "true",
  "mount-read-only": "false",
  "mounted": "false",
  "nfs-sharing": "false",
  "nvme-tcp-export": "false",
  "parent": "",
  "partition": "false",
  "raid-master": "none",
  "raid-member": "false",
  "raid-member-failed": "false",
  "self-encrypted-and-locked": "false",
  "self-encryption-enabled": "false",
  "self-encryption-supported": "false",
  "size": "17179852800",
  "slot": "file-disk1-swap",
  "slot-default": "",
  "smb-sharing": "false",
  "swap": "true",
  "swap-enabled": "true",
  "type": "file"
}
*/

// ResourceDisk A device in the `/disk` menu: block, file, RAM, RAID or remote share. Physical drives are
// imported; virtual ones are created here. Formatting is a separate action (`/disk/format-drive`), out of scope.
// https://help.mikrotik.com/docs/spaces/ROS/pages/91193346/Disks
// https://help.mikrotik.com/docs/spaces/ROS/pages/259031065/ROSE-storage
func ResourceDisk() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/disk"),
		MetaId:           PropId(Id),
		// Pure telemetry (I/O accounting, SMART counters); no configuration meaning.
		MetaSkipFields: PropSkipFields("active_time", "available_spare", "available_spare_threshold",
			"controller_busy_time", "critical_temperature", "critical_temperature_time", "critical_warning",
			"discard_bytes", "discard_merges", "discard_ops", "discard_time", "flush_ops", "flush_time",
			"free_inodes", "fw_version", "host_read_bytes", "host_read_cmds", "host_write_bytes",
			"host_write_cmds", "in_flight_ops", "interface_speed", "io_errors", "io_ops", "percentage_used",
			"power_cycles", "power_on_time", "read_bytes", "read_merges", "read_ops", "read_ops_per_second",
			"read_rate", "read_time", "sector_size", "temperature", "temperatures", "total_inodes",
			"unrecovered_integrity_errors", "unsafe_shutdowns", "use", "wait_time", "warning_temperature",
			"warning_temperature_time", "write_bytes", "write_merges", "write_ops", "write_ops_per_second",
			"write_rate", "write_time"),

		"acquired": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether the device is claimed by another subsystem, such as a RAID array.",
		},
		"block_device": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether the entry represents a block device.",
		},
		KeyComment: PropCommentRw,
		"compress": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Enables transparent compression on the file system of this device.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"crypted_backend": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The slot of the drive or partition that holds the encrypted data when `type` is " +
				"`crypted`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyDisabled: PropDisabledRw,
		"empty": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether the device has no file system on it.",
		},
		"encrypted": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether the device holds encrypted content.",
		},
		"encryption_key": {
			Type:             schema.TypeString,
			Optional:         true,
			Sensitive:        true,
			Description:      "The key used to unlock an encrypted device (`type` is `crypted`).",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"file_offset": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The offset, in bytes, inside the backing file at which the device starts.",
			DiffSuppressFunc: BytesEqual,
		},
		"file_path": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The path of the backing file when `type` is `file`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"file_size": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The size of the backing file when `type` is `file`. Accepts a byte count, a value with " +
				"a `K`, `M`, `G` or `T` suffix, or `auto`.",
			// Not BytesEqual: the property also accepts the literal `auto`, which the byte parser
			// used by BytesEqual cannot handle.
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"formatting": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether a format operation is currently running on this device.",
		},
		"free": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The free space of the file system, in bytes.",
		},
		"fs": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The file system detected on the device.",
		},
		"fs_label": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The label of the file system on the device.",
		},
		"fs_uuid": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The UUID of the file system on the device.",
		},
		"guid_partition_table": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether the device carries a GPT partition table.",
		},
		"interface": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The hardware interface the drive is attached through, for example `PCIe 4x8 GT/s`.",
		},
		"iscsi_address": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The address of the iSCSI target this device connects to.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"iscsi_export": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Exports this device as an iSCSI target.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"iscsi_iqn": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The iSCSI qualified name of the remote target.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"iscsi_port": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "The TCP port of the remote iSCSI target.",
			ValidateFunc:     validation.IntBetween(1, 65535),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"iscsi_server_iqn": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The iSCSI qualified name announced when this device is exported as a target.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"iscsi_server_port": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "The TCP port this device listens on when it is exported as an iSCSI target.",
			ValidateFunc:     validation.IntBetween(1, 65535),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"media_interface": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The interface used by the dynamic `/ip/media` instance created for this device.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"media_sharing": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Enables media (DLNA) sharing of this device.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"model": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The model reported by the drive.",
		},
		"mount_filesystem": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Whether the file system found on this device is mounted.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"mount_point": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The directory the file system is mounted at.",
		},
		"mount_point_template": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The template used to build the mount point. Supports the `[slot]`, `[model]`, " +
				"`[serial]`, `[fs-label]`, `[fs-uuid]` and `[fs]` variables.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"mount_read_only": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Mounts the file system read-only.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"mounted": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether the file system of this device is currently mounted.",
		},
		"nfs_address": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The address of the NFS server this device mounts a share from.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"nfs_share": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The remote folder mounted from the NFS server.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"nfs_sharing": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Exports this device over NFS.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"nvme_tcp_address": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The address of the NVMe over TCP target this device connects to.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"nvme_tcp_export": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Exports this device as an NVMe over TCP target.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"nvme_tcp_host_name": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The initiator host name used to identify this device to the NVMe over TCP target.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"nvme_tcp_nqn": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The NVMe qualified name of the remote target.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"nvme_tcp_password": {
			Type:             schema.TypeString,
			Optional:         true,
			Sensitive:        true,
			Description:      "The password used to authenticate against the remote NVMe over TCP target.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"nvme_tcp_port": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "The TCP port of the remote NVMe over TCP target.",
			ValidateFunc:     validation.IntBetween(1, 65535),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"nvme_tcp_secret": {
			Type:        schema.TypeString,
			Computed:    true,
			Sensitive:   true,
			Description: "The generated secret used by the NVMe over TCP initiator.",
		},
		"nvme_tcp_server_allow_host_name": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Restricts the exported NVMe over TCP target to the initiators announcing the configured " +
				"host name.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"nvme_tcp_server_nqn": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The NVMe qualified name announced when this device is exported as a target.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"nvme_tcp_server_password": {
			Type:             schema.TypeString,
			Optional:         true,
			Sensitive:        true,
			Description:      "The password the initiators have to supply to reach the exported target.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"nvme_tcp_server_port": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "The TCP port the exported NVMe over TCP target listens on.",
			ValidateFunc:     validation.IntBetween(1, 65535),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"nvme_tcp_server_secret": {
			Type:        schema.TypeString,
			Computed:    true,
			Sensitive:   true,
			Description: "The generated secret used by the exported NVMe over TCP target.",
		},
		"parent": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The slot of the parent device. Used, for example, when a partition is created on an " +
				"existing drive.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"partition": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether the entry is a partition of another device.",
		},
		"partition_number": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "The number of the partition inside the parent device.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"partition_offset": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The offset, in bytes, at which the partition starts inside the parent device, or " +
				"`auto` to let the system place it.",
			// Not BytesEqual: the property also accepts the literal `auto`, which the byte parser
			// used by BytesEqual cannot handle.
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"partition_size": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The size of the partition, in bytes or with a `K`, `M`, `G` or `T` suffix.",
			DiffSuppressFunc: BytesEqual,
		},
		"raid_chunk_size": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The size of the chunks or stripes used by the RAID array.",
			DiffSuppressFunc: BytesEqual,
		},
		"raid_device_count": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "The number of devices that make up the RAID array.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"raid_master": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The slot of the RAID array this device is a member of. `none` if it is not a member.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"raid_max_component_size": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The maximum amount of space used on each member of the RAID array, or `none` for " +
				"no limit.",
			// Not BytesEqual: the property also accepts the literal `none`, which the byte parser
			// used by BytesEqual cannot handle.
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"raid_member": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether this device is a member of a RAID array.",
		},
		"raid_member_failed": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Marks this member of a RAID array as failed.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"raid_member_state": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The state of this device inside the RAID array.",
		},
		"raid_role": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The position of this device inside the RAID array: the index of the member, or `spare` " +
				"for a hot spare.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"raid_type": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The RAID level of the array." +
				"\n  * 0;" +
				"\n  * 1;" +
				"\n  * 4;" +
				"\n  * 5;" +
				"\n  * 6;" +
				"\n  * linear.",
			ValidateFunc:     validation.StringInSlice([]string{"0", "1", "4", "5", "6", "linear"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"raid_uuid": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The UUID of the RAID array.",
		},
		"ramdisk_size": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The size of the RAM block device when `type` is `ramdisk`.",
			DiffSuppressFunc: BytesEqual,
		},
		"self_encrypted_and_locked": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether a self-encrypting drive is currently locked.",
		},
		"self_encryption_enabled": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether the self-encryption of the drive is enabled.",
		},
		"self_encryption_password": {
			Type:             schema.TypeString,
			Optional:         true,
			Sensitive:        true,
			Description:      "The password of a TCG OPAL self-encrypting drive.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"self_encryption_supported": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether the drive supports TCG OPAL self-encryption.",
		},
		"serial": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The serial number reported by the drive.",
		},
		"size": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The size of the device, in bytes.",
		},
		"slot": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			Description: "The name of the slot. For physical drives it is derived from the connection location, " +
				"for the devices created here it is the name of the new device.",
		},
		"slot_default": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The default slot name of a physical drive.",
		},
		"smb_address": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The address of the SMB server this device mounts a share from.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"smb_encryption": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Requires encryption on the SMB connection to the remote share.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"smb_password": {
			Type:             schema.TypeString,
			Optional:         true,
			Sensitive:        true,
			Description:      "The password used to mount the remote SMB share.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"smb_server_encryption": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Requires encryption from the clients of the SMB share exported by this device.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"smb_server_password": {
			Type:             schema.TypeString,
			Optional:         true,
			Sensitive:        true,
			Description:      "The password the clients have to supply to reach the exported SMB share.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"smb_server_user": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The user the clients have to authenticate as to reach the exported SMB share.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"smb_share": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The name of the remote SMB share that is mounted.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"smb_sharing": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Exports this device over SMB.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"smb_user": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The user name used to mount the remote SMB share.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"sshfs_address": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The address of the SSH server this device mounts a directory from.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"sshfs_local_user": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The local user the SSHFS mount runs as.",
		},
		"sshfs_password": {
			Type:             schema.TypeString,
			Optional:         true,
			Sensitive:        true,
			Description:      "The password used to authenticate against the SSH server.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"sshfs_path": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The remote directory mounted over SSHFS.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"sshfs_port": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "The TCP port of the remote SSH server.",
			ValidateFunc:     validation.IntBetween(1, 65535),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"sshfs_user": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The user name used to authenticate against the SSH server.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"state": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The current state of the device.",
		},
		"swap": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Uses this device, partition or file as swap space.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"swap_enabled": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether the device is currently active as swap space.",
		},
		"tmpfs_max_size": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The maximum amount of RAM used by a `tmpfs` device. Defaults to half of the available " +
				"RAM when it is not set.",
			DiffSuppressFunc: BytesEqual,
		},
		"type": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			Description: "The kind of device this entry represents." +
				"\n  * crypted - an encrypted device on top of `crypted_backend`;" +
				"\n  * file - a device backed by a file;" +
				"\n  * hardware - a physical drive detected by the system;" +
				"\n  * iscsi - a remote iSCSI target;" +
				"\n  * nfs - a remote NFS share;" +
				"\n  * nvme-tcp - a remote NVMe over TCP target;" +
				"\n  * partition - a partition of another device;" +
				"\n  * raid - a software RAID array;" +
				"\n  * ramdisk - a RAM backed block device;" +
				"\n  * smb - a remote SMB share;" +
				"\n  * sshfs - a remote directory mounted over SSH;" +
				"\n  * tmpfs - a RAM backed file system.",
			ValidateFunc: validation.StringInSlice([]string{"crypted", "file", "hardware", "iscsi", "nfs",
				"nvme-tcp", "partition", "raid", "ramdisk", "smb", "sshfs", "tmpfs"}, false),
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
