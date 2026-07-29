package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testInterfaceBridgeHost = "routeros_interface_bridge_host.test_bridge_host_x"

func TestAccInterfaceBridgeHostTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy: testCheckResourceDestroy("/interface/bridge/host",
					"routeros_interface_bridge_host"),
				Steps: []resource.TestStep{
					{
						Config: testAccInterfaceBridgeHostConfig("02:00:00:B4:11:01", "test_bridge_host_x"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfaceBridgeHost),
							resource.TestCheckResourceAttr(testInterfaceBridgeHost, "bridge", "test_bridge_host_x_br"),
							resource.TestCheckResourceAttr(testInterfaceBridgeHost, "interface", "test_bridge_host_x_veth"),
							resource.TestCheckResourceAttr(testInterfaceBridgeHost, "mac_address", "02:00:00:B4:11:01"),
							resource.TestCheckResourceAttr(testInterfaceBridgeHost, "comment", "test_bridge_host_x"),
							resource.TestCheckResourceAttr(testInterfaceBridgeHost, "disabled", "false"),
							resource.TestCheckResourceAttr(testInterfaceBridgeHost, "dynamic", "false"),
						),
					},
					{
						Config: testAccInterfaceBridgeHostConfig("02:00:00:B4:11:02", "test_bridge_host_x updated"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfaceBridgeHost),
							resource.TestCheckResourceAttr(testInterfaceBridgeHost, "mac_address", "02:00:00:B4:11:02"),
							resource.TestCheckResourceAttr(testInterfaceBridgeHost, "comment",
								"test_bridge_host_x updated"),
						),
					},
				},
			})

		})
	}
}

func testAccInterfaceBridgeHostConfig(mac, comment string) string {
	return providerConfig + fmt.Sprintf(`
resource "routeros_interface_veth" "test_bridge_host_x_veth" {
	name    = "test_bridge_host_x_veth"
	address = ["192.0.2.2/30"]
	gateway = "192.0.2.1"
}

resource "routeros_interface_bridge" "test_bridge_host_x_br" {
	name = "test_bridge_host_x_br"
}

resource "routeros_interface_bridge_port" "test_bridge_host_x_port" {
	bridge    = routeros_interface_bridge.test_bridge_host_x_br.name
	interface = routeros_interface_veth.test_bridge_host_x_veth.name
}

resource "routeros_interface_bridge_host" "test_bridge_host_x" {
	bridge      = routeros_interface_bridge_port.test_bridge_host_x_port.bridge
	interface   = routeros_interface_bridge_port.test_bridge_host_x_port.interface
	mac_address = "%v"
	comment     = "%v"
}
`, mac, comment)
}
