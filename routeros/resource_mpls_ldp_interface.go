package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  The menu is empty on the reference device (ROS 7.23.2, MPLS not configured),
  the field set below is taken from `/console/inspect request=syntax path="mpls,ldp,interface,add"`
  and from `print proplist=` completion on that device.

  {
    ".id": "*1",
    "accept-dynamic-neighbors": "true",
    "afi": "ip",
    "comment": "",
    "disabled": "false",
    "hello-interval": "5s",
    "hold-time": "15s",
    "interface": "ether1",
    "transport-addresses": "10.0.0.1"
  }
*/

// ResourceMplsLdpInterface https://help.mikrotik.com/docs/spaces/ROS/pages/121995275/LDP
func ResourceMplsLdpInterface() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/mpls/ldp/interface"),
		MetaId:           PropId(Id),

		"accept_dynamic_neighbors": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Whether to accept LDP neighbors discovered by the basic (link) hello messages received " +
				"on this interface.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"afi": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "Address families for which LDP is enabled on this interface.",
			Elem: &schema.Schema{
				Type:             schema.TypeString,
				ValidateDiagFunc: ValidationValInSlice([]string{"ip", "ipv6"}, false, false),
			},
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"hello_interval": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Interval between the LDP hello messages sent on this interface.",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEqual,
		},
		"hold_time": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Time the neighbor discovered on this interface is kept without receiving a hello " +
				"message from it.",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEqual,
		},
		"interface": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of the interface or of the interface list LDP is run on.",
		},
		"transport_addresses": {
			Type:     schema.TypeList,
			Optional: true,
			Description: "List of the IPv4/IPv6 addresses advertised as the LDP transport addresses in the hello " +
				"messages sent on this interface. Overrides the instance-wide setting. The order is significant.",
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPAddress,
			},
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
