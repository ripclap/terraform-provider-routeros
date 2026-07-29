package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// The connection tracking helpers are a fixed list that can be modified but not added or removed,
// so the test carries no CheckDestroy.
const testIPFirewallServicePortRtsp = "routeros_ip_firewall_service_port.test_service_port_rtsp_x"

func TestAccIPFirewallServicePortTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				Steps: []resource.TestStep{
					{
						Config: testAccIPFirewallServicePortConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIPFirewallServicePortRtsp),
							resource.TestCheckResourceAttr(testIPFirewallServicePortRtsp, "name", "rtsp"),
							resource.TestCheckResourceAttr(testIPFirewallServicePortRtsp, "ports", "554,8554"),
							resource.TestCheckResourceAttr(testIPFirewallServicePortRtsp, "disabled", "true"),
						),
					},
					{
						// Restores the RouterOS default state of the helper.
						Config: testAccIPFirewallServicePortRestoreConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIPFirewallServicePortRtsp),
							resource.TestCheckResourceAttr(testIPFirewallServicePortRtsp, "ports", "554"),
							resource.TestCheckResourceAttr(testIPFirewallServicePortRtsp, "disabled", "true"),
						),
					},
				},
			})
		})
	}
}

func testAccIPFirewallServicePortConfig() string {
	return providerConfig + `
resource "routeros_ip_firewall_service_port" "test_service_port_rtsp_x" {
	name     = "rtsp"
	ports    = "554,8554"
	disabled = true
}
`
}

func testAccIPFirewallServicePortRestoreConfig() string {
	return providerConfig + `
resource "routeros_ip_firewall_service_port" "test_service_port_rtsp_x" {
	name     = "rtsp"
	ports    = "554"
	disabled = true
}
`
}
