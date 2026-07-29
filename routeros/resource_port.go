package routeros

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
	{
	  ".id": "*1",
	  "baud-rate": "115200",
	  "channels": "1",
	  "data-bits": "8",
	  "device": "",
	  "flow-control": "none",
	  "inactive": "false",
	  "line-state": "dtr,rts,cts,dcd,dsr",
	  "name": "serial0",
	  "parity": "none",
	  "stop-bits": "1",
	  "used-by": "Serial Console(#0)"
	}

	Serial ports are created by the hardware: this resource adopts the existing port by its name, and
	destroying it only drops it from the Terraform state.
*/

// ResourcePort Serial port (RS-232 / USB serial) settings.
// https://help.mikrotik.com/docs/spaces/ROS/pages/8978525/Ports
func ResourcePort() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/port"),
		MetaId:           PropId(Id),

		"baud_rate": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Speed of the port in bits per second, or `auto` to detect it automatically, for example " +
				"`115200`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"channels": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "Number of channels the port provides.",
		},
		"data_bits": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Number of data bits per character. Use `7` for ASCII data and `8` for binary data.",
			ValidateFunc:     validation.IntInSlice([]int{7, 8}),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"device": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Name of the device that provides this port, for USB serial adapters.",
		},
		"dtr": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "State of the RS-232 DTR signal circuit used by hardware flow control." +
				"\n> RouterOS does not report this property back in `/port/print`, so its value is kept from the " +
				"configuration and cannot be refreshed from the device.",
			ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
		},
		"flow_control": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Method used to pause and resume the transmission.",
			ValidateFunc:     validation.StringInSlice([]string{"hardware", "none", "xon-xoff"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyInactive: PropInactiveRo,
		"line_state": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Comma separated list of the RS-232 signal lines that are currently asserted.",
		},
		KeyName: PropName("Name of the serial port, for example `serial0`. On the first apply the port with this " +
			"name must already exist on the router; afterwards the port is tracked by its internal identifier and " +
			"changing the name renames the port."),
		"parity": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Parity bit used for error detection.",
			ValidateFunc:     validation.StringInSlice([]string{"even", "none", "odd"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"rts": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "State of the RS-232 RTS signal circuit used by hardware flow control." +
				"\n> RouterOS does not report this property back in `/port/print`, so its value is kept from the " +
				"configuration and cannot be refreshed from the device.",
			ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
		},
		"stop_bits": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Number of stop bits that follow every character.",
			ValidateFunc:     validation.IntInSlice([]int{1, 2}),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"used_by": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Service that currently holds the port, for example `Serial Console(#0)`.",
		},
	}

	return &schema.Resource{
		CreateContext: resourcePortCreate(resSchema),
		ReadContext:   DefaultRead(resSchema),
		UpdateContext: DefaultUpdate(resSchema),
		DeleteContext: DefaultSystemDelete(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: ImportStateCustomContext(resSchema),
		},

		Schema: resSchema,
	}
}

// resourcePortCreate adopts the serial port that already exists on the router under the
// configured name and then applies the configuration to it.
func resourcePortCreate(s map[string]*schema.Schema) schema.CreateContextFunc {
	return func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		metadata := GetMetadata(s)

		name := d.Get(KeyName).(string)
		items, err := ReadItemsFiltered(buildReadFilter(map[string]any{KeyName: name}), metadata.Path, m.(Client))
		if err != nil {
			return diag.FromErr(err)
		}

		if items == nil || len(*items) != 1 {
			return diag.Errorf("expected exactly one '%v' entry named '%v'", metadata.Path, name)
		}

		d.SetId((*items)[0].GetID(Id))

		return ResourceUpdate(ctx, s, d, m)
	}
}
