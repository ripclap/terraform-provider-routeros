package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testRoutingIgmpProxyMfc = "routeros_routing_igmp_proxy_mfc.test_routing_igmp_proxy_mfc"

func TestAccRoutingIgmpProxyMfcTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy: testCheckResourceDestroy("/routing/igmp-proxy/mfc",
					"routeros_routing_igmp_proxy_mfc"),
				Steps: []resource.TestStep{
					{
						Config: testAccRoutingIgmpProxyMfcConfig("192.0.2.9", "false"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testRoutingIgmpProxyMfc),
							resource.TestCheckResourceAttr(testRoutingIgmpProxyMfc, "group", "239.9.9.9"),
							resource.TestCheckResourceAttr(testRoutingIgmpProxyMfc, "source", "192.0.2.9"),
							resource.TestCheckResourceAttr(testRoutingIgmpProxyMfc, "upstream_interface", "ether5"),
							resource.TestCheckResourceAttr(testRoutingIgmpProxyMfc, "downstream_interfaces.#", "1"),
							resource.TestCheckTypeSetElemAttr(testRoutingIgmpProxyMfc, "downstream_interfaces.*", "ether6"),
							resource.TestCheckResourceAttr(testRoutingIgmpProxyMfc, "disabled", "false"),
							resource.TestCheckResourceAttr(testRoutingIgmpProxyMfc, "dynamic", "false"),
							resource.TestCheckResourceAttr(testRoutingIgmpProxyMfc, "active", "false"),
						),
					},
					{
						Config: testAccRoutingIgmpProxyMfcConfig("192.0.2.10", "true"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testRoutingIgmpProxyMfc),
							resource.TestCheckResourceAttr(testRoutingIgmpProxyMfc, "source", "192.0.2.10"),
							resource.TestCheckResourceAttr(testRoutingIgmpProxyMfc, "disabled", "true"),
						),
					},
				},
			})

		})
	}
}

func testAccRoutingIgmpProxyMfcConfig(source, disabled string) string {
	return fmt.Sprintf(`%v

resource "routeros_routing_igmp_proxy_mfc" "test_routing_igmp_proxy_mfc" {
  group                 = "239.9.9.9"
  source                = "%v"
  disabled              = %v
  upstream_interface    = "ether5"
  downstream_interfaces = ["ether6"]
}
`, providerConfig, source, disabled)
}
