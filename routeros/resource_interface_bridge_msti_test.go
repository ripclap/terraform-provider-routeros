package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testInterfaceBridgeMsti = "routeros_interface_bridge_msti.test_bridge_msti_x"

func TestAccInterfaceBridgeMstiTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy: testCheckResourceDestroy("/interface/bridge/msti",
					"routeros_interface_bridge_msti"),
				Steps: []resource.TestStep{
					{
						Config: testAccInterfaceBridgeMstiConfig("0x7000", "1500-1510,1520", "test_bridge_msti_x"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfaceBridgeMsti),
							resource.TestCheckResourceAttr(testInterfaceBridgeMsti, "bridge", "test_bridge_msti_x_br"),
							resource.TestCheckResourceAttr(testInterfaceBridgeMsti, "identifier", "9"),
							resource.TestCheckResourceAttr(testInterfaceBridgeMsti, "priority", "0x7000"),
							resource.TestCheckResourceAttr(testInterfaceBridgeMsti, "vlan_mapping", "1500-1510,1520"),
							resource.TestCheckResourceAttr(testInterfaceBridgeMsti, "comment", "test_bridge_msti_x"),
							resource.TestCheckResourceAttr(testInterfaceBridgeMsti, "disabled", "false"),
							resource.TestCheckResourceAttr(testInterfaceBridgeMsti, "dynamic", "false"),
						),
					},
					{
						Config: testAccInterfaceBridgeMstiConfig("0x9000", "1530-1540", "test_bridge_msti_x updated"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfaceBridgeMsti),
							resource.TestCheckResourceAttr(testInterfaceBridgeMsti, "priority", "0x9000"),
							resource.TestCheckResourceAttr(testInterfaceBridgeMsti, "vlan_mapping", "1530-1540"),
							resource.TestCheckResourceAttr(testInterfaceBridgeMsti, "comment",
								"test_bridge_msti_x updated"),
						),
					},
				},
			})

		})
	}
}

func testAccInterfaceBridgeMstiConfig(priority, vlanMapping, comment string) string {
	return providerConfig + fmt.Sprintf(`
resource "routeros_interface_bridge" "test_bridge_msti_x_br" {
	name           = "test_bridge_msti_x_br"
	vlan_filtering = true
	protocol_mode  = "mstp"
}

resource "routeros_interface_bridge_msti" "test_bridge_msti_x" {
	bridge       = routeros_interface_bridge.test_bridge_msti_x_br.name
	identifier   = 9
	priority     = "%v"
	vlan_mapping = "%v"
	comment      = "%v"
}
`, priority, vlanMapping, comment)
}
