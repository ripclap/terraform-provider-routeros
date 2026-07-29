package routeros

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  {
    ".id": "*8",
    "address": "192.168.100.2/24",
    "comment": "comment",
    "disabled": "false",
    "gateway": "192.168.100.1",
    "name": "veth1",
    "running": "true"
  }
*/

// https://help.mikrotik.com/docs/display/ROS/Container
func ResourceInterfaceVeth() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/veth"),
		MetaId:           PropId(Id),

		"address": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "Ip address.",
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.IsCIDR,
			},
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"container_mac_address": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "MAC address presented to the container side of the veth pair.",
		},
		"dhcp": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"dhcp_address": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Address obtained by the veth DHCP client, when dhcp is enabled.",
		},
		"gateway": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Gateway IP address.",
			// The device reports an empty string when no gateway is set.
			ValidateFunc: validation.Any(validation.StringIsEmpty, validation.IsIPv4Address),
		},
		"gateway6": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Gateway IPv6 address.",
			// The device reports an empty string when no gateway is set.
			ValidateFunc: validation.Any(validation.StringIsEmpty, validation.IsIPv6Address),
		},
		KeyMacAddress: {
			Type:         schema.TypeString,
			Description:  "MAC address.",
			Optional:     true,
			Computed:     true,
			ValidateFunc: validation.IsMACAddress,
			DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
				if old != "" && d.GetRawConfig().GetAttr(k).IsNull() {
					return true
				}
				return strings.EqualFold(old, new)
			},
		},
		KeyName:    PropName("Interface name."),
		KeyRunning: PropRunningRo,
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
