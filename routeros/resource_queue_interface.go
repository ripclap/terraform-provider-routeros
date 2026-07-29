package routeros

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
	{
	  ".id": "*7",
	  "active-queue": "only-hardware-queue",
	  "default-queue": "only-hardware-queue",
	  "interface": "ether1",
	  "queue": "only-hardware-queue"
	}

	RouterOS creates one entry per interface automatically: this resource adopts the existing entry, and
	destroying it only drops it from the Terraform state.
*/

// ResourceQueueInterface Assigns an interface queue type to an interface.
// https://help.mikrotik.com/docs/spaces/ROS/pages/74678348/Queues
func ResourceQueueInterface() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/queue/interface"),
		MetaId:           PropId(Id),

		"active_queue": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Queue type that is actually applied to the interface at the moment.",
		},
		"default_queue": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Queue type the interface driver reports as its default.",
		},
		KeyInterface: {
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: "Name of the interface whose automatically created queue entry is managed.",
		},
		"queue": {
			Type:     schema.TypeString,
			Required: true,
			Description: "Name of the `/queue/type` entry applied to the interface, for example " +
				"`only-hardware-queue`, `ethernet-default` or `no-queue`.",
		},
	}

	return &schema.Resource{
		CreateContext: resourceQueueInterfaceCreate(resSchema),
		ReadContext:   DefaultRead(resSchema),
		UpdateContext: DefaultUpdate(resSchema),
		DeleteContext: DefaultSystemDelete(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: ImportStateCustomContext(resSchema),
		},

		Schema: resSchema,
	}
}

// resourceQueueInterfaceCreate adopts the interface queue entry that RouterOS has already created
// for the requested interface and then applies the configuration to it.
func resourceQueueInterfaceCreate(s map[string]*schema.Schema) schema.CreateContextFunc {
	return func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		metadata := GetMetadata(s)

		iface := d.Get(KeyInterface).(string)
		items, err := ReadItemsFiltered(buildReadFilter(map[string]any{KeyInterface: iface}), metadata.Path, m.(Client))
		if err != nil {
			return diag.FromErr(err)
		}

		if items == nil || len(*items) != 1 {
			return diag.Errorf("expected exactly one '%v' entry for interface '%v'", metadata.Path, iface)
		}

		d.SetId((*items)[0].GetID(Id))

		return ResourceUpdate(ctx, s, d, m)
	}
}
