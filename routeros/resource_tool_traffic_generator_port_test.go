package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testResourceToolTrafficGeneratorPort = "routeros_tool_traffic_generator_port.test_tg_port_x"

// Registering a port moves no packets; the test never runs /tool/traffic-generator/start.
func TestAccToolTrafficGeneratorPortTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy: testCheckResourceDestroy("/tool/traffic-generator/port",
					"routeros_tool_traffic_generator_port"),
				Steps: []resource.TestStep{
					{
						Config: testAccToolTrafficGeneratorPortConfig(false),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testResourceToolTrafficGeneratorPort),
							resource.TestCheckResourceAttr(testResourceToolTrafficGeneratorPort, "name", "test_tg_port_x"),
							resource.TestCheckResourceAttr(testResourceToolTrafficGeneratorPort, "interface", "ether6"),
							resource.TestCheckResourceAttr(testResourceToolTrafficGeneratorPort, "disabled", "false"),
							resource.TestCheckResourceAttr(testResourceToolTrafficGeneratorPort, "dynamic", "false"),
							resource.TestCheckResourceAttrSet(testResourceToolTrafficGeneratorPort, "first_header"),
						),
					},
					{
						Config: testAccToolTrafficGeneratorPortConfig(true),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testResourceToolTrafficGeneratorPort),
							resource.TestCheckResourceAttr(testResourceToolTrafficGeneratorPort, "disabled", "true"),
						),
					},
				},
			})
		})
	}
}

func testAccToolTrafficGeneratorPortConfig(disabled bool) string {
	state := "false"
	if disabled {
		state = "true"
	}
	return providerConfig + `
resource "routeros_tool_traffic_generator_port" "test_tg_port_x" {
	name      = "test_tg_port_x"
	interface = "ether6"
	disabled  = ` + state + `
}
`
}
