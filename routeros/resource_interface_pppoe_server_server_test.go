package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testInterfacePppoeServerServerAddress = "routeros_interface_pppoe_server_server.test_pppoe_server_server_x"

func TestAccInterfacePppoeServerServerTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy: testCheckResourceDestroy("/interface/pppoe-server/server",
					"routeros_interface_pppoe_server_server"),
				Steps: []resource.TestStep{
					{
						Config: testAccInterfacePppoeServerServerConfig("test_pppoe_server_server_x", 100),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfacePppoeServerServerAddress),
							resource.TestCheckResourceAttr(testInterfacePppoeServerServerAddress, "service_name",
								"test_pppoe_server_server_x"),
							resource.TestCheckResourceAttr(testInterfacePppoeServerServerAddress, "interface",
								"test_pppoe_server_server_x_veth"),
							resource.TestCheckResourceAttr(testInterfacePppoeServerServerAddress, "comment",
								"test_pppoe_server_server_x"),
							resource.TestCheckResourceAttr(testInterfacePppoeServerServerAddress, "pado_delay", "100"),
							resource.TestCheckResourceAttr(testInterfacePppoeServerServerAddress, "disabled", "true"),
							resource.TestCheckResourceAttr(testInterfacePppoeServerServerAddress, "one_session_per_host", "true"),
							resource.TestCheckResourceAttr(testInterfacePppoeServerServerAddress, "accept_empty_service", "false"),
							resource.TestCheckResourceAttr(testInterfacePppoeServerServerAddress, "accept_untagged", "false"),
							resource.TestCheckResourceAttr(testInterfacePppoeServerServerAddress, "pppoe_over_vlan_range", "100-115,120"),
							resource.TestCheckResourceAttr(testInterfacePppoeServerServerAddress, "default_profile", "default"),
							resource.TestCheckResourceAttr(testInterfacePppoeServerServerAddress, "invalid", "false"),
							resource.TestCheckResourceAttr(testInterfacePppoeServerServerAddress, "authentication.#", "2"),
						),
					},
					{
						Config: testAccInterfacePppoeServerServerConfig("test_pppoe_server_server_x2", 250),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfacePppoeServerServerAddress),
							resource.TestCheckResourceAttr(testInterfacePppoeServerServerAddress, "service_name",
								"test_pppoe_server_server_x2"),
							resource.TestCheckResourceAttr(testInterfacePppoeServerServerAddress, "pado_delay", "250"),
						),
					},
				},
			})

		})
	}
}

func testAccInterfacePppoeServerServerConfig(serviceName string, padoDelay int) string {
	return providerConfig + fmt.Sprintf(`
resource "routeros_interface_veth" "test_pppoe_server_server_x_veth" {
	name = "test_pppoe_server_server_x_veth"
}

resource "routeros_interface_pppoe_server_server" "test_pppoe_server_server_x" {
	interface             = routeros_interface_veth.test_pppoe_server_server_x_veth.name
	service_name          = "%v"
	comment               = "test_pppoe_server_server_x"
	disabled              = true
	accept_empty_service  = false
	accept_untagged       = false
	authentication        = ["mschap1", "mschap2"]
	default_profile       = "default"
	keepalive_timeout     = "10"
	max_mru               = "1480"
	max_mtu               = "1480"
	max_sessions          = "unlimited"
	mrru                  = "disabled"
	one_session_per_host  = true
	pado_delay            = %v
	pppoe_over_vlan_range = "100-115,120"
}
`, serviceName, padoDelay)
}
