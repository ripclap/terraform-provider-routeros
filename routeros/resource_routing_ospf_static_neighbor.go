package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  {
    ".id": "*1",
    "address": "",
    "area": "",
    "comment": "",
    "disabled": "false",
    "inactive": "false",
    "instance-id": "",
    "poll-interval": ""
  }
*/

// ResourceRoutingOspfStaticNeighbor https://help.mikrotik.com/docs/spaces/ROS/pages/9863229/OSPF
func ResourceRoutingOspfStaticNeighbor() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/routing/ospf/static-neighbor"),
		MetaId:           PropId(Id),

		"address": {
			Type:     schema.TypeString,
			Required: true,
			Description: "The address of the neighbor. For OSPFv3 the link-local address of the neighbor together " +
				"with the interface it is reachable through is used, for example `fe80::1%ether1`.",
		},
		"area": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The OSPF area the neighbor belongs to.",
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		KeyInactive: PropInactiveRo,
		"instance_id": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "The OSPF instance id used in the packets sent to this neighbor.",
			ValidateFunc:     validation.IntBetween(0, 255),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"poll_interval": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The interval between the HELLO packets sent to a neighbor that is currently down. It is " +
				"normally much longer than the hello interval of the interface.",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEqual,
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
