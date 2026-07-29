package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testRoutingSettings = "routeros_routing_settings.test_routing_settings"

func TestAccRoutingSettingsTest_basic(t *testing.T) {
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
						Config: testAccRoutingSettingsConfig("3", "11s", "500ms"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testRoutingSettings),
							resource.TestCheckResourceAttr(testRoutingSettings, "check_gateway_ping_count", "3"),
							resource.TestCheckResourceAttr(testRoutingSettings, "check_gateway_ping_interval", "11s"),
							resource.TestCheckResourceAttr(testRoutingSettings, "check_gateway_ping_timeout", "500ms"),
							resource.TestCheckResourceAttr(testRoutingSettings, "single_process", "false"),
							resource.TestCheckResourceAttr(testRoutingSettings, "policy_rules.#", "6"),
							resource.TestCheckResourceAttr(testRoutingSettings, "policy_rules.0", "mangle"),
							resource.TestCheckResourceAttr(testRoutingSettings, "policy_rules.5", "main"),
						),
					},
					{
						Config: testAccRoutingSettingsConfig("2", "10s", "1s"),
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr(testRoutingSettings, "check_gateway_ping_count", "2"),
							resource.TestCheckResourceAttr(testRoutingSettings, "check_gateway_ping_interval", "10s"),
							resource.TestCheckResourceAttr(testRoutingSettings, "check_gateway_ping_timeout", "1s"),
						),
					},
				},
			})

		})
	}
}

func testAccRoutingSettingsConfig(count, interval, timeout string) string {
	return fmt.Sprintf(`%v

resource "routeros_routing_settings" "test_routing_settings" {
	check_gateway_ping_count    = %v
	check_gateway_ping_interval = "%v"
	check_gateway_ping_timeout  = "%v"
	single_process              = false
	policy_rules                = ["mangle", "vrf-lookup", "vrf-unreach", "local", "user", "main"]
}
`, providerConfig, count, interval, timeout)
}
