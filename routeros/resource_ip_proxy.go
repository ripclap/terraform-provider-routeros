package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
{
  "always-from-cache": "false",
  "anonymous": "false",
  "cache-administrator": "webmaster",
  "cache-hit-dscp": "4",
  "cache-on-disk": "false",
  "cache-path": "web-proxy",
  "enabled": "false",
  "max-cache-object-size": "2048",
  "max-cache-size": "unlimited",
  "max-client-connections": "600",
  "max-fresh-time": "3d",
  "max-server-connections": "600",
  "parent-proxy": "::",
  "parent-proxy-port": "0",
  "port": "8080",
  "serialize-connections": "false",
  "src-address": "::"
}
*/

// ResourceIpProxy Web proxy service settings.
// https://help.mikrotik.com/docs/spaces/ROS/pages/132350000/Proxy
func ResourceIpProxy() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/ip/proxy"),
		MetaId:           PropId(Id),

		"always_from_cache": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Ignore the client refresh requests if the cached content is considered fresh.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"anonymous": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Do not pass the client's IP address on in the `X-Forwarded-For` header.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"cache_administrator": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Administrator e-mail address displayed on the proxy error pages.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"cache_hit_dscp": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "DSCP value the proxy sets on the packets it serves from the cache.",
			ValidateFunc:     validation.IntBetween(0, 63),
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"cache_on_disk": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Store the cache on the router's disk instead of in RAM.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"cache_path": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Directory used for the on-disk cache.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		KeyEnabled: {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Enables the web proxy service.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"max_cache_object_size": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Maximum size of a single cached object, in KiB.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"max_cache_size": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Total cache capacity: `none`, `unlimited` or a size in KiB.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"max_client_connections": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Maximum number of simultaneous client connections that are accepted.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"max_fresh_time": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Maximum period an object is kept in the cache.",
			DiffSuppressFunc: TimeEqual,
		},
		"max_server_connections": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "Maximum number of simultaneous connections towards the servers.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"parent_proxy": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "IPv4 or IPv6 address of the upstream (parent) proxy server.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"parent_proxy_port": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "TCP port the upstream (parent) proxy listens on.",
			ValidateFunc:     Validation64k,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"port": {
			Type:             schema.TypeInt,
			Optional:         true,
			Description:      "TCP port the proxy service listens on.",
			ValidateFunc:     Validation64k,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"serialize_connections": {
			Type:             schema.TypeBool,
			Optional:         true,
			Description:      "Process the persistent client connections one after another instead of in parallel.",
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
		},
		"src_address": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "Local address used as the source of the connections the proxy makes towards the servers.",
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
