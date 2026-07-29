package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  {
    ".id": "*1",
    "afi": "",
    "disabled": "false",
    "in-filter-chain": "",
    "name": "",
    "originate-default": "",
    "out-filter-chain": "",
    "out-filter-select": "",
    "redistribute": "",
    "route-gc-timeout": "",
    "route-timeout": "",
    "routing-table": "",
    "update-interval": "",
    "vrf": "main"
  }
*/

// ResourceRoutingRipInstance https://help.mikrotik.com/docs/spaces/ROS/pages/328211/RIP
func ResourceRoutingRipInstance() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/routing/rip/instance"),
		MetaId:           PropId(Id),

		"afi": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "The address family this instance operates on.",
			ValidateFunc: validation.StringInSlice([]string{"ip", "ipv6"}, false),
		},
		KeyDisabled: PropDisabledRw,
		"in_filter_chain": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Name of the routing filter chain that is applied to the routes received by this " +
				"instance.",
		},
		KeyName: PropName("Name of the RIP instance, referenced by the interface templates and the static neighbors."),
		"originate_default": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "Controls the origination of the default route into RIP.",
			ValidateFunc: validation.StringInSlice([]string{"always", "if-installed", "never"}, false),
		},
		"out_filter_chain": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Name of the routing filter chain that is applied to the routes advertised by this " +
				"instance.",
		},
		"out_filter_select": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Name of the `/routing/filter/select-rule` chain that decides which of the candidate " +
				"routes are advertised by this instance.",
		},
		"redistribute": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
				ValidateDiagFunc: ValidationMultiValInSlice([]string{"bgp", "bgp-mpls-vpn", "connected", "dhcp",
					"fantasy", "isis", "modem", "ospf", "rip", "slaac", "static", "vpn"}, false, false),
			},
			Description: "Enable the redistribution of the listed route types into RIP.",
		},
		"route_gc_timeout": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The garbage collection time. A route that has expired is kept in the table and " +
				"advertised as unreachable for this long before it is removed.",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEqual,
		},
		"route_timeout": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The time after which a route that has not been refreshed by an update is declared " +
				"expired.",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEqual,
		},
		"routing_table": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Name of the routing table the routes of this instance are installed into.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"update_interval": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The interval between the periodic RIP updates.",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEqual,
		},
		KeyVrf: PropVrfRw,
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
