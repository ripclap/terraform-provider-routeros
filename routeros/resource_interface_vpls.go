package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
	The menu was empty on the reference device (RouterOS 7.23.2):
	`GET /rest/interface/vpls` -> `[]`, so no value sample is available.

	Writable fields reported by the device:
	/console/inspect request=child path="interface,vpls,add"
	  arp  arp-timeout  bridge  bridge-cost  bridge-horizon  bridge-pvid  cisco-static-id  comment
	  copy-from  disable-running-check  disabled  mac-address  mtu  name  peer  pw-control-word
	  pw-l2mtu  pw-type  vpls-id

	Read-only fields reported by the same device:
	  bgp-signaled  bgp-vpls  bgp-vpls-prfx  cisco-bgp-signaled  dynamic  running
*/

// ResourceInterfaceVpls VPLS pseudowire interface.
// https://help.mikrotik.com/docs/spaces/ROS/pages/40992798/VPLS
func ResourceInterfaceVpls() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/vpls"),
		MetaId:           PropId(Id),

		KeyArp:        PropArpRw,
		KeyArpTimeout: PropArpTimeoutRw,
		"bgp_signaled": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether the pseudowire was signalled by BGP.",
		},
		"bgp_vpls": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The BGP VPLS instance the dynamically created pseudowire belongs to.",
		},
		"bgp_vpls_prfx": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The BGP VPLS prefix the dynamically created pseudowire was built from.",
		},
		"bridge": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The bridge the interface is dynamically added to as a bridge port.",
		},
		"bridge_cost": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Spanning tree cost of the dynamically created bridge port.",
			ValidateFunc:     validation.IntBetween(1, 200000000),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"bridge_horizon": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Bridge horizon of the dynamically created bridge port. Accepts an integer or `none`, in " +
				"which case the bridge horizon is not used.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"bridge_pvid": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Port VLAN ID (PVID) of the dynamically created bridge port.",
			ValidateFunc:     validation.IntBetween(1, 4094),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"cisco_bgp_signaled": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether the pseudowire was signalled by BGP using the Cisco style encoding.",
		},
		"cisco_static_id": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Cisco-style VPLS tunnel identifier. Used together with `pw_type` to signal a Cisco " +
				"compatible static pseudowire.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyComment: PropCommentRw,
		"disable_running_check": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Whether to detect if the interface is running or not.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyDisabled:   PropDisabledRw,
		KeyDynamic:    PropDynamicRo,
		KeyMacAddress: PropMacAddressRw("MAC address of the VPLS interface.", false),
		"mtu": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Layer3 Maximum Transmission Unit of the VPLS interface.",
			ValidateFunc:     validation.IntBetween(32, 65536),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyName: PropName("Name of the VPLS interface."),
		"peer": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The IP address of the remote peer of the pseudowire.",
			ValidateFunc:     validation.IsIPAddress,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"pw_control_word": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Enables or disables the usage of the pseudowire Control Word.",
			ValidateFunc:     validation.StringInSlice([]string{"default", "disabled", "enabled"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"pw_l2mtu": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "The L2MTU value advertised to the remote peer.",
			ValidateFunc:     validation.IntBetween(0, 65536),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"pw_type": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Pseudowire type.",
			ValidateFunc:     validation.StringInSlice([]string{"raw-ethernet", "tagged-ethernet", "vpls"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyRunning: PropRunningRo,
		"vpls_id": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "A unique number that identifies the VPLS tunnel. Written either as `AS:number` or as " +
				"`IP:number`, e.g. `65000:1` or `10.0.0.1:1`.",
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
