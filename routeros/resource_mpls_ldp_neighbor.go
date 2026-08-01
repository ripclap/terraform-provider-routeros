package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
  The menu is empty on RouterOS 7.23 (ROS 7.23, MPLS not configured),
  the field set below is taken from `/console/inspect request=syntax path="mpls,ldp,neighbor,add"`
  and from `print proplist=` completion on that device.

  Static entries are used to set up targeted LDP sessions; sessions discovered by the link hello
  messages appear in the same menu as dynamic entries and are not managed by Terraform.

  {
    ".id": "*1",
    "active-connect": "true",
    "addresses": "10.0.0.2",
    "comment": "",
    "disabled": "false",
    "dynamic": "false",
    "inactive": "false",
    "local-transport": "10.0.0.1",
    "on-demand": "false",
    "operational": "true",
    "passive": "false",
    "passive-wait": "false",
    "path-vector-limit": "",
    "peer": "10.0.0.2:0",
    "send-targeted": "true",
    "sending-targeted-hello": "true",
    "throttled": "false",
    "transport": "10.0.0.2",
    "used-afi": "ip",
    "vpls": "false"
  }
*/

// ResourceMplsLdpNeighbor https://help.mikrotik.com/docs/spaces/ROS/pages/121995275/LDP
func ResourceMplsLdpNeighbor() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/mpls/ldp/neighbor"),
		MetaId:           PropId(Id),

		"active_connect": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether this router initiates the TCP connection to the neighbor.",
		},
		"addresses": {
			Type:        schema.TypeList,
			Computed:    true,
			Description: "Addresses advertised by the neighbor.",
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		KeyDynamic:  PropDynamicRo,
		KeyInactive: PropInactiveRo,
		"local_transport": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Local transport address used for the session with this neighbor.",
		},
		"on_demand": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether the downstream-on-demand label distribution mode is used with this neighbor.",
		},
		"operational": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether the LDP session with this neighbor is operational.",
		},
		"passive": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether this router waits for the neighbor to initiate the TCP connection.",
		},
		"passive_wait": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether the session is waiting for the neighbor to connect.",
		},
		"path_vector_limit": {
			Type:     schema.TypeString,
			Computed: true,
			Description: "Path vector limit negotiated with this neighbor. " +
				"Reported as a number by the console, exposed as a string because the property is " +
				"absent unless the loop detection is negotiated.",
		},
		"peer": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "LDP identifier of the neighbor, in the `lsr-id:label-space` notation.",
		},
		"send_targeted": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Whether to send the targeted (extended) hello messages to the transport address, " +
				"which is what establishes a targeted LDP session.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"sending_targeted_hello": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether the targeted hello messages are currently being sent to this neighbor.",
		},
		"throttled": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether the session establishment with this neighbor is currently throttled.",
		},
		"transport": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Transport address of the neighbor a targeted LDP session is set up with. The console " +
				"grammar accepts an IPv4/IPv6 address, an interface name or a VRF name.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"used_afi": {
			Type:     schema.TypeString,
			Computed: true,
			Description: "Address families used on the session with this neighbor. " +
				"The exact value shape is not documented, exposed as a string.",
		},
		"vpls": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether the neighbor is used for a VPLS pseudowire.",
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
