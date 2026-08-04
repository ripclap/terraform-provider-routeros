package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
  ".id": "*8",
  "address": "time.cloudflare.com",
  "auth-key": "none",
  "comment": "Cloudflare NTP",
  "disabled": "false",
  "dynamic": "false",
  "iburst": "true",
  "max-poll": "10",
  "min-poll": "6",
  "resolved-address": "162.159.200.123"
}
*/

// ResourceSystemNtpClientServer https://help.mikrotik.com/docs/spaces/ROS/pages/40992869/NTP
func ResourceSystemNtpClientServer() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/system/ntp/client/servers"),
		MetaId:           PropId(Id),

		"address": {
			Type:     schema.TypeString,
			Required: true,
			Description: "The address of the NTP server. The following formats are accepted: FQDN, `ipv4`, " +
				"`ipv4@vrf`, `ipv6`, `ipv6@vrf`, `ipv6-linklocal%interface`.",
		},
		"auth_key": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The NTP symmetric key used for the authentication between the NTP client and the server. " +
				"References an entry of `/system/ntp/key`; `none` disables the authentication.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
			Sensitive:        true,
		},
		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		KeyDynamic:  PropDynamicRo,
		"iburst": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Sends a burst of packets when the association is first initialized, which speeds up the " +
				"initial synchronization.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"max_poll": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "The maximum polling interval, as a power of two in seconds (`10` means 1024 seconds). " +
				"Must not be lower than `min_poll`.",
			ValidateFunc:     validation.IntBetween(3, 17),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"min_poll": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "The minimum polling interval, as a power of two in seconds (`6` means 64 seconds). " +
				"Must not be higher than `max_poll`.",
			ValidateFunc:     validation.IntBetween(3, 17),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"resolved_address": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The IP address the `address` was resolved to, when an FQDN is used.",
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
