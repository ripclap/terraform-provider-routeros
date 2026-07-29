package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testInterfaceBridgePortMstOverride = "routeros_interface_bridge_port_mst_override.test_mst_override_x"

func TestAccInterfaceBridgePortMstOverrideTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy: testCheckResourceDestroy("/interface/bridge/port/mst-override",
					"routeros_interface_bridge_port_mst_override"),
				Steps: []resource.TestStep{
					{
						Config: testAccInterfaceBridgePortMstOverrideConfig("1000", "0x80"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfaceBridgePortMstOverride),
							resource.TestCheckResourceAttr(testInterfaceBridgePortMstOverride, "identifier", "17"),
							resource.TestCheckResourceAttr(testInterfaceBridgePortMstOverride, "interface",
								"test_mst_override_x_vlan"),
							resource.TestCheckResourceAttr(testInterfaceBridgePortMstOverride, "internal_path_cost", "1000"),
							resource.TestCheckResourceAttr(testInterfaceBridgePortMstOverride, "priority", "0x80"),
							resource.TestCheckResourceAttr(testInterfaceBridgePortMstOverride, "disabled", "false"),
							resource.TestCheckResourceAttr(testInterfaceBridgePortMstOverride, "comment",
								"test_mst_override_x"),
						),
					},
					{
						Config: testAccInterfaceBridgePortMstOverrideConfig("20000", "0x40"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfaceBridgePortMstOverride),
							resource.TestCheckResourceAttr(testInterfaceBridgePortMstOverride, "internal_path_cost", "20000"),
							resource.TestCheckResourceAttr(testInterfaceBridgePortMstOverride, "priority", "0x40"),
						),
					},
				},
			})
		})
	}
}

// The override needs an active MSTP bridge port and a pre-existing MST instance; the veth fixtures are
// down, so a VLAN interface over ether1 provides the active port without touching ether1's addressing.
func testAccInterfaceBridgePortMstOverrideConfig(cost, priority string) string {
	return providerConfig + `

resource "routeros_interface_bridge" "test_mst_override_x_br" {
	name           = "test_mst_override_x_br"
	protocol_mode  = "mstp"
	vlan_filtering = true
}

resource "routeros_interface_vlan" "test_mst_override_x_vlan" {
	name      = "test_mst_override_x_vlan"
	interface = "ether1"
	vlan_id   = 3921
}

resource "routeros_interface_bridge_port" "test_mst_override_x_port" {
	bridge    = routeros_interface_bridge.test_mst_override_x_br.name
	interface = routeros_interface_vlan.test_mst_override_x_vlan.name
}

resource "routeros_interface_bridge_msti" "test_mst_override_x_msti" {
	bridge       = routeros_interface_bridge.test_mst_override_x_br.name
	identifier   = 17
	vlan_mapping = "3921"
}

resource "routeros_interface_bridge_port_mst_override" "test_mst_override_x" {
	identifier         = routeros_interface_bridge_msti.test_mst_override_x_msti.identifier
	interface          = routeros_interface_bridge_port.test_mst_override_x_port.interface
	internal_path_cost = ` + cost + `
	priority           = "` + priority + `"
	comment            = "test_mst_override_x"
}
`
}
