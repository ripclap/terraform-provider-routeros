package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIpSocksAccessAddress = "routeros_ip_socks_access.test_ip_socks_access"

func TestAccIpSocksAccessTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ip/socks/access", "routeros_ip_socks_access"),
				Steps: []resource.TestStep{
					{
						Config: testAccIpSocksAccessConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpSocksAccessAddress),
							resource.TestCheckResourceAttr(testIpSocksAccessAddress, "action", "deny"),
							resource.TestCheckResourceAttr(testIpSocksAccessAddress, "comment", "test_ip_socks_access"),
							resource.TestCheckResourceAttr(testIpSocksAccessAddress, "disabled", "true"),
							resource.TestCheckResourceAttr(testIpSocksAccessAddress, "dst_address", "198.51.100.30"),
							resource.TestCheckResourceAttr(testIpSocksAccessAddress, "dst_port", "8080"),
							resource.TestCheckResourceAttr(testIpSocksAccessAddress, "src_address", "192.0.2.0/24"),
							resource.TestCheckResourceAttr(testIpSocksAccessAddress, "src_port", "1024-65535"),
						),
					},
					{
						Config: testAccIpSocksAccessUpdatedConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpSocksAccessAddress),
							resource.TestCheckResourceAttr(testIpSocksAccessAddress, "action", "allow"),
							resource.TestCheckResourceAttr(testIpSocksAccessAddress, "comment", "test_ip_socks_access updated"),
							resource.TestCheckResourceAttr(testIpSocksAccessAddress, "dst_address", "198.51.100.31"),
							resource.TestCheckResourceAttr(testIpSocksAccessAddress, "dst_port", "8443"),
						),
					},
				},
			})

		})
	}
}

func testAccIpSocksAccessConfig() string {
	return providerConfig + `

resource "routeros_ip_socks_access" "test_ip_socks_access" {
	action      = "deny"
	comment     = "test_ip_socks_access"
	disabled    = true
	dst_address = "198.51.100.30"
	dst_port    = "8080"
	src_address = "192.0.2.0/24"
	src_port    = "1024-65535"
}

`
}

func testAccIpSocksAccessUpdatedConfig() string {
	return providerConfig + `

resource "routeros_ip_socks_access" "test_ip_socks_access" {
	action      = "allow"
	comment     = "test_ip_socks_access updated"
	disabled    = true
	dst_address = "198.51.100.31"
	dst_port    = "8443"
	src_address = "192.0.2.0/24"
	src_port    = "1024-65535"
}

`
}
