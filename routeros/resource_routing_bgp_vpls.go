package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  {
    ".id": "*1",
    "bridge": "",
    "bridge-cost": "",
    "bridge-horizon": "",
    "bridge-pvid": "",
    "cisco-id": "",
    "comment": "",
    "current-peers": "",
    "disabled": "false",
    "export-route-targets": "",
    "import-route-targets": "",
    "inactive": "false",
    "interface-list": "",
    "local-pref": "",
    "name": "",
    "pw-control-word": "",
    "pw-l2mtu": "",
    "pw-type": "",
    "rd": "",
    "site-id": "",
    "vrf": "main"
  }
*/

// ResourceRoutingBgpVpls A BGP signaled VPLS instance. The pseudowires to the remote sites are discovered and
// signaled over BGP (RFC 4761) and attached to the configured bridge.
// https://help.mikrotik.com/docs/spaces/ROS/pages/331612228/routing+bgp
func ResourceRoutingBgpVpls() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/routing/bgp/vpls"),
		MetaId:           PropId(Id),

		"bridge": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Name of the bridge the dynamically created pseudowire interfaces are added to as ports.",
		},
		"bridge_cost": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "The spanning tree path cost assigned to the dynamically created bridge ports.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"bridge_horizon": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The split horizon group of the dynamically created bridge ports. Ports that share the " +
				"same horizon value do not forward traffic to each other. The value is either `none` or a number.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"bridge_pvid": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "The port VLAN id assigned to the dynamically created bridge ports.",
			ValidateFunc:     validation.IntBetween(1, 4094),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"cisco_id": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The VPLS identifier used for the interoperability with the Cisco style BGP signaled " +
				"VPLS, written as `RouteDistinguisher:VPLSId`.",
		},
		KeyComment: PropCommentRw,
		"current_peers": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The peers this VPLS instance currently has a signaled pseudowire with.",
		},
		KeyDisabled: PropDisabledRw,
		"export_route_targets": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "The route targets attached to the VPLS NLRI advertised by this instance.",
		},
		"import_route_targets": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "The route targets a received VPLS NLRI has to carry to be accepted into this instance.",
		},
		KeyInactive: PropInactiveRo,
		"interface_list": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Name of the interface list the dynamically created pseudowire interfaces are added to.",
		},
		"local_pref": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "The LOCAL_PREF attribute value of the VPLS NLRI advertised by this instance.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyName: PropName("Name of the VPLS instance."),
		"pw_control_word": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Whether the pseudowire control word is used. `default` follows the global MPLS " +
				"settings.",
			ValidateFunc:     validation.StringInSlice([]string{"default", "disabled", "enabled"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"pw_l2mtu": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "The layer2 MTU signaled for the pseudowires of this instance.",
			ValidateFunc:     validation.IntBetween(0, 65535),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"pw_type": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "The pseudowire encapsulation type signaled for this instance.",
			ValidateFunc: validation.StringInSlice([]string{"raw-ethernet", "tagged-ethernet", "vpls"}, false),
		},
		"rd": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The route distinguisher prepended to the VPLS NLRI advertised by this instance, written " +
				"as `ASN:number` or `IP:number`.",
		},
		"site_id": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "The VPLS site identifier of this router inside the instance. It has to be unique among " +
				"all the sites of the VPLS.",
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
