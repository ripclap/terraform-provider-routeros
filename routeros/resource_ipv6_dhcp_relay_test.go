package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIPv6DhcpRelayAddress = "routeros_ipv6_dhcp_relay.test_ipv6_dhcp_relay"

func TestAccIPv6DhcpRelayTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ipv6/dhcp-relay", "routeros_ipv6_dhcp_relay"),
				Steps: []resource.TestStep{
					{
						Config: testAccIPv6DhcpRelayConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIPv6DhcpRelayAddress),
							resource.TestCheckResourceAttr(testIPv6DhcpRelayAddress, "name", "test_ipv6_dhcp_relay"),
							resource.TestCheckResourceAttr(testIPv6DhcpRelayAddress, "comment", "test_ipv6_dhcp_relay"),
							resource.TestCheckResourceAttr(testIPv6DhcpRelayAddress, "delay_threshold", "5s"),
							resource.TestCheckResourceAttr(testIPv6DhcpRelayAddress, "disabled", "true"),
							resource.TestCheckResourceAttr(testIPv6DhcpRelayAddress, "interface", testIPv6DhcpRelayInterface),
							resource.TestCheckResourceAttr(testIPv6DhcpRelayAddress, "link_address", "2001:db8:1::1"),
							resource.TestCheckResourceAttr(testIPv6DhcpRelayAddress, "store_relayed_bindings", "true"),
							resource.TestCheckResourceAttr(testIPv6DhcpRelayAddress, "dhcp_server.#", "2"),
							resource.TestCheckTypeSetElemAttr(testIPv6DhcpRelayAddress, "dhcp_server.*", "2001:db8::1"),
							resource.TestCheckTypeSetElemAttr(testIPv6DhcpRelayAddress, "dhcp_server.*", "2001:db8::2"),
							resource.TestCheckResourceAttr(testIPv6DhcpRelayAddress, "dhcp_options.#", "1"),
							resource.TestCheckTypeSetElemAttr(testIPv6DhcpRelayAddress, "dhcp_options.*", "client_mac"),
						),
					},
					{
						Config: testAccIPv6DhcpRelayUpdatedConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIPv6DhcpRelayAddress),
							resource.TestCheckResourceAttr(testIPv6DhcpRelayAddress, "comment", "test_ipv6_dhcp_relay updated"),
							resource.TestCheckResourceAttr(testIPv6DhcpRelayAddress, "delay_threshold", "10s"),
							resource.TestCheckResourceAttr(testIPv6DhcpRelayAddress, "link_address", "2001:db8:2::1"),
							resource.TestCheckResourceAttr(testIPv6DhcpRelayAddress, "store_relayed_bindings", "false"),
							resource.TestCheckResourceAttr(testIPv6DhcpRelayAddress, "dhcp_server.#", "1"),
							resource.TestCheckTypeSetElemAttr(testIPv6DhcpRelayAddress, "dhcp_server.*", "2001:db8::3"),
						),
					},
				},
			})

		})
	}
}

// Reuse an existing interface instead of creating one: an extra throwaway object can trip up the concurrent acceptance tests.
const testIPv6DhcpRelayInterface = "ether6"

func testAccIPv6DhcpRelayConfig() string {
	return providerConfig + `

resource "routeros_ipv6_dhcp_relay" "test_ipv6_dhcp_relay" {
	name                   = "test_ipv6_dhcp_relay"
	comment                = "test_ipv6_dhcp_relay"
	delay_threshold        = "5s"
	dhcp_options           = ["client_mac"]
	dhcp_server            = ["2001:db8::1", "2001:db8::2"]
	disabled               = true
	interface              = "` + testIPv6DhcpRelayInterface + `"
	link_address           = "2001:db8:1::1"
	store_relayed_bindings = true
}

`
}

func testAccIPv6DhcpRelayUpdatedConfig() string {
	return providerConfig + `

resource "routeros_ipv6_dhcp_relay" "test_ipv6_dhcp_relay" {
	name                   = "test_ipv6_dhcp_relay"
	comment                = "test_ipv6_dhcp_relay updated"
	delay_threshold        = "10s"
	dhcp_options           = ["client_mac"]
	dhcp_server            = ["2001:db8::3"]
	disabled               = true
	interface              = "` + testIPv6DhcpRelayInterface + `"
	link_address           = "2001:db8:2::1"
	store_relayed_bindings = false
}

`
}
