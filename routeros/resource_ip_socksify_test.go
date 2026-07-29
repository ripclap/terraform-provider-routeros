package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIpSocksifyAddress = "routeros_ip_socksify.test_ip_socksify"

func TestAccIpSocksifyTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ip/socksify", "routeros_ip_socksify"),
				Steps: []resource.TestStep{
					{
						Config: testAccIpSocksifyConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpSocksifyAddress),
							resource.TestCheckResourceAttr(testIpSocksifyAddress, "name", "test_ip_socksify"),
							resource.TestCheckResourceAttr(testIpSocksifyAddress, "comment", "test_ip_socksify"),
							resource.TestCheckResourceAttr(testIpSocksifyAddress, "connection_timeout", "30"),
							resource.TestCheckResourceAttr(testIpSocksifyAddress, "disabled", "true"),
							resource.TestCheckResourceAttr(testIpSocksifyAddress, "port", "21081"),
							resource.TestCheckResourceAttr(testIpSocksifyAddress, "socks5_password", "TestIpSocksify1"),
							resource.TestCheckResourceAttr(testIpSocksifyAddress, "socks5_port", "1080"),
							resource.TestCheckResourceAttr(testIpSocksifyAddress, "socks5_server", "192.0.2.71"),
							resource.TestCheckResourceAttr(testIpSocksifyAddress, "socks5_user", "test_ip_socksify_u1"),
						),
					},
					{
						Config: testAccIpSocksifyUpdatedConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpSocksifyAddress),
							resource.TestCheckResourceAttr(testIpSocksifyAddress, "comment", "test_ip_socksify updated"),
							resource.TestCheckResourceAttr(testIpSocksifyAddress, "connection_timeout", "60"),
							resource.TestCheckResourceAttr(testIpSocksifyAddress, "port", "21082"),
							resource.TestCheckResourceAttr(testIpSocksifyAddress, "socks5_password", "TestIpSocksify2"),
							resource.TestCheckResourceAttr(testIpSocksifyAddress, "socks5_port", "11080"),
							resource.TestCheckResourceAttr(testIpSocksifyAddress, "socks5_server", "192.0.2.72"),
							resource.TestCheckResourceAttr(testIpSocksifyAddress, "socks5_user", "test_ip_socksify_u2"),
						),
					},
				},
			})

		})
	}
}

func testAccIpSocksifyConfig() string {
	return providerConfig + `

resource "routeros_ip_socksify" "test_ip_socksify" {
	name               = "test_ip_socksify"
	comment            = "test_ip_socksify"
	connection_timeout = 30
	disabled           = true
	port               = 21081
	socks5_password    = "TestIpSocksify1"
	socks5_port        = 1080
	socks5_server      = "192.0.2.71"
	socks5_user        = "test_ip_socksify_u1"
}

`
}

func testAccIpSocksifyUpdatedConfig() string {
	return providerConfig + `

resource "routeros_ip_socksify" "test_ip_socksify" {
	name               = "test_ip_socksify"
	comment            = "test_ip_socksify updated"
	connection_timeout = 60
	disabled           = true
	port               = 21082
	socks5_password    = "TestIpSocksify2"
	socks5_port        = 11080
	socks5_server      = "192.0.2.72"
	socks5_user        = "test_ip_socksify_u2"
}

`
}
