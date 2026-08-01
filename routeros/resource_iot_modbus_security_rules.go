package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
  `GET /rest/iot/modbus/security-rules` returns `[]` on RouterOS 7.23
  (RouterOS 7.23) - no rules are configured, so no live record could be
  captured. The shape below uses the field names reported by the device and the values from
  the MikroTik example `/iot modbus security-rules add ip-range=0.0.0.0/0 allowed-function-codes=3,6`:

  {
    ".id": "*1",
    "allowed-function-codes": "3,6",
    "comment": "",
    "disabled": "false",
    "ip-range": "0.0.0.0/0"
  }
*/

// ResourceIotModbusSecurityRules Modbus security rules of the `iot` package. Each rule permits a set of
// Modbus function codes to an address range; rules are only evaluated while
// `/iot/modbus disable-security-rules=no`.
// https://help.mikrotik.com/docs/spaces/UM/pages/61046813/RB924i-2nD-BT5+BG77+Modbus+configuration
func ResourceIotModbusSecurityRules() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/iot/modbus/security-rules"),
		MetaId:           PropId(Id),

		"allowed_function_codes": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeInt},
			Description: "Modbus function codes that will be accessible from the configured address range, for " +
				"example `[3, 6]`. A rule with no function codes permits nothing.",
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"ip_range": {
			Type:     schema.TypeString,
			Required: true,
			Description: "IP address range or network, in address/netmask notation, that the rule applies to, for " +
				"example `0.0.0.0/0` or `192.168.88.0/24`.",
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
