package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIPv6NdSettings = "routeros_ipv6_nd_settings.test_ipv6_nd_settings_x"

func TestAccIPv6NdSettingsTest_basic(t *testing.T) {
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
						Config: testAccIPv6NdSettingsConfig(`["dns", "mtu"]`, 17),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIPv6NdSettings),
							resource.TestCheckResourceAttr(testIPv6NdSettings,
								"router_advertisement_route_distance", "17"),
							resource.TestCheckResourceAttr(testIPv6NdSettings,
								"router_advertisement_ignored_options.#", "2"),
							resource.TestCheckTypeSetElemAttr(testIPv6NdSettings,
								"router_advertisement_ignored_options.*", "dns"),
							resource.TestCheckTypeSetElemAttr(testIPv6NdSettings,
								"router_advertisement_ignored_options.*", "mtu"),
						),
					},
					// Restore the factory defaults of this settings object.
					{
						Config: testAccIPv6NdSettingsConfig(`[]`, 0),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIPv6NdSettings),
							resource.TestCheckResourceAttr(testIPv6NdSettings,
								"router_advertisement_route_distance", "0"),
							resource.TestCheckResourceAttr(testIPv6NdSettings,
								"router_advertisement_ignored_options.#", "0"),
						),
					},
				},
			})

		})
	}
}

func testAccIPv6NdSettingsConfig(ignoredOptions string, distance int) string {
	return fmt.Sprintf(`%v

resource "routeros_ipv6_nd_settings" "test_ipv6_nd_settings_x" {
  router_advertisement_ignored_options = %v
  router_advertisement_route_distance  = %v
}
`, providerConfig, ignoredOptions, distance)
}
