package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIpProxyAccess = "routeros_ip_proxy_access.test_proxy_access_x"

func TestAccIpProxyAccessTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ip/proxy/access", "routeros_ip_proxy_access"),
				Steps: []resource.TestStep{
					{
						Config: testAccIpProxyAccessConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpProxyAccess),
							resource.TestCheckResourceAttr(testIpProxyAccess, "action", "deny"),
							resource.TestCheckResourceAttr(testIpProxyAccess, "dst_host", "test-proxy-access-x.example.com"),
							resource.TestCheckResourceAttr(testIpProxyAccess, "dst_port", "80"),
							resource.TestCheckResourceAttr(testIpProxyAccess, "method", "GET"),
							resource.TestCheckResourceAttr(testIpProxyAccess, "path", "/test_proxy_access_x/*"),
							resource.TestCheckResourceAttr(testIpProxyAccess, "src_address", "192.0.2.70"),
							resource.TestCheckResourceAttr(testIpProxyAccess, "comment", "test_proxy_access_x"),
							resource.TestCheckResourceAttr(testIpProxyAccess, "disabled", "true"),
						),
					},
					{
						Config: testAccIpProxyAccessUpdatedConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpProxyAccess),
							resource.TestCheckResourceAttr(testIpProxyAccess, "action", "redirect"),
							resource.TestCheckResourceAttr(testIpProxyAccess, "action_data",
								"test-proxy-access-x.example.com/blocked"),
							resource.TestCheckResourceAttr(testIpProxyAccess, "dst_port", "8080"),
							resource.TestCheckResourceAttr(testIpProxyAccess, "method", "POST"),
							resource.TestCheckResourceAttr(testIpProxyAccess, "comment", "test_proxy_access_x_upd"),
						),
					},
				},
			})

		})
	}
}

func testAccIpProxyAccessConfig() string {
	return providerConfig + `
resource "routeros_ip_proxy_access" "test_proxy_access_x" {
	action      = "deny"
	dst_host    = "test-proxy-access-x.example.com"
	dst_port    = "80"
	method      = "GET"
	path        = "/test_proxy_access_x/*"
	src_address = "192.0.2.70"
	comment     = "test_proxy_access_x"
	disabled    = true
}
`
}

func testAccIpProxyAccessUpdatedConfig() string {
	return providerConfig + `
resource "routeros_ip_proxy_access" "test_proxy_access_x" {
	action      = "redirect"
	action_data = "test-proxy-access-x.example.com/blocked"
	dst_host    = "test-proxy-access-x.example.com"
	dst_port    = "8080"
	method      = "POST"
	path        = "/test_proxy_access_x/*"
	src_address = "192.0.2.70"
	comment     = "test_proxy_access_x_upd"
	disabled    = true
}
`
}
