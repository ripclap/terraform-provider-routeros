package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
  ".id": "*1",
  "address": "192.0.2.10",
  "comment": "",
  "complete": "true",
  "dhcp": "false",
  "disabled": "false",
  "dynamic": "true",
  "interface": "ether1",
  "invalid": "false",
  "mac-address": "00:00:5E:00:53:00",
  "published": "false"
}
*/

// ResourceIpArp https://help.mikrotik.com/docs/display/ROS/ARP
func ResourceIpArp() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/arp"),
		MetaId:           PropId(Id),

		"address": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "IP address to be mapped.",
			ValidateFunc: validation.IsIPAddress,
		},
		KeyComment: PropCommentRw,
		"complete": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether the entry has both an address and a MAC address.",
		},
		"dhcp": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether the entry was created by the DHCP server.",
		},
		KeyDisabled:   PropDisabledRw,
		KeyDynamic:    PropDynamicRo,
		KeyInterface:  PropInterfaceRw,
		KeyInvalid:    PropInvalidRo,
		KeyMacAddress: PropMacAddressRw("MAC address to be mapped to.", false),
		"published": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Answer ARP requests for this address on behalf of the host it belongs to, so that " +
				"the router replies for a host that cannot answer for itself.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
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
