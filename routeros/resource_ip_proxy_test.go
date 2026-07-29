package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIpProxy = "routeros_ip_proxy.test_proxy_x"

// `/ip/proxy` is a settings singleton (no CheckDestroy); left disabled so no listening port is opened.
func TestAccIpProxyTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				Steps: []resource.TestStep{
					{
						Config: testAccIpProxyConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpProxy),
							resource.TestCheckResourceAttr(testIpProxy, "enabled", "false"),
							resource.TestCheckResourceAttr(testIpProxy, "port", "8081"),
							resource.TestCheckResourceAttr(testIpProxy, "src_address", "192.0.2.79"),
							resource.TestCheckResourceAttr(testIpProxy, "parent_proxy", "192.0.2.80"),
							resource.TestCheckResourceAttr(testIpProxy, "parent_proxy_port", "3128"),
							resource.TestCheckResourceAttr(testIpProxy, "cache_administrator", "test_proxy_x_admin"),
							resource.TestCheckResourceAttr(testIpProxy, "max_cache_size", "none"),
							resource.TestCheckResourceAttr(testIpProxy, "max_cache_object_size", "1024"),
							resource.TestCheckResourceAttr(testIpProxy, "max_client_connections", "500"),
							resource.TestCheckResourceAttr(testIpProxy, "max_server_connections", "500"),
							resource.TestCheckResourceAttr(testIpProxy, "max_fresh_time", "2d"),
							resource.TestCheckResourceAttr(testIpProxy, "anonymous", "true"),
							resource.TestCheckResourceAttr(testIpProxy, "always_from_cache", "true"),
							resource.TestCheckResourceAttr(testIpProxy, "serialize_connections", "true"),
							resource.TestCheckResourceAttr(testIpProxy, "cache_hit_dscp", "8"),
							resource.TestCheckResourceAttr(testIpProxy, "cache_on_disk", "false"),
						),
					},
					{
						// Restores the RouterOS defaults of the menu.
						Config: testAccIpProxyRestoreConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpProxy),
							resource.TestCheckResourceAttr(testIpProxy, "port", "8080"),
							resource.TestCheckResourceAttr(testIpProxy, "src_address", "::"),
							resource.TestCheckResourceAttr(testIpProxy, "parent_proxy", "::"),
							resource.TestCheckResourceAttr(testIpProxy, "parent_proxy_port", "0"),
							resource.TestCheckResourceAttr(testIpProxy, "cache_administrator", "webmaster"),
							resource.TestCheckResourceAttr(testIpProxy, "max_cache_size", "unlimited"),
							resource.TestCheckResourceAttr(testIpProxy, "max_cache_object_size", "2048"),
							resource.TestCheckResourceAttr(testIpProxy, "max_fresh_time", "3d"),
							resource.TestCheckResourceAttr(testIpProxy, "anonymous", "false"),
						),
					},
				},
			})

		})
	}
}

func testAccIpProxyConfig() string {
	return providerConfig + `
resource "routeros_ip_proxy" "test_proxy_x" {
	enabled                = false
	port                   = 8081
	src_address            = "192.0.2.79"
	parent_proxy           = "192.0.2.80"
	parent_proxy_port      = 3128
	cache_administrator    = "test_proxy_x_admin"
	cache_on_disk          = false
	max_cache_size         = "none"
	max_cache_object_size  = 1024
	max_client_connections = 500
	max_server_connections = 500
	max_fresh_time         = "2d"
	anonymous              = true
	always_from_cache      = true
	serialize_connections  = true
	cache_hit_dscp         = 8
}
`
}

func testAccIpProxyRestoreConfig() string {
	return providerConfig + `
resource "routeros_ip_proxy" "test_proxy_x" {
	enabled                = false
	port                   = 8080
	src_address            = "::"
	parent_proxy           = "::"
	parent_proxy_port      = 0
	cache_administrator    = "webmaster"
	cache_on_disk          = false
	max_cache_size         = "unlimited"
	max_cache_object_size  = 2048
	max_client_connections = 600
	max_server_connections = 600
	max_fresh_time         = "3d"
	anonymous              = false
	always_from_cache      = false
	serialize_connections  = false
	cache_hit_dscp         = 4
}
`
}
