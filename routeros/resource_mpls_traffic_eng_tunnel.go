package routeros

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  The menu is empty on RouterOS 7.23 (ROS 7.23, MPLS not configured),
  the field set below is taken from `/console/inspect request=syntax path="mpls,traffic-eng,tunnel,add"`
  and from `print proplist=` completion on that device.

  {
    ".id": "*1",
    "affinity-exclude": "0x00000000",
    "affinity-include-all": "0x00000000",
    "affinity-include-any": "0x00000000",
    "auto-bandwidth-avg-interval": "5m",
    "auto-bandwidth-range": "0-0",
    "auto-bandwidth-reserve": "0",
    "auto-bandwidth-update-interval": "1h",
    "bandwidth": "10000000",
    "bandwidth-limit": "disabled",
    "comment": "",
    "disabled": "false",
    "forwarding": "",
    "forwarding-on": "",
    "from-address": "10.0.0.1",
    "holding-priority": "7",
    "invalid": "false",
    "name": "tun1",
    "primary": "",
    "primary-path": "tun-1-link",
    "primary-pending": "",
    "primary-retry-interval": "1m",
    "record-route": "true",
    "reoptimize-interval": "",
    "secondary": "",
    "secondary-paths": "",
    "secondary-pending": "",
    "secondary-standby": "false",
    "session": "",
    "setup-priority": "7",
    "to-address": "10.0.0.5",
    "vrf": "main"
  }
*/

// ResourceMplsTrafficEngTunnel https://help.mikrotik.com/docs/spaces/ROS/pages/40992796/Traffic+Eng
func ResourceMplsTrafficEngTunnel() *schema.Resource {
	hexMask := validation.StringMatch(regexp.MustCompile(`^(0[xX][0-9a-fA-F]{1,8}|\d+)$`),
		"value must be a decimal number or a '0x'-prefixed hexadecimal number")

	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/mpls/traffic-eng/tunnel"),
		MetaId:           PropId(Id),

		"affinity_exclude": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Do not use an interface if its `resource_class` matches any of the bits specified " +
				"here. Written as a decimal number or as a `0x`-prefixed hexadecimal number.",
			ValidateFunc:     hexMask,
			DiffSuppressFunc: HexEqual,
		},
		"affinity_include_all": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Use an interface only if its `resource_class` matches all of the bits specified here. " +
				"Written as a decimal number or as a `0x`-prefixed hexadecimal number.",
			ValidateFunc:     hexMask,
			DiffSuppressFunc: HexEqual,
		},
		"affinity_include_any": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Use an interface if its `resource_class` matches any of the bits specified here. " +
				"Written as a decimal number or as a `0x`-prefixed hexadecimal number.",
			ValidateFunc:     hexMask,
			DiffSuppressFunc: HexEqual,
		},
		"auto_bandwidth_avg_interval": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Interval over which the actual amount of the transferred data is measured to calculate " +
				"the average bandwidth of the tunnel.",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEqual,
		},
		"auto_bandwidth_range": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Range the automatically adjusted reservation is kept within, written as `Min[-Max]` " +
				"in bits per second.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"auto_bandwidth_reserve": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Percentage of the additional bandwidth reserved on top of the measured average when " +
				"the automatic bandwidth is used.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"auto_bandwidth_update_interval": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Interval during which the tunnel keeps track of the highest average rate before the " +
				"reservation is updated.",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEqual,
		},
		"bandwidth": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Amount of the bandwidth reserved for this TE tunnel. The console grammar is `Num[bps]`, " +
				"and the router reports the value back as a plain number of bits per second.",
			DiffSuppressFunc: BitsEqual,
		},
		"bandwidth_limit": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Limits the traffic sent into the tunnel as a percentage of the configured tunnel " +
				"bandwidth, or `disabled` for no limit.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"forwarding": {
			Type:     schema.TypeString,
			Computed: true,
			Description: "Forwarding state of the tunnel. " +
				"The exact value shape is not documented, exposed as a string.",
		},
		"forwarding_on": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Path the tunnel currently forwards the traffic over.",
		},
		"from_address": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Ingress address of the tunnel. When left unset RouterOS picks the address " +
				"automatically.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"holding_priority": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Priority used to decide whether this session can be preempted by another session. " +
				"`0` is the highest priority.",
			ValidateFunc:     validation.IntBetween(0, 7),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyInvalid: PropInvalidRo,
		KeyName:    PropName("Name of the traffic engineering tunnel."),
		"primary": {
			Type:     schema.TypeString,
			Computed: true,
			Description: "State of the primary path of the tunnel. " +
				"The exact value shape is not documented, exposed as a string.",
		},
		"primary_path": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Name of the primary label switching path from the `/mpls/traffic-eng/path` menu.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"primary_pending": {
			Type:     schema.TypeString,
			Computed: true,
			Description: "Pending state of the primary path. " +
				"The exact value shape is not documented, exposed as a string.",
		},
		"primary_retry_interval": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Interval after which the tunnel tries to move back to the primary path.",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEqual,
		},
		"record_route": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Whether the sender node asks to record the actual route the LSP tunnel traverses, " +
				"which is also used for the loop detection.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"reoptimize_interval": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Interval after which the tunnel re-optimizes the current path.",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEqual,
		},
		"secondary": {
			Type:     schema.TypeString,
			Computed: true,
			Description: "State of the secondary path of the tunnel. " +
				"The exact value shape is not documented, exposed as a string.",
		},
		"secondary_paths": {
			Type:     schema.TypeList,
			Optional: true,
			Description: "Names of the label switching paths from the `/mpls/traffic-eng/path` menu used when " +
				"the primary path fails. The order is significant.",
			Elem: &schema.Schema{Type: schema.TypeString},
		},
		"secondary_pending": {
			Type:     schema.TypeString,
			Computed: true,
			Description: "Pending state of the secondary path. " +
				"The exact value shape is not documented, exposed as a string.",
		},
		"secondary_standby": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Whether the secondary paths are signalled and kept up in advance so that the switchover " +
				"does not require a new reservation.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"session": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "RSVP session of the tunnel.",
		},
		"setup_priority": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Priority used to decide whether this session can preempt another session. " +
				"`0` is the highest priority.",
			ValidateFunc:     validation.IntBetween(0, 7),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"to_address": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Address of the remote end of the TE tunnel.",
			ValidateFunc:     validation.IsIPAddress,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
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
