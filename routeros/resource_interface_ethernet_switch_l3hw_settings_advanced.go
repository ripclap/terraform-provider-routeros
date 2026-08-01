package routeros

import (
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
	{
	  "neigh-discovery-burst-delay": "300ms",
	  "neigh-discovery-burst-limit": "64",
	  "neigh-discovery-interval": "1m37s",
	  "neigh-dump-retries": "3",
	  "neigh-keepalive-interval": "15s",
	  "partial-offload-chunk": "1024",
	  "route-index-delay-max": "10s",
	  "route-index-delay-min": "1s",
	  "route-queue-limit-high": "256",
	  "route-queue-limit-low": "0",
	  "shwp-reset-counter": "128"
	}
*/

// ResourceInterfaceEthernetSwitchL3HwSettingsAdvanced Advanced tuning of the L3 hardware offloading engine.
// https://help.mikrotik.com/docs/spaces/ROS/pages/62390319/L3+Hardware+Offloading
func ResourceInterfaceEthernetSwitchL3HwSettingsAdvanced() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/ethernet/switch/l3hw-settings/advanced"),
		MetaId:           PropId(Id),

		"neigh_discovery_burst_delay": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Delay between subsequent ARP/ND request bursts. The minimum value is `10ms`.",
			DiffSuppressFunc: TimeEqualU(time.Millisecond),
		},
		"neigh_discovery_burst_limit": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Maximum number of ARP/ND requests that are sent simultaneously.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"neigh_discovery_interval": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Interval for sending ARP/ND requests used to check whether the offloaded hosts are still " +
				"active. The minimum value is `30s`.",
			DiffSuppressFunc: TimeEqual,
		},
		"neigh_dump_retries": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Number of retries used when the neighbor table is dumped." +
				"\nThis property is exposed by RouterOS 7.23 but is not covered by the MikroTik " +
				"documentation; only the name and the default value (3) are known.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"neigh_keepalive_interval": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Keepalive interval of the offloaded neighbor (host) entries. The minimum value is `5s`.",
			DiffSuppressFunc: TimeEqual,
		},
		"partial_offload_chunk": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Minimum number of routes that are added incrementally in Partial Offloading mode.",
			ValidateFunc:     validation.IntAtLeast(16),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"route_index_delay_max": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Maximum delay between route processing and offloading.",
			DiffSuppressFunc: TimeEqual,
		},
		"route_index_delay_min": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Minimum delay between route processing and offloading.",
			DiffSuppressFunc: TimeEqual,
		},
		"route_queue_limit_high": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Threshold at which the driver stops route indexing.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"route_queue_limit_low": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Threshold at which the driver re-enables route indexing.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"shwp_reset_counter": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Reset the Shortest HW Prefix and try full route table offloading after this amount of " +
				"changes.",
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
