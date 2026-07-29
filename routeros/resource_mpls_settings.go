package routeros

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  {
    "allow-fast-path": "yes",
    "dynamic-label-range": "16-1048575",
    "mpls-fast-path-bytes": "0",
    "mpls-fast-path-packets": "0",
    "propagate-ttl": "yes"
  }
*/

// ResourceMplsSettings https://help.mikrotik.com/docs/spaces/ROS/pages/128974876/MPLS+MTU%2C+Forwarding+and+Label+Bindings
func ResourceMplsSettings() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/mpls/settings"),
		MetaId:           PropId(Id),
		MetaSkipFields:   PropSkipFields("mpls_fast_path_bytes", "mpls_fast_path_packets"),

		"allow_fast_path": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Whether to allow MPLS FastPath processing.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"dynamic_label_range": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Range of the label numbers used for the dynamic label allocation, written as `Start-End` " +
				"or as a single value. Valid label numbers are 16..1048575.",
			ValidateFunc: validation.StringMatch(regexp.MustCompile(`^\d+(-\d+)?$`),
				"value must be a label number or a 'Start-End' label range, for example '16-1048575'"),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"propagate_ttl": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Whether to copy the TTL value from the IP header to the MPLS header.",
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
