package routeros

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
	{
	  ".id": "*3",
	  "byte-use": "0",
	  "invalid": "false",
	  "map": "default",
	  "name": "sfp-sfpplus1",
	  "packet-use": "0",
	  "pfc": "disabled",
	  "pfc-paused-tc": "",
	  "profile": "default",
	  "queue1-byte-cap": "54784",
	  "queue1-packet-cap": "199",
	  "queue1-shared-byte-cap": "687104",
	  "queue1-shared-packet-cap": "2555",
	  "trust-l2": "ignore",
	  "trust-l3": "ignore",
	  "tx-manager": "default"
	  ... (per-queue and per-PFC-class counters omitted)
	}

	The entries of this menu are created by the switch driver, one per switch port; they can only be
	changed, never added or removed. Writable fields reported by the reference device
	(RouterOS 7.23.2):

	/console/inspect request=child path="interface,ethernet,switch,qos,port,set"
	  egress-rate-queue0 .. egress-rate-queue7  map  pfc  profile  trust-l2  trust-l3  tx-manager

	Everything else the menu reports is a statistics counter and is skipped.
*/

// ResourceInterfaceEthernetSwitchQosPort Per switch port QoS configuration.
// https://help.mikrotik.com/docs/spaces/ROS/pages/189497483/Quality+of+Service
func ResourceInterfaceEthernetSwitchQosPort() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/ethernet/switch/qos/port"),
		MetaId:           PropId(Id),
		MetaSkipFields: PropSkipFields("byte_max", "byte_use", "name", "packet_use", "pfc_paused_tc", "pfc_rx", "pfc_tx",
			"pfc_unknown", "pfc0_pause_threshold", "pfc0_resume_threshold", "pfc0_use", "pfc1_pause_threshold",
			"pfc1_resume_threshold", "pfc1_use", "pfc2_pause_threshold", "pfc2_resume_threshold", "pfc2_use",
			"pfc3_pause_threshold", "pfc3_resume_threshold", "pfc3_use", "pfc4_pause_threshold", "pfc4_resume_threshold",
			"pfc4_use", "pfc5_pause_threshold", "pfc5_resume_threshold", "pfc5_use", "pfc6_pause_threshold",
			"pfc6_resume_threshold", "pfc6_use", "pfc7_pause_threshold", "pfc7_resume_threshold", "pfc7_use",
			"queue0_byte_cap", "queue0_byte_max", "queue0_byte_use", "queue0_packet_cap", "queue0_packet_use",
			"queue0_shared_byte_cap", "queue0_shared_packet_cap", "queue1_byte_cap", "queue1_byte_max", "queue1_byte_use",
			"queue1_packet_cap", "queue1_packet_use", "queue1_shared_byte_cap", "queue1_shared_packet_cap",
			"queue2_byte_cap", "queue2_byte_max", "queue2_byte_use", "queue2_packet_cap", "queue2_packet_use",
			"queue2_shared_byte_cap", "queue2_shared_packet_cap", "queue3_byte_cap", "queue3_byte_max", "queue3_byte_use",
			"queue3_packet_cap", "queue3_packet_use", "queue3_shared_byte_cap", "queue3_shared_packet_cap",
			"queue4_byte_cap", "queue4_byte_max", "queue4_byte_use", "queue4_packet_cap", "queue4_packet_use",
			"queue4_shared_byte_cap", "queue4_shared_packet_cap", "queue5_byte_cap", "queue5_byte_max", "queue5_byte_use",
			"queue5_packet_cap", "queue5_packet_use", "queue5_shared_byte_cap", "queue5_shared_packet_cap",
			"queue6_byte_cap", "queue6_byte_max", "queue6_byte_use", "queue6_packet_cap", "queue6_packet_use",
			"queue6_shared_byte_cap", "queue6_shared_packet_cap", "queue7_byte_cap", "queue7_byte_max", "queue7_byte_use",
			"queue7_packet_cap", "queue7_packet_use", "queue7_shared_byte_cap", "queue7_shared_packet_cap", "tx_bytes",
			"tx_drop_byte", "tx_drop_packet", "tx_drop_queue0_byte", "tx_drop_queue0_packet", "tx_drop_queue1_byte",
			"tx_drop_queue1_packet", "tx_drop_queue2_byte", "tx_drop_queue2_packet", "tx_drop_queue3_byte",
			"tx_drop_queue3_packet", "tx_drop_queue4_byte", "tx_drop_queue4_packet", "tx_drop_queue5_byte",
			"tx_drop_queue5_packet", "tx_drop_queue6_byte", "tx_drop_queue6_packet", "tx_drop_queue7_byte",
			"tx_drop_queue7_packet", "tx_packet", "tx_queue0_byte", "tx_queue0_packet", "tx_queue1_byte",
			"tx_queue1_packet", "tx_queue2_byte", "tx_queue2_packet", "tx_queue3_byte", "tx_queue3_packet",
			"tx_queue4_byte", "tx_queue4_packet", "tx_queue5_byte", "tx_queue5_packet", "tx_queue6_byte",
			"tx_queue6_packet", "tx_queue7_byte", "tx_queue7_packet"),

		"egress_rate_queue0": propQosPortEgressRate(0),
		"egress_rate_queue1": propQosPortEgressRate(1),
		"egress_rate_queue2": propQosPortEgressRate(2),
		"egress_rate_queue3": propQosPortEgressRate(3),
		"egress_rate_queue4": propQosPortEgressRate(4),
		"egress_rate_queue5": propQosPortEgressRate(5),
		"egress_rate_queue6": propQosPortEgressRate(6),
		"egress_rate_queue7": propQosPortEgressRate(7),
		KeyInvalid:           PropInvalidRo,
		"map": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The priority-to-profile mapping table used for the ingress packets of this port.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyName: PropNameOptional("Port name."),
		"pfc": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The Priority-based Flow Control profile that controls the ingress priority traffic of this " +
				"port.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"profile": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The QoS profile assigned by default to the ingress packets of this port.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyRunning: PropRunningRo,
		"switch": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Name of the switch the port belongs to.",
		},
		"trust_l2": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Whether to trust the Layer-2 VLAN header (PCP field) of the ingress packets.",
			ValidateFunc:     validation.StringInSlice([]string{"ignore", "keep", "trust"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"trust_l3": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Whether to trust the Layer-3 IP header (DSCP field) of the ingress packets.",
			ValidateFunc:     validation.StringInSlice([]string{"ignore", "keep", "trust"}, false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"tx_manager": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The transmission manager responsible for the packet egress of this port.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
	}

	resCreateUpdate := func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		item, metadata := TerraformResourceDataToMikrotik(resSchema, d)

		res, err := ReadItems(&ItemId{Name, d.Get("name").(string)}, metadata.Path, m.(Client))
		if err != nil {
			// API/REST client error.
			ColorizedDebug(ctx, fmt.Sprintf(ErrorMsgPatch, err))
			return diag.FromErr(err)
		}

		// Resource not found.
		if len(*res) == 0 {
			d.SetId("")
			ColorizedDebug(ctx, fmt.Sprintf(ErrorMsgPatch, err))
			return diag.FromErr(errorNoLongerExists)
		}

		d.SetId((*res)[0].GetID(Id))
		item[".id"] = d.Id()

		var resUrl string
		if m.(Client).GetTransport() == TransportREST {
			resUrl = "/set"
		}

		err = m.(Client).SendRequest(crudPost, &URL{Path: metadata.Path + resUrl}, item, nil)
		if err != nil {
			return diag.FromErr(err)
		}

		return ResourceRead(ctx, resSchema, d, m)
	}

	return &schema.Resource{
		CreateContext: resCreateUpdate,
		ReadContext:   DefaultRead(resSchema),
		UpdateContext: resCreateUpdate,
		DeleteContext: DefaultSystemDelete(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: ImportStateCustomContext(resSchema),
		},

		Schema: resSchema,
	}
}

// propQosPortEgressRate builds the schema of the per-queue egress rate limit of a switch port.
func propQosPortEgressRate(queue int) *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Description: fmt.Sprintf("Egress traffic limitation in bits per second for queue%d. Accepts an integer "+
			"with an optional `k`, `M` or `G` suffix.", queue),
		DiffSuppressFunc: BitsEqual,
	}
}
