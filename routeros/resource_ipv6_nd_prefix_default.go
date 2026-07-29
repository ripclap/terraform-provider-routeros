package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
{
  "autonomous": "true",
  "dhcp6-pd-preferred": "false",
  "preferred-lifetime": "1w",
  "valid-lifetime": "4w2d"
}
*/

// ResourceIPv6NdPrefixDefault Default template for RouterOS auto-generated advertised prefixes.
// https://help.mikrotik.com/docs/spaces/ROS/pages/40992815/IPv6+Neighbor+Discovery
func ResourceIPv6NdPrefixDefault() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ipv6/nd/prefix/default"),
		MetaId:           PropId(Id),

		"autonomous": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "When set, indicates that this prefix can be used for autonomous address configuration. " +
				"Otherwise, the prefix information is silently ignored. *Default: `yes`*",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"dhcp6_pd_preferred": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Advertise the prefixes obtained by the DHCPv6-PD client with this default template. " +
				"*Default: `no`*",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"preferred_lifetime": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Timeframe (relative to the time the packet is sent) after which the generated address " +
				"becomes `deprecated`. A deprecated address is used only for already existing connections and is " +
				"usable until the valid lifetime expires. *Default: `1w`*",
			DiffSuppressFunc: TimeEqual,
		},
		"valid_lifetime": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The length of time (relative to the time the packet is sent) an address remains in the " +
				"valid state. The valid lifetime must be greater than or equal to the preferred lifetime. " +
				"*Default: `4w2d`*",
			DiffSuppressFunc: TimeEqual,
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
