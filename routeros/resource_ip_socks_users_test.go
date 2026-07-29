package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIpSocksUsersAddress = "routeros_ip_socks_users.test_ip_socks_users"

func TestAccIpSocksUsersTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ip/socks/users", "routeros_ip_socks_users"),
				Steps: []resource.TestStep{
					{
						Config: testAccIpSocksUsersConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpSocksUsersAddress),
							resource.TestCheckResourceAttr(testIpSocksUsersAddress, "name", "test_ip_socks_users"),
							resource.TestCheckResourceAttr(testIpSocksUsersAddress, "disabled", "true"),
							resource.TestCheckResourceAttr(testIpSocksUsersAddress, "only_one", "true"),
							resource.TestCheckResourceAttr(testIpSocksUsersAddress, "password", "TestIpSocksUsers1"),
							resource.TestCheckResourceAttr(testIpSocksUsersAddress, "rate_limit", "1M/1M"),
						),
					},
					{
						Config: testAccIpSocksUsersUpdatedConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpSocksUsersAddress),
							resource.TestCheckResourceAttr(testIpSocksUsersAddress, "only_one", "false"),
							resource.TestCheckResourceAttr(testIpSocksUsersAddress, "password", "TestIpSocksUsers2"),
							resource.TestCheckResourceAttr(testIpSocksUsersAddress, "rate_limit", "64k/64k"),
						),
					},
				},
			})

		})
	}
}

func testAccIpSocksUsersConfig() string {
	return providerConfig + `

resource "routeros_ip_socks_users" "test_ip_socks_users" {
	name       = "test_ip_socks_users"
	disabled   = true
	only_one   = true
	password   = "TestIpSocksUsers1"
	rate_limit = "1M/1M"
}

`
}

func testAccIpSocksUsersUpdatedConfig() string {
	return providerConfig + `

resource "routeros_ip_socks_users" "test_ip_socks_users" {
	name       = "test_ip_socks_users"
	disabled   = true
	only_one   = false
	password   = "TestIpSocksUsers2"
	rate_limit = "64k/64k"
}

`
}
