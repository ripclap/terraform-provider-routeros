package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIpDhcpServerAlertAddress = "routeros_ip_dhcp_server_alert.test_dhcp_server_alert_x"

func TestAccIpDhcpServerAlertTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ip/dhcp-server/alert", "routeros_ip_dhcp_server_alert"),
				Steps: []resource.TestStep{
					{
						Config: testAccIpDhcpServerAlertConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpDhcpServerAlertAddress),
							resource.TestCheckResourceAttr(testIpDhcpServerAlertAddress, "interface", "ether3"),
							resource.TestCheckResourceAttr(testIpDhcpServerAlertAddress, "alert_timeout", "1h"),
							resource.TestCheckResourceAttr(testIpDhcpServerAlertAddress, "valid_server", "00:11:22:33:44:55"),
							resource.TestCheckResourceAttr(testIpDhcpServerAlertAddress, "comment", "test_dhcp_server_alert_x"),
							resource.TestCheckResourceAttr(testIpDhcpServerAlertAddress, "disabled", "true"),
						),
					},
					{
						Config: testAccIpDhcpServerAlertUpdatedConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpDhcpServerAlertAddress),
							resource.TestCheckResourceAttr(testIpDhcpServerAlertAddress, "alert_timeout", "none"),
							resource.TestCheckResourceAttr(testIpDhcpServerAlertAddress, "valid_server", "00:11:22:33:44:55,00:11:22:33:44:66"),
							resource.TestCheckResourceAttr(testIpDhcpServerAlertAddress, "comment", "test_dhcp_server_alert_x_upd"),
						),
					},
				},
			})

		})
	}
}

func testAccIpDhcpServerAlertConfig() string {
	return providerConfig + `
resource "routeros_ip_dhcp_server_alert" "test_dhcp_server_alert_x" {
	interface     = "ether3"
	alert_timeout = "1h"
	valid_server  = "00:11:22:33:44:55"
	comment       = "test_dhcp_server_alert_x"
	disabled      = true
}
`
}

func testAccIpDhcpServerAlertUpdatedConfig() string {
	return providerConfig + `
resource "routeros_ip_dhcp_server_alert" "test_dhcp_server_alert_x" {
	interface     = "ether3"
	alert_timeout = "none"
	valid_server  = "00:11:22:33:44:55,00:11:22:33:44:66"
	comment       = "test_dhcp_server_alert_x_upd"
	disabled      = true
}
`
}
