package routeros

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
{
  ".id": "*1",
  "balancing": "false",
  "default-subvolume": "data",
  "label": "BtrfsDisk",
  "replacing": "false",
  "scrubbing": "false",
  "total-devs": "1",
  "uuid": "00000000-0000-0000-0000-000000000000"
}
*/

// ResourceDiskBtrfsFilesystem A Btrfs file system. It cannot be added or removed from configuration, so
// this resource looks it up by `uuid` and only manages its writable attributes.
// https://help.mikrotik.com/docs/spaces/ROS/pages/295239711/Btrfs
func ResourceDiskBtrfsFilesystem() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/disk/btrfs/filesystem"),
		MetaId:           PropId(Id),
		MetaSkipFields:   PropSkipFields("uuid"),

		"balance_status": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The progress of the running or last balance operation.",
		},
		"balancing": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether a balance operation is currently running.",
		},
		"corruption_errors": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The number of corruption errors recorded by the file system.",
		},
		"default_subvolume": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The subvolume that is mounted as the root of the file system when no subvolume is " +
				"given explicitly.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"dev_ids": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The internal Btrfs device ids of the devices that make up the file system.",
		},
		"devs": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The devices that make up the file system.",
		},
		"flush_errors": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The number of flush errors recorded by the file system.",
		},
		"generation_errors": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The number of generation errors recorded by the file system.",
		},
		"label": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The label of the Btrfs file system.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"missing_devs": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The devices of the file system that are currently missing.",
		},
		"read_errors": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The number of read errors recorded by the file system.",
		},
		"replace_status": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The progress of the running or last device replace operation.",
		},
		"replacing": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether a device replace operation is currently running.",
		},
		"scrub_status": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The progress of the running or last scrub operation.",
		},
		"scrubbing": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether a scrub operation is currently running.",
		},
		"spaces": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The allocation of the data, metadata and system block groups.",
		},
		"total_devs": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The number of devices that make up the file system.",
		},
		"uuid": {
			Type:     schema.TypeString,
			Required: true,
			Description: "The UUID of the Btrfs file system. The file system has to exist already; it is looked " +
				"up by this UUID.",
		},
		"write_errors": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The number of write errors recorded by the file system.",
		},
	}

	resCreateUpdate := func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		item, metadata := TerraformResourceDataToMikrotik(resSchema, d)

		uuid := d.Get("uuid").(string)
		res, err := ReadItemsFiltered([]string{"uuid=" + uuid}, metadata.Path, m.(Client))
		if err != nil {
			return diag.FromErr(err)
		}

		if len(*res) == 0 {
			d.SetId("")
			return diag.Errorf("Btrfs filesystem uuid=%v not found", uuid)
		}

		d.SetId((*res)[0].GetID(Id))
		item[".id"] = d.Id()

		var resUrl string
		if m.(Client).GetTransport() == TransportREST {
			resUrl = "/set"
		}

		if err := m.(Client).SendRequest(crudPost, &URL{Path: metadata.Path + resUrl}, item, nil); err != nil {
			return diag.FromErr(err)
		}

		return ResourceRead(ctx, resSchema, d, m)
	}

	return &schema.Resource{
		CreateContext: resCreateUpdate,
		ReadContext:   DefaultRead(resSchema),
		UpdateContext: resCreateUpdate,
		DeleteContext: DefaultSystemDelete(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: ImportStateCustomContext(resSchema),
		},

		Schema: resSchema,
	}
}
