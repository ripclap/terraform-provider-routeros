package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
	The menu was empty on the RouterOS 7.23:
	`GET /rest/interface/mesh` -> `[]`, so no value sample is available.

	Writable fields reported by the device:
	/console/inspect request=child path="interface,mesh,add"
	  admin-mac  arp  arp-timeout  auto-mac  comment  copy-from  disabled  hwmp-default-hoplimit
	  hwmp-prep-lifetime  hwmp-preq-destination-only  hwmp-preq-reply-and-forward  hwmp-preq-retries
	  hwmp-preq-waiting-time  hwmp-rann-interval  hwmp-rann-lifetime  hwmp-rann-propagation-delay
	  mesh-portal  mtu  name  reoptimize-paths

	Read-only fields reported by the same device: mac-address, running

	The HWMP+ timers are reported by RouterOS as time values, so they are modelled as strings with
	a time aware diff.

	HWMP+ mesh has no page on help.mikrotik.com; the property descriptions below come from
	the legacy MikroTik HWMP+ mesh manual. The field names, the value types and the enumerations
	were taken from the device.
*/

// ResourceInterfaceMesh HWMP+ mesh interface.
func ResourceInterfaceMesh() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/mesh"),
		MetaId:           PropId(Id),

		"admin_mac": PropMacAddressRw("Static MAC address of the mesh interface, used when `auto_mac` is disabled.",
			false),
		KeyArp:        PropArpRw,
		KeyArpTimeout: PropArpTimeoutRw,
		"auto_mac": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Whether to use the MAC address of the first mesh port instead of the address configured " +
				"in `admin_mac`.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"hwmp_default_hoplimit": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "The TTL put into the HWMP+ messages originated by this device.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"hwmp_prep_lifetime": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Lifetime of the reverse route that is created by a received HWMP+ PREP message.",
			DiffSuppressFunc: TimeEqual,
		},
		"hwmp_preq_destination_only": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Whether only the destination of a PREQ message is allowed to answer it. When disabled, " +
				"intermediate nodes with a valid path may answer instead.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"hwmp_preq_reply_and_forward": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Whether an intermediate node that answers a PREQ message also forwards the message to the " +
				"destination.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"hwmp_preq_retries": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "How many times a PREQ message is retried before the destination is considered unreachable.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"hwmp_preq_waiting_time": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "How long to wait for a PREP answer before retrying the PREQ message.",
			DiffSuppressFunc: TimeEqual,
		},
		"hwmp_rann_interval": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "How often the mesh portal originates a RANN (root announcement) message.",
			DiffSuppressFunc: TimeEqual,
		},
		"hwmp_rann_lifetime": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Lifetime of the route to the mesh portal created by a received RANN message.",
			DiffSuppressFunc: TimeEqual,
		},
		"hwmp_rann_propagation_delay": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "How long a node waits before propagating a received RANN message further.",
			DiffSuppressFunc: TimeEqual,
		},
		KeyMacAddress: PropMacAddressRo,
		"mesh_portal": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Whether this device acts as a mesh portal, i.e. announces itself as a gateway out of the " +
				"mesh with RANN messages.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"mtu": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Layer3 Maximum Transmission Unit of the mesh interface.",
			ValidateFunc:     Validation64k,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyName: PropName("Name of the mesh interface."),
		"reoptimize_paths": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Whether to periodically look for better paths to the known destinations even when the " +
				"current paths still work.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyRunning: PropRunningRo,
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
