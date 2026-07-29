package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIPv6DhcpRelayOptionAddress = "routeros_ipv6_dhcp_relay_option.test_ipv6_dhcp_relay_option"

func TestAccIPv6DhcpRelayOptionTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ipv6/dhcp-relay/option", "routeros_ipv6_dhcp_relay_option"),
				Steps: []resource.TestStep{
					{
						Config: testAccIPv6DhcpRelayOptionConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIPv6DhcpRelayOptionAddress),
							resource.TestCheckResourceAttr(testIPv6DhcpRelayOptionAddress, "name", "test_ipv6_dhcp_relay_option"),
							resource.TestCheckResourceAttr(testIPv6DhcpRelayOptionAddress, "code", "17"),
							resource.TestCheckResourceAttr(testIPv6DhcpRelayOptionAddress, "value", "0x00000E04"),
							resource.TestCheckResourceAttr(testIPv6DhcpRelayOptionAddress, "raw_value", "00000e04"),
							resource.TestCheckResourceAttr(testIPv6DhcpRelayOptionAddress, "only_if_mac_available", "false"),
							// `default` is reported only for the built-in `client_mac`; absent for user options, so it's not asserted.
							resource.TestCheckNoResourceAttr(testIPv6DhcpRelayOptionAddress, "default"),
						),
					},
					{
						Config: testAccIPv6DhcpRelayOptionUpdatedConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIPv6DhcpRelayOptionAddress),
							resource.TestCheckResourceAttr(testIPv6DhcpRelayOptionAddress, "code", "37"),
							resource.TestCheckResourceAttr(testIPv6DhcpRelayOptionAddress, "value", "0x0001"),
							resource.TestCheckResourceAttr(testIPv6DhcpRelayOptionAddress, "only_if_mac_available", "true"),
							// `raw_value` isn't asserted here: the API `/set` returns a bare `!done`, so device-recomputed values reach state only on the next re-read.
						),
					},
					{
						// No config change: the refresh from the device makes the recomputed `raw_value` visible on both transports.
						Config: testAccIPv6DhcpRelayOptionUpdatedConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIPv6DhcpRelayOptionAddress),
							resource.TestCheckResourceAttr(testIPv6DhcpRelayOptionAddress, "raw_value", "0001"),
						),
					},
				},
			})

		})
	}
}

func testAccIPv6DhcpRelayOptionConfig() string {
	return providerConfig + `

resource "routeros_ipv6_dhcp_relay_option" "test_ipv6_dhcp_relay_option" {
	name                  = "test_ipv6_dhcp_relay_option"
	code                  = 17
	value                 = "0x00000E04"
	only_if_mac_available = false
}

`
}

func testAccIPv6DhcpRelayOptionUpdatedConfig() string {
	return providerConfig + `

resource "routeros_ipv6_dhcp_relay_option" "test_ipv6_dhcp_relay_option" {
	name                  = "test_ipv6_dhcp_relay_option"
	code                  = 37
	value                 = "0x0001"
	only_if_mac_available = true
}

`
}
