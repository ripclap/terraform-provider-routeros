package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testInterfaceVpls = "routeros_interface_vpls.test_vpls_x"

func TestAccInterfaceVplsTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/interface/vpls", "routeros_interface_vpls"),
				Steps: []resource.TestStep{
					{
						Config: testAccInterfaceVplsConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfaceVpls),
							resource.TestCheckResourceAttr(testInterfaceVpls, "name", "test_vpls_x"),
							resource.TestCheckResourceAttr(testInterfaceVpls, "peer", "192.0.2.77"),
							resource.TestCheckResourceAttr(testInterfaceVpls, "vpls_id", "65099:77"),
							resource.TestCheckResourceAttr(testInterfaceVpls, "comment", "test_vpls_x"),
							resource.TestCheckResourceAttr(testInterfaceVpls, "disabled", "true"),
							resource.TestCheckResourceAttr(testInterfaceVpls, "mac_address", "02:CA:FE:00:00:77"),
							resource.TestCheckResourceAttr(testInterfaceVpls, "running", "false"),
						),
					},
					{
						Config: testAccInterfaceVplsUpdatedConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfaceVpls),
							resource.TestCheckResourceAttr(testInterfaceVpls, "arp", "proxy-arp"),
							resource.TestCheckResourceAttr(testInterfaceVpls, "arp_timeout", "30s"),
							resource.TestCheckResourceAttr(testInterfaceVpls, "bridge", "test_vpls_br_x"),
							resource.TestCheckResourceAttr(testInterfaceVpls, "bridge_cost", "100"),
							resource.TestCheckResourceAttr(testInterfaceVpls, "bridge_horizon", "none"),
							resource.TestCheckResourceAttr(testInterfaceVpls, "bridge_pvid", "1"),
							resource.TestCheckResourceAttr(testInterfaceVpls, "disable_running_check", "true"),
							resource.TestCheckResourceAttr(testInterfaceVpls, "mtu", "1500"),
							resource.TestCheckResourceAttr(testInterfaceVpls, "pw_control_word", "enabled"),
							resource.TestCheckResourceAttr(testInterfaceVpls, "pw_l2mtu", "1520"),
							resource.TestCheckResourceAttr(testInterfaceVpls, "pw_type", "tagged-ethernet"),
							resource.TestCheckResourceAttr(testInterfaceVpls, "comment", "updated"),
						),
					},
				},
			})

		})
	}
}

func testAccInterfaceVplsConfig() string {
	return providerConfig + `

resource "routeros_interface_bridge" "test_vpls_br_x" {
	name = "test_vpls_br_x"
}

resource "routeros_interface_vpls" "test_vpls_x" {
	name        = "test_vpls_x"
	peer        = "192.0.2.77"
	vpls_id     = "65099:77"
	mac_address = "02:CA:FE:00:00:77"
	comment     = "test_vpls_x"
	disabled    = true
}
`
}

func testAccInterfaceVplsUpdatedConfig() string {
	return providerConfig + `

resource "routeros_interface_bridge" "test_vpls_br_x" {
	name = "test_vpls_br_x"
}

resource "routeros_interface_vpls" "test_vpls_x" {
	name                  = "test_vpls_x"
	peer                  = "192.0.2.77"
	vpls_id               = "65099:77"
	mac_address           = "02:CA:FE:00:00:77"
	comment               = "updated"
	disabled              = true
	arp                   = "proxy-arp"
	arp_timeout           = "30s"
	bridge                = routeros_interface_bridge.test_vpls_br_x.name
	bridge_cost           = 100
	bridge_horizon        = "none"
	bridge_pvid           = 1
	disable_running_check = true
	mtu                   = 1500
	pw_control_word       = "enabled"
	pw_l2mtu              = 1520
	pw_type               = "tagged-ethernet"
}
`
}
