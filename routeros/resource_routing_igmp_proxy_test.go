package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testRoutingIgmpProxy = "routeros_routing_igmp_proxy.test_routing_igmp_proxy"

// /routing/igmp-proxy is a settings singleton: it cannot be created or destroyed, so the resource uses the
// DefaultSystem* handlers and the test does not declare a CheckDestroy. The last step restores the RouterOS
// factory values.
func TestAccRoutingIgmpProxyTest_basic(t *testing.T) {
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
						Config: testAccRoutingIgmpProxyConfig("3m20s", "12s", "true"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testRoutingIgmpProxy),
							resource.TestCheckResourceAttr(testRoutingIgmpProxy, "query_interval", "3m20s"),
							resource.TestCheckResourceAttr(testRoutingIgmpProxy, "query_response_interval", "12s"),
							resource.TestCheckResourceAttr(testRoutingIgmpProxy, "quick_leave", "true"),
						),
					},
					{
						// RouterOS factory defaults.
						Config: testAccRoutingIgmpProxyConfig("2m5s", "10s", "false"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testRoutingIgmpProxy),
							resource.TestCheckResourceAttr(testRoutingIgmpProxy, "query_interval", "2m5s"),
							resource.TestCheckResourceAttr(testRoutingIgmpProxy, "query_response_interval", "10s"),
							resource.TestCheckResourceAttr(testRoutingIgmpProxy, "quick_leave", "false"),
						),
					},
				},
			})

		})
	}
}

func testAccRoutingIgmpProxyConfig(queryInterval, queryResponseInterval, quickLeave string) string {
	return fmt.Sprintf(`%v

resource "routeros_routing_igmp_proxy" "test_routing_igmp_proxy" {
  query_interval          = "%v"
  query_response_interval = "%v"
  quick_leave             = %v
}
`, providerConfig, queryInterval, queryResponseInterval, quickLeave)
}
