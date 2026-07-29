package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testPortRemoteAccessAddress = "routeros_port_remote_access.test_port_remote_access_x"

func TestAccPortRemoteAccessTest_basic(t *testing.T) {
	testCheckMenu(t, "/port")
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy: testCheckResourceDestroy("/port/remote-access",
					"routeros_port_remote_access"),
				Steps: []resource.TestStep{
					{
						Config: testAccPortRemoteAccessConfig("32108", "tcp-server"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testPortRemoteAccessAddress),
							resource.TestCheckResourceAttr(testPortRemoteAccessAddress, "port", "serial0"),
							resource.TestCheckResourceAttr(testPortRemoteAccessAddress, "channel", "0"),
							resource.TestCheckResourceAttr(testPortRemoteAccessAddress, "ip_port", "32108"),
							resource.TestCheckResourceAttr(testPortRemoteAccessAddress, "protocol", "tcp-server"),
							resource.TestCheckResourceAttr(testPortRemoteAccessAddress, "disabled", "false"),
							resource.TestCheckResourceAttr(testPortRemoteAccessAddress, "comment",
								"test_port_remote_access_x"),
						),
					},
					{
						Config: testAccPortRemoteAccessConfig("32109", "rfc2217"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testPortRemoteAccessAddress),
							resource.TestCheckResourceAttr(testPortRemoteAccessAddress, "ip_port", "32109"),
							resource.TestCheckResourceAttr(testPortRemoteAccessAddress, "protocol", "rfc2217"),
						),
					},
				},
			})

		})
	}
}

func testAccPortRemoteAccessConfig(ipPort, protocol string) string {
	return providerConfig + `

resource "routeros_port_remote_access" "test_port_remote_access_x" {
	port     = "serial0"
	channel  = 0
	ip_port  = ` + ipPort + `
	protocol = "` + protocol + `"
	disabled = false
	comment  = "test_port_remote_access_x"
}

`
}
