package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
  "router-advertisement-ignored-options": "",
  "router-advertisement-route-distance": "0"
}

  The accepted keywords of "router-advertisement-ignored-options" were enumerated on the
  device with the console completion engine:

    /console/inspect request=completion path="ipv6,nd,settings"
                     input="set router-advertisement-ignored-options="
      -> dns, mtu

  which agrees with help.mikrotik.com (selection: "mtu, dns"). The same page documents
  "router-advertisement-route-distance" as integer 0..255, default 0.
*/

// ResourceIPv6NdSettings https://help.mikrotik.com/docs/spaces/ROS/pages/40992815/IPv6+Neighbor+Discovery
func ResourceIPv6NdSettings() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ipv6/nd/settings"),
		MetaId:           PropId(Id),

		"router_advertisement_ignored_options": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"dns", "mtu"}, false),
			},
			Description: "Neighbor Discovery options that are disregarded in the received router advertisements." +
				"\n  * dns - ignore the RDNSS/DNSSL options;" +
				"\n  * mtu - ignore the MTU option.",
		},
		"router_advertisement_route_distance": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Distance (administrative metric) applied to the default routes installed from received " +
				"router advertisements. *Default: `0`*",
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
