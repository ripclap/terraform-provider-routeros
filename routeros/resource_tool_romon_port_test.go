package routeros

import (
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testResourceToolRomonPort = "routeros_tool_romon_port.test_tool_romon_port_x"

// RoMON only runs on ethernet-like ports; other interfaces are rejected with
// "input does not match any value of interface", so the test uses the existing idle ether6 fixture.
func TestAccToolRomonPortTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/tool/romon/port", "routeros_tool_romon_port"),
				Steps: []resource.TestStep{
					{
						Config: testAccToolRomonPortConfig(200, false),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testResourceToolRomonPort),
							resource.TestCheckResourceAttr(testResourceToolRomonPort, "interface", "ether6"),
							resource.TestCheckResourceAttr(testResourceToolRomonPort, "comment", "test_tool_romon_port_x"),
							resource.TestCheckResourceAttr(testResourceToolRomonPort, "cost", "200"),
							resource.TestCheckResourceAttr(testResourceToolRomonPort, "forbid", "false"),
							resource.TestCheckResourceAttr(testResourceToolRomonPort, "disabled", "false"),
							resource.TestCheckResourceAttr(testResourceToolRomonPort, "dynamic", "false"),
							resource.TestCheckResourceAttr(testResourceToolRomonPort, "secrets", "test_tool_romon_port_secret"),
						),
					},
					{
						Config: testAccToolRomonPortConfig(300, true),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testResourceToolRomonPort),
							resource.TestCheckResourceAttr(testResourceToolRomonPort, "cost", "300"),
							resource.TestCheckResourceAttr(testResourceToolRomonPort, "forbid", "true"),
						),
					},
				},
			})
		})
	}
}

func testAccToolRomonPortConfig(cost int, forbid bool) string {
	forbidState := "false"
	if forbid {
		forbidState = "true"
	}
	return providerConfig + `
resource "routeros_tool_romon_port" "test_tool_romon_port_x" {
	interface = "ether6"
	comment   = "test_tool_romon_port_x"
	cost      = ` + strconv.Itoa(cost) + `
	forbid    = ` + forbidState + `
	disabled  = false
	secrets   = "test_tool_romon_port_secret"
}
`
}
