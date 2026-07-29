package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIpProxyCacheAddress = "routeros_ip_proxy_cache.test_ip_proxy_cache"

func TestAccIpProxyCacheTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ip/proxy/cache", "routeros_ip_proxy_cache"),
				Steps: []resource.TestStep{
					{
						Config: testAccIpProxyCacheConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpProxyCacheAddress),
							resource.TestCheckResourceAttr(testIpProxyCacheAddress, "action", "deny"),
							resource.TestCheckResourceAttr(testIpProxyCacheAddress, "comment", "test_ip_proxy_cache"),
							resource.TestCheckResourceAttr(testIpProxyCacheAddress, "disabled", "true"),
							resource.TestCheckResourceAttr(testIpProxyCacheAddress, "dst_address", "198.51.100.10"),
							resource.TestCheckResourceAttr(testIpProxyCacheAddress, "dst_host", "cache.test-ip-proxy-cache.example.com"),
							resource.TestCheckResourceAttr(testIpProxyCacheAddress, "dst_port", "8080"),
							resource.TestCheckResourceAttr(testIpProxyCacheAddress, "method", "GET"),
							resource.TestCheckResourceAttr(testIpProxyCacheAddress, "path", "/test_ip_proxy_cache/*"),
							resource.TestCheckResourceAttr(testIpProxyCacheAddress, "src_address", "192.0.2.0/24"),
						),
					},
					{
						Config: testAccIpProxyCacheUpdatedConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpProxyCacheAddress),
							resource.TestCheckResourceAttr(testIpProxyCacheAddress, "action", "allow"),
							resource.TestCheckResourceAttr(testIpProxyCacheAddress, "comment", "test_ip_proxy_cache updated"),
							resource.TestCheckResourceAttr(testIpProxyCacheAddress, "method", "POST"),
							resource.TestCheckResourceAttr(testIpProxyCacheAddress, "path", "/test_ip_proxy_cache/updated/*"),
						),
					},
				},
			})

		})
	}
}

func testAccIpProxyCacheConfig() string {
	return providerConfig + `

resource "routeros_ip_proxy_cache" "test_ip_proxy_cache" {
	action      = "deny"
	comment     = "test_ip_proxy_cache"
	disabled    = true
	dst_address = "198.51.100.10"
	dst_host    = "cache.test-ip-proxy-cache.example.com"
	dst_port    = "8080"
	method      = "GET"
	path        = "/test_ip_proxy_cache/*"
	src_address = "192.0.2.0/24"
}

`
}

func testAccIpProxyCacheUpdatedConfig() string {
	return providerConfig + `

resource "routeros_ip_proxy_cache" "test_ip_proxy_cache" {
	action      = "allow"
	comment     = "test_ip_proxy_cache updated"
	disabled    = true
	dst_address = "198.51.100.10"
	dst_host    = "cache.test-ip-proxy-cache.example.com"
	dst_port    = "8080"
	method      = "POST"
	path        = "/test_ip_proxy_cache/updated/*"
	src_address = "192.0.2.0/24"
}

`
}
