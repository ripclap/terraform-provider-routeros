package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testInterfacePppClientAddress = "routeros_interface_ppp_client.test_ppp_client_x"

func TestAccInterfacePppClientTest_basic(t *testing.T) {
	testCheckMenu(t, "/port")
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/interface/ppp-client", "routeros_interface_ppp_client"),
				Steps: []resource.TestStep{
					{
						Config: testAccInterfacePppClientConfig("test_ppp_client_x", "*99#"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfacePppClientAddress),
							resource.TestCheckResourceAttr(testInterfacePppClientAddress, "name", "test_ppp_client_x"),
							resource.TestCheckResourceAttr(testInterfacePppClientAddress, "comment", "test_ppp_client_x"),
							resource.TestCheckResourceAttr(testInterfacePppClientAddress, "phone", "*99#"),
							resource.TestCheckResourceAttr(testInterfacePppClientAddress, "user", "test_ppp_client_x_user"),
							resource.TestCheckResourceAttr(testInterfacePppClientAddress, "password", "test_ppp_client_x_pw"),
							resource.TestCheckResourceAttr(testInterfacePppClientAddress, "disabled", "true"),
							resource.TestCheckResourceAttr(testInterfacePppClientAddress, "add_default_route", "false"),
							resource.TestCheckResourceAttr(testInterfacePppClientAddress, "dial_on_demand", "false"),
							resource.TestCheckResourceAttr(testInterfacePppClientAddress, "null_modem", "true"),
							resource.TestCheckResourceAttr(testInterfacePppClientAddress, "profile", "default"),
							resource.TestCheckResourceAttr(testInterfacePppClientAddress, "port", "serial0"),
							resource.TestCheckResourceAttr(testInterfacePppClientAddress, "default_route_distance", "3"),
							resource.TestCheckResourceAttr(testInterfacePppClientAddress, "use_peer_dns", "false"),
						),
					},
					{
						Config: testAccInterfacePppClientConfig("test_ppp_client_x updated", "*98#"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfacePppClientAddress),
							resource.TestCheckResourceAttr(testInterfacePppClientAddress, "comment", "test_ppp_client_x updated"),
							resource.TestCheckResourceAttr(testInterfacePppClientAddress, "phone", "*98#"),
						),
					},
				},
			})

		})
	}
}

func testAccInterfacePppClientConfig(comment, phone string) string {
	return providerConfig + fmt.Sprintf(`
resource "routeros_interface_ppp_client" "test_ppp_client_x" {
	name              = "test_ppp_client_x"
	comment           = "%v"
	phone             = "%v"
	user              = "test_ppp_client_x_user"
	password          = "test_ppp_client_x_pw"
	disabled          = true
	add_default_route = false
	dial_on_demand    = false
	null_modem        = true
	profile           = "default"
	keepalive_timeout = "30"
	max_mru           = 1500
	max_mtu           = 1500
	mrru              = "disabled"
	use_peer_dns      = false

	dial_command           = "ATDT"
	default_route_distance = 3
	modem_init             = ""

	# 'port' has no diff suppression in the schema, and RouterOS always reports the
	# placeholder '*FFFFFFFF' when no serial port is bound, so the port has to be set
	# explicitly to keep the plan empty. serial0 is the free serial port of the device
	# (serial0 is taken by the serial console).
	port = "serial0"
}
`, comment, phone)
}
