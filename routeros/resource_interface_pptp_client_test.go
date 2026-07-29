package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testInterfacePptpClientAddress = "routeros_interface_pptp_client.test_pptp_client_x"

func TestAccInterfacePptpClientTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/interface/pptp-client", "routeros_interface_pptp_client"),
				Steps: []resource.TestStep{
					{
						Config: testAccInterfacePptpClientConfig("192.0.2.55", "test_pptp_client_x"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfacePptpClientAddress),
							resource.TestCheckResourceAttr(testInterfacePptpClientAddress, "name", "test_pptp_client_x"),
							resource.TestCheckResourceAttr(testInterfacePptpClientAddress, "connect_to", "192.0.2.55"),
							resource.TestCheckResourceAttr(testInterfacePptpClientAddress, "comment", "test_pptp_client_x"),
							resource.TestCheckResourceAttr(testInterfacePptpClientAddress, "user", "test_pptp_client_x_user"),
							resource.TestCheckResourceAttr(testInterfacePptpClientAddress, "password", "test_pptp_client_x_pw"),
							resource.TestCheckResourceAttr(testInterfacePptpClientAddress, "add_default_route", "false"),
							resource.TestCheckResourceAttr(testInterfacePptpClientAddress, "dial_on_demand", "false"),
							resource.TestCheckResourceAttr(testInterfacePptpClientAddress, "disabled", "true"),
							resource.TestCheckResourceAttr(testInterfacePptpClientAddress, "profile", "default-encryption"),
							resource.TestCheckResourceAttr(testInterfacePptpClientAddress, "allow.#", "2"),
						),
					},
					{
						Config: testAccInterfacePptpClientConfig("192.0.2.56", "test_pptp_client_x updated"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfacePptpClientAddress),
							resource.TestCheckResourceAttr(testInterfacePptpClientAddress, "connect_to", "192.0.2.56"),
							resource.TestCheckResourceAttr(testInterfacePptpClientAddress, "comment", "test_pptp_client_x updated"),
						),
					},
				},
			})

		})
	}
}

func testAccInterfacePptpClientConfig(connectTo, comment string) string {
	return providerConfig + fmt.Sprintf(`
resource "routeros_interface_pptp_client" "test_pptp_client_x" {
	name              = "test_pptp_client_x"
	connect_to        = "%v"
	comment           = "%v"
	user              = "test_pptp_client_x_user"
	password          = "test_pptp_client_x_pw"
	add_default_route = false
	dial_on_demand    = false
	disabled          = true
	profile           = "default-encryption"
	allow             = ["mschap1", "mschap2"]
	keepalive_timeout = "60"
	max_mru           = 1450
	max_mtu           = 1450
	mrru              = "disabled"
	use_peer_dns      = "no"
}
`, connectTo, comment)
}
