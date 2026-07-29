package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testInterfaceMeshPortAddress = "routeros_interface_mesh_port.test_mesh_port_x"

func TestAccInterfaceMeshPortTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/interface/mesh/port", "routeros_interface_mesh_port"),
				Steps: []resource.TestStep{
					{
						Config: testAccInterfaceMeshPortConfig("15s", 25, "ethernet"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfaceMeshPortAddress),
							resource.TestCheckResourceAttr(testInterfaceMeshPortAddress, "mesh", "test_mesh_port_x_mesh"),
							resource.TestCheckResourceAttr(testInterfaceMeshPortAddress, "interface", "test_mesh_port_x_veth"),
							resource.TestCheckResourceAttr(testInterfaceMeshPortAddress, "comment", "test_mesh_port_x"),
							resource.TestCheckResourceAttr(testInterfaceMeshPortAddress, "hello_interval", "15s"),
							resource.TestCheckResourceAttr(testInterfaceMeshPortAddress, "path_cost", "25"),
							resource.TestCheckResourceAttr(testInterfaceMeshPortAddress, "port_type", "ethernet"),
							resource.TestCheckResourceAttr(testInterfaceMeshPortAddress, "disabled", "false"),
							resource.TestCheckResourceAttr(testInterfaceMeshPortAddress, "dynamic", "false"),
						),
					},
					{
						Config: testAccInterfaceMeshPortConfig("30s", 10, "auto"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfaceMeshPortAddress),
							resource.TestCheckResourceAttr(testInterfaceMeshPortAddress, "hello_interval", "30s"),
							resource.TestCheckResourceAttr(testInterfaceMeshPortAddress, "path_cost", "10"),
							resource.TestCheckResourceAttr(testInterfaceMeshPortAddress, "port_type", "auto"),
						),
					},
				},
			})

		})
	}
}

func testAccInterfaceMeshPortConfig(helloInterval string, pathCost int, portType string) string {
	return providerConfig + fmt.Sprintf(`
resource "routeros_interface_veth" "test_mesh_port_x_veth" {
	name = "test_mesh_port_x_veth"
}

resource "routeros_interface_mesh" "test_mesh_port_x_mesh" {
	name = "test_mesh_port_x_mesh"
}

resource "routeros_interface_mesh_port" "test_mesh_port_x" {
	mesh           = routeros_interface_mesh.test_mesh_port_x_mesh.name
	interface      = routeros_interface_veth.test_mesh_port_x_veth.name
	comment        = "test_mesh_port_x"
	hello_interval = "%v"
	path_cost      = %v
	port_type      = "%v"
	disabled       = false
}
`, helloInterval, pathCost, portType)
}
