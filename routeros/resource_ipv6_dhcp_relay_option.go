package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  {
    ".id": "*FFFFFFF3",
    "code": "79",
    "default": "true",
    "name": "client_mac",
    "only-if-mac-available": "true",
    "raw-value": "0001",
    "value": "0x0001$(CLIENT_MAC)"
  }
*/

// ResourceIPv6DhcpRelayOption https://help.mikrotik.com/docs/spaces/ROS/pages/24805500/DHCP
func ResourceIPv6DhcpRelayOption() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ipv6/dhcp-relay/option"),
		MetaId:           PropId(Id),

		"code": {
			Type:     schema.TypeInt,
			Required: true,
			Description: "DHCPv6 option code. DHCPv6 option codes are 16 bit wide " +
				"([RFC 8415](https://www.rfc-editor.org/rfc/rfc8415)); the built-in `client_mac` entry uses code 79 " +
				"(OPTION_CLIENT_LINKLAYER_ADDR).",
			ValidateFunc: validation.IntBetween(1, 65535),
		},
		KeyComment: PropCommentRw,
		KeyDefault: PropDefaultRo,
		KeyName:    PropName("Descriptive name of the option."),
		"only_if_mac_available": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Insert the option only when the MAC address of the client is known to the relay, that is, " +
				"when the client is directly attached to the relay.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"raw_value": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Read-only field which shows the raw DHCP option value (the format actually sent out).",
		},
		"value": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Parameter's value. Available data types for options are:\n" +
				"    - `'test'` -> ASCII to Hex 0x74657374\n" +
				"    - `s'160'` -> ASCII to Hex 0x313630\n" +
				"    - `'10'` -> Decimal to Hex 0x0a\n" +
				"    - `0x0a0a` -> No conversion\n" +
				"    - `$(VARIABLE)` -> hardcoded values\n\n" +
				"Data types can be combined, the built-in `client_mac` option is defined as " +
				"`0x0001$(CLIENT_MAC)`.",
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
