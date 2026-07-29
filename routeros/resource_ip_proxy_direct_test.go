package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIpProxyDirectAddress = "routeros_ip_proxy_direct.test_ip_proxy_direct"

func TestAccIpProxyDirectTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ip/proxy/direct", "routeros_ip_proxy_direct"),
				Steps: []resource.TestStep{
					{
						Config: testAccIpProxyDirectConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpProxyDirectAddress),
							resource.TestCheckResourceAttr(testIpProxyDirectAddress, "action", "allow"),
							resource.TestCheckResourceAttr(testIpProxyDirectAddress, "comment", "test_ip_proxy_direct"),
							resource.TestCheckResourceAttr(testIpProxyDirectAddress, "disabled", "true"),
							resource.TestCheckResourceAttr(testIpProxyDirectAddress, "dst_address", "198.51.100.20"),
							resource.TestCheckResourceAttr(testIpProxyDirectAddress, "dst_host", "direct.test-ip-proxy-direct.example.com"),
							resource.TestCheckResourceAttr(testIpProxyDirectAddress, "dst_port", "443"),
							resource.TestCheckResourceAttr(testIpProxyDirectAddress, "method", "CONNECT"),
							resource.TestCheckResourceAttr(testIpProxyDirectAddress, "path", "/test_ip_proxy_direct/*"),
							resource.TestCheckResourceAttr(testIpProxyDirectAddress, "src_address", "192.0.2.0/24"),
						),
					},
					{
						Config: testAccIpProxyDirectUpdatedConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpProxyDirectAddress),
							resource.TestCheckResourceAttr(testIpProxyDirectAddress, "action", "deny"),
							resource.TestCheckResourceAttr(testIpProxyDirectAddress, "comment", "test_ip_proxy_direct updated"),
							resource.TestCheckResourceAttr(testIpProxyDirectAddress, "method", "HEAD"),
							resource.TestCheckResourceAttr(testIpProxyDirectAddress, "dst_port", "8443"),
						),
					},
				},
			})

		})
	}
}

func testAccIpProxyDirectConfig() string {
	return providerConfig + `

resource "routeros_ip_proxy_direct" "test_ip_proxy_direct" {
	action      = "allow"
	comment     = "test_ip_proxy_direct"
	disabled    = true
	dst_address = "198.51.100.20"
	dst_host    = "direct.test-ip-proxy-direct.example.com"
	dst_port    = "443"
	method      = "CONNECT"
	path        = "/test_ip_proxy_direct/*"
	src_address = "192.0.2.0/24"
}

`
}

func testAccIpProxyDirectUpdatedConfig() string {
	return providerConfig + `

resource "routeros_ip_proxy_direct" "test_ip_proxy_direct" {
	action      = "deny"
	comment     = "test_ip_proxy_direct updated"
	disabled    = true
	dst_address = "198.51.100.20"
	dst_host    = "direct.test-ip-proxy-direct.example.com"
	dst_port    = "8443"
	method      = "HEAD"
	path        = "/test_ip_proxy_direct/*"
	src_address = "192.0.2.0/24"
}

`
}
