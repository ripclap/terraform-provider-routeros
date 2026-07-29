package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
	The menu was empty on the reference device (RouterOS 7.23.2), so no value sample is available.
*/

// ResourceIpPacking MikroTik Packet Packer Protocol (M3P) per interface.
// https://help.mikrotik.com/docs/spaces/ROS/pages/139526171/IP+packing
func ResourceIpPacking() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/packing"),
		MetaId:           PropId(Id),

		"aggregated_size": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Size of the aggregated packet that packing will try to achieve before sending.",
			ValidateFunc:     validation.IntBetween(20, 16384),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyDisabled:  PropDisabledRw,
		KeyInterface: PropInterfaceRw,
		"packing": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "How the outgoing packets of this interface are processed." +
				"\n  * none - no processing." +
				"\n  * simple - aggregate several small packets into one large packet." +
				"\n  * compress-headers - aggregate and compress the IP headers." +
				"\n  * compress-all - aggregate and compress both headers and payload.",
			ValidateFunc: validation.StringInSlice([]string{"compress-all", "compress-headers", "none", "simple"},
				false),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"unpacking": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "How the incoming packets of this interface are processed." +
				"\n  * none - no processing." +
				"\n  * simple - unpack aggregated packets." +
				"\n  * compress-headers - unpack and decompress the IP headers." +
				"\n  * compress-all - unpack and decompress both headers and payload.",
			ValidateFunc: validation.StringInSlice([]string{"compress-all", "compress-headers", "none", "simple"},
				false),
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
