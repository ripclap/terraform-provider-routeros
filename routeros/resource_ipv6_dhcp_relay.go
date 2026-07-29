package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
	Menu is empty on the reference device (RouterOS 7.23.2), so no GET sample is available.
*/

// ResourceIPv6DhcpRelay https://help.mikrotik.com/docs/spaces/ROS/pages/24805500/DHCP
func ResourceIPv6DhcpRelay() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ipv6/dhcp-relay"),
		MetaId:           PropId(Id),

		KeyComment: PropCommentRw,
		"delay_threshold": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "If secs field in DHCP packet is smaller than delay-threshold, then this packet is ignored.",
			DiffSuppressFunc: TimeEqual,
		},
		"dhcp_options": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Description: "List of DHCPv6 options (names of the entries in `/ipv6/dhcp-relay/option`) that the relay " +
				"inserts into the forwarded DHCP packets. *Default: `client_mac`*",
		},
		"dhcp_server": {
			Type:     schema.TypeSet,
			Required: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Description: "A list of DHCPv6 server addresses the DHCP requests should be forwarded to. A link-local " +
				"server address has to be written together with the outgoing interface, `[IPv6]%interface`.",
		},
		KeyDisabled: PropDisabledRw,
		KeyInterface: {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Interface name the DHCPv6 relay will be working on.",
		},
		KeyInvalid: PropInvalidRo,
		"link_address": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "An IPv6 address that may be used by the server to identify the link on which the client is " +
				"located. *Default: `::`*",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyName: PropName("Descriptive name for the relay."),
		"store_relayed_bindings": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Inspects the relayed DHCP advertisements and stores the assigned prefixes.",
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
