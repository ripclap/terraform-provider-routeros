package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
	{
	  "latency-distribution-max": "100us",
	  "latency-distribution-measure-interval": "0-853.ns",
	  "latency-distribution-samples": "64",
	  "measure-out-of-order": "false",
	  "running": "false",
	  "stats-samples-to-keep": "100",
	  "test-id": "0"
	}
*/

// ResourceToolTrafficGenerator Global traffic generator settings.
// https://help.mikrotik.com/docs/spaces/ROS/pages/128221376/Traffic+Generator
func ResourceToolTrafficGenerator() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/tool/traffic-generator"),
		MetaId:           PropId(Id),
		MetaSkipFields: PropSkipFields("latency_distribution_measure_interval", "latency_distribution_samples",
			"running"),

		"latency_distribution_max": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Upper bound of the latency distribution histogram, for example `100us`. Latency values " +
				"above this bound are counted in the last bucket." +
				"\n> The value is a RouterOS time interval and may use the `us`, `ms`, `s` postfixes. It is compared " +
				"as an opaque string, so use the same notation that the router reports back.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"measure_out_of_order": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Whether out of order packets are detected and counted during a test.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"stats_samples_to_keep": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "How many statistics samples are kept in memory (0..4294967295).",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"test_id": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Identifier that is embedded into the generated packets so that the receiving side can " +
				"tell several parallel tests apart.",
			ValidateFunc:     validation.IntBetween(0, 255),
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
