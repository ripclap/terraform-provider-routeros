package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIpSocksAddress = "routeros_ip_socks.test_ip_socks"

// `/ip/socks` is a settings singleton (no CheckDestroy); the server is never enabled.
func TestAccIpSocksTest_basic(t *testing.T) {
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
						Config: testAccIpSocksConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpSocksAddress),
							resource.TestCheckResourceAttr(testIpSocksAddress, "enabled", "false"),
							resource.TestCheckResourceAttr(testIpSocksAddress, "auth_method", "password"),
							resource.TestCheckResourceAttr(testIpSocksAddress, "connection_idle_timeout", "3m"),
							resource.TestCheckResourceAttr(testIpSocksAddress, "max_connections", "123"),
							resource.TestCheckResourceAttr(testIpSocksAddress, "port", "21080"),
							resource.TestCheckResourceAttr(testIpSocksAddress, "version", "5"),
							resource.TestCheckResourceAttr(testIpSocksAddress, "vrf", "main"),
						),
					},
					{
						Config: testAccIpSocksRestoredConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpSocksAddress),
							resource.TestCheckResourceAttr(testIpSocksAddress, "enabled", "false"),
							resource.TestCheckResourceAttr(testIpSocksAddress, "auth_method", "none"),
							resource.TestCheckResourceAttr(testIpSocksAddress, "connection_idle_timeout", "2m"),
							resource.TestCheckResourceAttr(testIpSocksAddress, "max_connections", "200"),
							resource.TestCheckResourceAttr(testIpSocksAddress, "port", "1080"),
							resource.TestCheckResourceAttr(testIpSocksAddress, "version", "4"),
						),
					},
				},
			})

		})
	}
}

func testAccIpSocksConfig() string {
	return providerConfig + `
resource "routeros_ip_socks" "test_ip_socks" {
	enabled                 = false
	auth_method             = "password"
	connection_idle_timeout = "3m"
	max_connections         = 123
	port                    = 21080
	version                 = "5"
	vrf                     = "main"
}`
}

func testAccIpSocksRestoredConfig() string {
	return providerConfig + `
resource "routeros_ip_socks" "test_ip_socks" {
	enabled                 = false
	auth_method             = "none"
	connection_idle_timeout = "2m"
	max_connections         = 200
	port                    = 1080
	version                 = "4"
	vrf                     = "main"
}`
}
