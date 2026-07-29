package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testInterfacePptpServerBinding = "routeros_interface_pptp_server.test_pptp_server_binding_x"

func TestAccInterfacePptpServerTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/interface/pptp-server", "routeros_interface_pptp_server"),
				Steps: []resource.TestStep{
					{
						Config: testAccInterfacePptpServerConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfacePptpServerBinding),
							resource.TestCheckResourceAttr(testInterfacePptpServerBinding, "name", "test_pptp_server_binding_x"),
							resource.TestCheckResourceAttr(testInterfacePptpServerBinding, "user", "test_pptp_srv_user_x"),
							resource.TestCheckResourceAttr(testInterfacePptpServerBinding, "comment", "test_pptp_server_binding_x"),
							resource.TestCheckResourceAttr(testInterfacePptpServerBinding, "disabled", "false"),
							resource.TestCheckResourceAttr(testInterfacePptpServerBinding, "running", "false"),
						),
					},
					{
						Config: testAccInterfacePptpServerUpdatedConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfacePptpServerBinding),
							resource.TestCheckResourceAttr(testInterfacePptpServerBinding, "user", "test_pptp_srv_user_x2"),
							resource.TestCheckResourceAttr(testInterfacePptpServerBinding, "comment", "updated"),
							resource.TestCheckResourceAttr(testInterfacePptpServerBinding, "disabled", "true"),
						),
					},
				},
			})

		})
	}
}

func testAccInterfacePptpServerConfig() string {
	return providerConfig + `

resource "routeros_interface_pptp_server" "test_pptp_server_binding_x" {
	name    = "test_pptp_server_binding_x"
	user    = "test_pptp_srv_user_x"
	comment = "test_pptp_server_binding_x"
}
`
}

func testAccInterfacePptpServerUpdatedConfig() string {
	return providerConfig + `

resource "routeros_interface_pptp_server" "test_pptp_server_binding_x" {
	name     = "test_pptp_server_binding_x"
	user     = "test_pptp_srv_user_x2"
	comment  = "updated"
	disabled = true
}
`
}
