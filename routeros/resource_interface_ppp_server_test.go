package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testInterfacePppServerAddress = "routeros_interface_ppp_server.test_ppp_server_x"

func TestAccInterfacePppServerTest_basic(t *testing.T) {
	testCheckMenu(t, "/port")
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/interface/ppp-server", "routeros_interface_ppp_server"),
				Steps: []resource.TestStep{
					{
						Config: testAccInterfacePppServerConfig("test_ppp_server_x", 2),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfacePppServerAddress),
							resource.TestCheckResourceAttr(testInterfacePppServerAddress, "name", "test_ppp_server_x"),
							resource.TestCheckResourceAttr(testInterfacePppServerAddress, "comment", "test_ppp_server_x"),
							resource.TestCheckResourceAttr(testInterfacePppServerAddress, "ring_count", "2"),
							resource.TestCheckResourceAttr(testInterfacePppServerAddress, "disabled", "true"),
							resource.TestCheckResourceAttr(testInterfacePppServerAddress, "null_modem", "true"),
							resource.TestCheckResourceAttr(testInterfacePppServerAddress, "profile", "default"),
							resource.TestCheckResourceAttr(testInterfacePppServerAddress, "port", "serial0"),
							resource.TestCheckResourceAttr(testInterfacePppServerAddress, "authentication.#", "2"),
						),
					},
					{
						Config: testAccInterfacePppServerConfig("test_ppp_server_x updated", 4),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfacePppServerAddress),
							resource.TestCheckResourceAttr(testInterfacePppServerAddress, "comment", "test_ppp_server_x updated"),
							resource.TestCheckResourceAttr(testInterfacePppServerAddress, "ring_count", "4"),
						),
					},
				},
			})

		})
	}
}

func testAccInterfacePppServerConfig(comment string, ringCount int) string {
	return providerConfig + fmt.Sprintf(`
resource "routeros_interface_ppp_server" "test_ppp_server_x" {
	name           = "test_ppp_server_x"
	comment        = "%v"
	ring_count     = %v
	disabled       = true
	null_modem     = true
	profile        = "default"
	authentication = ["mschap1", "mschap2"]
	max_mru        = 1500
	max_mtu        = 1500
	mrru           = "disabled"
	modem_init     = ""

	# 'port' has no diff suppression in the schema, and RouterOS always reports the
	# placeholder '*FFFFFFFF' when no serial port is bound, so the port has to be set
	# explicitly to keep the plan empty. serial0 is the free serial port of the device
	# (serial0 is taken by the serial console).
	port = "serial0"
}
`, comment, ringCount)
}
