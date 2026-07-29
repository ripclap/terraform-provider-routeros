package routeros

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  The menu is empty on the reference device (ROS 7.23.2, MPLS not configured),
  the field set below is taken from `/console/inspect request=syntax path="mpls,traffic-eng,interface,add"`
  and from `print proplist=` completion on that device.

  {
    ".id": "*1",
    "bandwidth": "100000000",
    "blockade-k-factor": "4",
    "comment": "",
    "disabled": "false",
    "down-flood-thresholds": "100,99,98,97,96,95,90,85,80,75,60,45,30,15",
    "igp-flood-period": "3m",
    "interface": "ether1",
    "invalid": "false",
    "k-factor": "3",
    "lih": "",
    "local-address-ip": "10.0.0.1",
    "local-address-ip6": "",
    "refresh-time": "30s",
    "remaining-bw": "",
    "remaining-bw-prios": "",
    "resource-class": "0x00000000",
    "te-metric": "0",
    "up-flood-thresholds": "15,30,45,60,75,80,85,90,95,96,97,98,99,100",
    "use-udp": "false"
  }
*/

// ResourceMplsTrafficEngInterface https://help.mikrotik.com/docs/spaces/ROS/pages/40992796/Traffic+Eng
func ResourceMplsTrafficEngInterface() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/mpls/traffic-eng/interface"),
		MetaId:           PropId(Id),

		"bandwidth": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Amount of the bandwidth announced as available for the RSVP-TE reservations on this " +
				"interface. The console grammar is `Num[bps]`, and the router reports the value back as a plain " +
				"number of bits per second.",
			DiffSuppressFunc: BitsEqual,
		},
		// blockade_k_factor, retransmit_multiplier and te_metric are uint32; string-typed to avoid 32-bit overflow.
		"blockade_k_factor": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Multiplier used to calculate the lifetime of the RSVP blockade state (RFC 2205).",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"down_flood_thresholds": {
			Type:     schema.TypeList,
			Optional: true,
			Description: "Reserved bandwidth thresholds, in percent, that trigger an IGP traffic engineering " +
				"update when the amount of the reserved bandwidth decreases. The order is significant.",
			Elem: &schema.Schema{Type: schema.TypeString},
		},
		"igp_flood_period": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Maximum interval between the periodic IGP traffic engineering updates advertised for " +
				"this interface.",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEqual,
		},
		"interface": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of the interface RSVP-TE is enabled on.",
		},
		KeyInvalid: PropInvalidRo,
		"k_factor": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "RSVP refresh multiplier K used to calculate the state lifetime (RFC 2205).",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"lih": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Logical Interface Handle assigned to this interface by RSVP.",
		},
		"local_address_ip": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Local IPv4 address used for the RSVP messages sent on this interface.",
		},
		"local_address_ip6": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Local IPv6 address used for the RSVP messages sent on this interface.",
		},
		"refresh_time": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Interval between the RSVP refresh messages sent for the sessions on this interface.",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEqual,
		},
		"remaining_bw": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Bandwidth still available for the new reservations on this interface.",
		},
		"remaining_bw_prios": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Bandwidth still available for the new reservations, per setup priority.",
		},
		"resource_class": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Resource class (administrative group, link colour) bit mask of this interface, matched " +
				"by the `affinity_include_any`, `affinity_include_all` and `affinity_exclude` properties of the " +
				"tunnels and paths. Written as a decimal number or as a `0x`-prefixed hexadecimal number.",
			ValidateFunc: validation.StringMatch(regexp.MustCompile(`^(0[xX][0-9a-fA-F]{1,8}|\d+)$`),
				"value must be a decimal number or a '0x'-prefixed hexadecimal number"),
			DiffSuppressFunc: HexEqual,
		},
		"te_metric": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Traffic engineering metric of this interface, used by CSPF instead of the IGP metric.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"up_flood_thresholds": {
			Type:     schema.TypeList,
			Optional: true,
			Description: "Reserved bandwidth thresholds, in percent, that trigger an IGP traffic engineering " +
				"update when the amount of the reserved bandwidth increases. The order is significant.",
			Elem: &schema.Schema{Type: schema.TypeString},
		},
		"use_udp": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Whether to send the RSVP messages encapsulated in UDP instead of raw IP.",
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
