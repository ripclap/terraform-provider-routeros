package routeros

import (
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testResourceToolTrafficMonitor = "routeros_tool_traffic_monitor.test_tool_traffic_monitor_x"

// The ether6 fixture carries no traffic and `on_event` is empty, so nothing is ever executed.
func TestAccToolTrafficMonitorTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/tool/traffic-monitor", "routeros_tool_traffic_monitor"),
				Steps: []resource.TestStep{
					{
						Config: testAccToolTrafficMonitorConfig("received", "above", 1000000),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testResourceToolTrafficMonitor),
							resource.TestCheckResourceAttr(testResourceToolTrafficMonitor, "name", "test_tool_traffic_monitor_x"),
							resource.TestCheckResourceAttr(testResourceToolTrafficMonitor, "comment", "test_tool_traffic_monitor_x"),
							resource.TestCheckResourceAttr(testResourceToolTrafficMonitor, "interface", "ether6"),
							resource.TestCheckResourceAttr(testResourceToolTrafficMonitor, "traffic", "received"),
							resource.TestCheckResourceAttr(testResourceToolTrafficMonitor, "trigger", "above"),
							resource.TestCheckResourceAttr(testResourceToolTrafficMonitor, "threshold", "1000000"),
							resource.TestCheckResourceAttr(testResourceToolTrafficMonitor, "disabled", "false"),
							resource.TestCheckResourceAttr(testResourceToolTrafficMonitor, "invalid", "false"),
						),
					},
					{
						Config: testAccToolTrafficMonitorConfig("transmitted", "always", 2000000),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testResourceToolTrafficMonitor),
							resource.TestCheckResourceAttr(testResourceToolTrafficMonitor, "traffic", "transmitted"),
							resource.TestCheckResourceAttr(testResourceToolTrafficMonitor, "trigger", "always"),
							resource.TestCheckResourceAttr(testResourceToolTrafficMonitor, "threshold", "2000000"),
						),
					},
				},
			})
		})
	}
}

func testAccToolTrafficMonitorConfig(traffic, trigger string, threshold int) string {
	return providerConfig + `
resource "routeros_tool_traffic_monitor" "test_tool_traffic_monitor_x" {
	name      = "test_tool_traffic_monitor_x"
	comment   = "test_tool_traffic_monitor_x"
	interface = "ether6"
	traffic   = "` + traffic + `"
	trigger   = "` + trigger + `"
	threshold = ` + strconv.Itoa(threshold) + `
	disabled  = false
}
`
}
