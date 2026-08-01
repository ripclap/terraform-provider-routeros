package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
  Live `GET /rest/iot/modbus` from a RouterOS 7.23 device running RouterOS 7.23:

  {
    "disable-security-rules": "true",
    "disabled": "true",
    "hardware-port": "serial0",
    "interframe-gap": "0",
    "rx-switch-offset": "0",
    "tcp-port": "502",
    "timeout": "1000"
  }
*/

// ResourceIotModbus Modbus server (slave) settings of the `iot` package.
// https://help.mikrotik.com/docs/spaces/UM/pages/61046813/RB924i-2nD-BT5+BG77+Modbus+configuration
func ResourceIotModbus() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/iot/modbus"),
		MetaId:           PropId(Id),

		"disable_security_rules": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Enables or disables the security-rule feature. While the rules are disabled every host is " +
				"allowed to use every Modbus function code.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyDisabled: PropDisabledRw,
		"hardware_port": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Name of the port from the `/port` menu that is assigned to the Modbus RTU service. On boards " +
				"with a dedicated Modbus port the port is named `modbus`; on such boards the only available port " +
				"is the serial console port `serial0`. Modbus TCP does not depend on this setting.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"interframe_gap": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Additional gap inserted between Modbus RTU frames. The value is expressed in milliseconds " +
				"and `0` represents the default gap of roughly 4 ms.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"rx_switch_offset": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Controls how fast the Modbus window is switched from transmit to receive. The value is " +
				"expressed in microseconds.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"tcp_port": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "TCP port that the device will use for Modbus TCP. Default: `502`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"timeout": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Request timeout value in milliseconds. MikroTik documents the range as 0..1000, default " +
				"`1000`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
	}

	return &schema.Resource{
		CreateContext: DefaultSystemCreate(resSchema),
		ReadContext:   DefaultSystemRead(resSchema),
		UpdateContext: DefaultSystemUpdate(resSchema),
		DeleteContext: DefaultSystemDelete(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resSchema,
	}
}
