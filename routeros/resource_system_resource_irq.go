package routeros

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
{
  ".id": "*15",
  "active-cpu": "5",
  "count": "17359",
  "cpu": "auto",
  "irq": "21",
  "per-cpu-count": "3,0,0,0,0,17356,0,0,0,0,0,0,0,0,0,0",
  "read-only": "false",
  "users": "ttyS0"
}
*/

// ResourceSystemResourceIrq CPU assignment of a hardware interrupt.
// https://help.mikrotik.com/docs/spaces/ROS/pages/40992875/Resource
func ResourceSystemResourceIrq() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/system/resource/irq"),
		MetaId:           PropId(Id),
		MetaSkipFields:   PropSkipFields("irq"),
		// `count` is a name reserved by Terraform, expose the RouterOS field as `irq_count`.
		MetaTransformSet: PropTransformSet("irq_count: count"),

		"active_cpu": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The CPU that currently serves the interrupt on a multicore system.",
		},
		"irq_count": {
			Type:     schema.TypeString,
			Computed: true,
			Description: "The number of interrupts that have been served (the RouterOS `count` field). " +
				"On Ethernet interfaces one interrupt equals one packet.",
		},
		"cpu": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The CPU that has to serve this interrupt: `auto` lets the system distribute the " +
				"interrupt, an integer pins it to that CPU core.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"irq": {
			Type:     schema.TypeInt,
			Required: true,
			Description: "The interrupt number. The interrupt has to exist already; the entry is looked up by " +
				"this number.",
		},
		"per_cpu_count": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The number of interrupts served by each of the CPU cores, in core order.",
		},
		"read_only": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether the CPU assignment of this interrupt is fixed and cannot be changed.",
		},
		"users": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The device or process the interrupt is assigned to.",
		},
	}

	resCreateUpdate := func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		item, metadata := TerraformResourceDataToMikrotik(resSchema, d)

		irq := strconv.Itoa(d.Get("irq").(int))
		res, err := ReadItemsFiltered([]string{"irq=" + irq}, metadata.Path, m.(Client))
		if err != nil {
			return diag.FromErr(err)
		}

		if len(*res) == 0 {
			d.SetId("")
			return diag.Errorf("interrupt irq=%v not found", irq)
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
