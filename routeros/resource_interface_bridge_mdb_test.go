package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testInterfaceBridgeMdb = "routeros_interface_bridge_mdb.test_bridge_mdb_x"

func TestAccInterfaceBridgeMdbTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy: testCheckResourceDestroy("/interface/bridge/mdb",
					"routeros_interface_bridge_mdb"),
				Steps: []resource.TestStep{
					{
						Config: testAccInterfaceBridgeMdbConfig("239.10.11.12", "test_bridge_mdb_x"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfaceBridgeMdb),
							resource.TestCheckResourceAttr(testInterfaceBridgeMdb, "bridge", "test_bridge_mdb_x_br"),
							resource.TestCheckResourceAttr(testInterfaceBridgeMdb, "group", "239.10.11.12"),
							resource.TestCheckResourceAttr(testInterfaceBridgeMdb, "interface.#", "1"),
							resource.TestCheckResourceAttr(testInterfaceBridgeMdb, "interface.0",
								"test_bridge_mdb_x_veth"),
							resource.TestCheckResourceAttr(testInterfaceBridgeMdb, "vid", "1"),
							resource.TestCheckResourceAttr(testInterfaceBridgeMdb, "comment", "test_bridge_mdb_x"),
							resource.TestCheckResourceAttr(testInterfaceBridgeMdb, "disabled", "false"),
							resource.TestCheckResourceAttr(testInterfaceBridgeMdb, "dynamic", "false"),
						),
					},
					{
						Config: testAccInterfaceBridgeMdbConfig("ff02::1:3", "test_bridge_mdb_x updated"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfaceBridgeMdb),
							resource.TestCheckResourceAttr(testInterfaceBridgeMdb, "group", "ff02::1:3"),
							resource.TestCheckResourceAttr(testInterfaceBridgeMdb, "comment",
								"test_bridge_mdb_x updated"),
						),
					},
				},
			})

		})
	}
}

func testAccInterfaceBridgeMdbConfig(group, comment string) string {
	return providerConfig + fmt.Sprintf(`
resource "routeros_interface_veth" "test_bridge_mdb_x_veth" {
	name    = "test_bridge_mdb_x_veth"
	address = ["192.0.2.6/30"]
	gateway = "192.0.2.5"
}

resource "routeros_interface_bridge" "test_bridge_mdb_x_br" {
	name           = "test_bridge_mdb_x_br"
	vlan_filtering = true
	igmp_snooping  = true
}

resource "routeros_interface_bridge_port" "test_bridge_mdb_x_port" {
	bridge    = routeros_interface_bridge.test_bridge_mdb_x_br.name
	interface = routeros_interface_veth.test_bridge_mdb_x_veth.name
}

resource "routeros_interface_bridge_mdb" "test_bridge_mdb_x" {
	bridge    = routeros_interface_bridge_port.test_bridge_mdb_x_port.bridge
	group     = "%v"
	interface = [routeros_interface_bridge_port.test_bridge_mdb_x_port.interface]
	vid       = 1
	comment   = "%v"
}
`, group, comment)
}
