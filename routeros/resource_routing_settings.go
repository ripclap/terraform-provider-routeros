package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
{
  "check-gateway-ping-count": "2",
  "check-gateway-ping-interval": "10s",
  "check-gateway-ping-timeout": "1s",
  "policy-rules": "mangle,vrf-lookup,vrf-unreach,local,user,main",
  "single-process": "false"
}
*/

// ResourceRoutingSettings https://help.mikrotik.com/docs/spaces/ROS/pages/328084/IP+Routing
func ResourceRoutingSettings() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/routing/settings"),
		MetaId:           PropId(Id),

		"check_gateway_ping_count": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "How many consecutive gateway checks have to fail before the gateway is considered " +
				"unreachable.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"check_gateway_ping_interval": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The interval between the gateway reachability checks.",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEqual,
		},
		"check_gateway_ping_timeout": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "How long to wait for a reply before a single gateway check is considered failed.",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEqual,
		},
		"connected_in_chain": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Name of the routing filter chain that is applied to the connected routes before they are " +
				"installed in the routing table.",
		},
		"dynamic_in_chain": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Name of the routing filter chain that is applied to the dynamically added routes, for " +
				"example the ones coming from DHCP or PPP, before they are installed in the routing table.",
		},
		"policy_rules": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
				ValidateDiagFunc: ValidationMultiValInSlice([]string{"local", "main", "mangle", "user",
					"vrf-lookup", "vrf-unreach"}, false, false),
			},
			Description: "The ordered list of the built-in routing policy stages. It defines in which order the " +
				"routing decision consults the mangle marks, the VRF lookup and unreachable rules, the local and " +
				"the user defined `/routing/rule` entries and the main table. The order is significant.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"single_process": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Run all the routing protocol processes in a single process instead of spreading them " +
				"over the available cores. Useful for troubleshooting.",
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
