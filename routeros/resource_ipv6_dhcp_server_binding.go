package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
	Menu is empty on the reference device (RouterOS 7.23.2), so no GET sample is available.
*/

// ResourceIPv6DhcpServerBinding https://help.mikrotik.com/docs/spaces/ROS/pages/24805500/DHCP
func ResourceIPv6DhcpServerBinding() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ipv6/dhcp-server/binding"),
		MetaId:           PropId(Id),

		"address": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The IPv6 address or prefix that will be assigned to the client. May be left unset when the " +
				"assignment is taken from `prefix_pool`.",
		},
		"address_lists": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Description: "Firewall address lists to which the allocated address or prefix will be added while the " +
				"binding is bound.",
		},
		"allow_dual_stack_queue": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Creates a single simple queue entry for both the IPv4 and the IPv6 address, using the MAC " +
				"address and the DUID for identification.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyComment: PropCommentRw,
		"dhcp_option": {
			Type:        schema.TypeSet,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "Add additional DHCP options from the `/ipv6/dhcp-server/option` list.",
		},
		KeyDisabled: PropDisabledRw,
		"duid": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "DUID value of the client the binding belongs to (hexadecimal format only).",
		},
		KeyDynamic: PropDynamicRo,
		"expires_after": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Time left until the binding expires.",
		},
		"ia_type": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Identity Association type the binding is created for:" +
				"\n  * na - non-temporary address (IA_NA), a single address handed out to the client;" +
				"\n  * pd - prefix delegation (IA_PD), a prefix delegated to a downstream router.",
			ValidateFunc:     validation.StringInSlice([]string{"na", "pd"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"iaid": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Identity Association Identifier, part of the client ID. Integer `0..4294967295`. " +
				"The upper bound is not enforced here because the constant overflows `int` on 32-bit builds.",
			ValidateFunc:     validation.IntAtLeast(0),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"insert_queue_before": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Specify where to place the dynamic simple queue entry created for a static binding with the " +
				"`rate_limit` parameter set. The console offers `bottom` and `first`; a queue name is accepted too, " +
				"so the value is not constrained here.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyInvalid: PropInvalidRo,
		"last_seen": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Time elapsed since the client was last heard from.",
		},
		"life_time": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Time period after which the binding expires. *Default: `3d`*",
			DiffSuppressFunc: TimeEqual,
		},
		"parent_queue": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The dynamically created queue for this binding will be configured as a child queue of the " +
				"specified parent queue.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"prefix_pool": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Prefix pool the prefix advertised to the DHCPv6 client is taken from.",
		},
		"queue_type": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Queue type used for the dynamically created simple queue of this binding.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"radius": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Shows whether this dynamic binding was authenticated by RADIUS.",
		},
		"rate_limit": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Adds a dynamic simple queue that limits the bandwidth of the client to the specified rate. " +
				"Requires the binding to be static.",
		},
		"reconfigure_key": {
			Type:     schema.TypeString,
			Computed: true,
			Description: "Key used to authenticate the DHCPv6 Reconfigure messages sent to this client. " +
				"VERIFY: field presence confirmed on ROS 7.23.2, the exact value format was not observable " +
				"(no bindings on the reference device).",
		},
		"reconfigure_last_sent": {
			Type:     schema.TypeString,
			Computed: true,
			Description: "When the last DHCPv6 Reconfigure message was sent to this client. " +
				"VERIFY: field presence confirmed on ROS 7.23.2, the exact value format was not observable " +
				"(no bindings on the reference device).",
		},
		"reconfigure_status": {
			Type:     schema.TypeString,
			Computed: true,
			Description: "Result of the last DHCPv6 Reconfigure exchange with this client. " +
				"VERIFY: field presence confirmed on ROS 7.23.2, the set of possible values was not observable " +
				"(no bindings on the reference device).",
		},
		"server": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Name of the DHCPv6 server this binding belongs to, or `all`. *Default: `all`*",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"status": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Binding status (`waiting`, `offered`, `bound`).",
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
