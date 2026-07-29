package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceIpCloudV0 is the schema as it stood while ddns_enabled was a string.
// Kept only so the v0 -> v1 state upgrader can derive the prior type.
func ResourceIpCloudV0() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/cloud"),
		MetaId:           PropId(Id),
		MetaSkipFields: PropSkipFields("vpn_dns_name", "vpn_interface", "vpn_peer_private_key", "vpn_peer_public_key",
			"vpn_port", "vpn_private_key", "vpn_public_key", "vpn_relay_addressess", "vpn_relay_addressess_ipv6",
			"vpn_relay_codes", "vpn_relay_ipv4_status", "vpn_relay_ipv6_status", "vpn_relay_regions", "vpn_relay_rtts",
			"vpn_status", "vpn_wireguard_client_config", "vpn_wireguard_client_config_qrcode"),

		"back_to_home_vpn": {
			Type:             schema.TypeString,
			Optional:         true,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"ddns_enabled": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"ddns_update_interval": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "none",
		},
		"dns_name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"public_address": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"public_address_ipv6": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"status": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"update_time": {
			Type:             schema.TypeString,
			Optional:         true,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"vpn_prefer_relay_code": {
			Type:             schema.TypeString,
			Optional:         true,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"warning": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}

	return &schema.Resource{Schema: resSchema}
}
