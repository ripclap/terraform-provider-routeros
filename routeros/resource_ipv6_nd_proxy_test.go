package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIPv6NdProxy = "routeros_ipv6_nd_proxy.test_ipv6_nd_proxy_x"

// Uses ether1, not a veth fixture: on RouterOS 7.23.2 `/ipv6/nd/proxy/remove` fails with
// "could not get interface index for device" when the interface is down, blocking destroy.
func TestAccIPv6NdProxyTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ipv6/nd/proxy", "routeros_ipv6_nd_proxy"),
				Steps: []resource.TestStep{
					{
						Config: testAccIPv6NdProxyConfig("2001:db8:7f00:73::1", "false"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIPv6NdProxy),
							resource.TestCheckResourceAttr(testIPv6NdProxy, "address", "2001:db8:7f00:73::1"),
							resource.TestCheckResourceAttr(testIPv6NdProxy, "interface", "ether1"),
							resource.TestCheckResourceAttr(testIPv6NdProxy, "disabled", "false"),
							resource.TestCheckResourceAttr(testIPv6NdProxy, "comment", "test_ipv6_nd_proxy_x"),
						),
					},
					{
						Config: testAccIPv6NdProxyConfig("2001:db8:7f00:73::2", "true"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIPv6NdProxy),
							resource.TestCheckResourceAttr(testIPv6NdProxy, "address", "2001:db8:7f00:73::2"),
							resource.TestCheckResourceAttr(testIPv6NdProxy, "disabled", "true"),
						),
					},
				},
			})

		})
	}
}

func testAccIPv6NdProxyConfig(address, disabled string) string {
	return fmt.Sprintf(`%v

resource "routeros_ipv6_nd_proxy" "test_ipv6_nd_proxy_x" {
  address   = %q
  interface = "ether1"
  disabled  = %v
  comment   = "test_ipv6_nd_proxy_x"
}
`, providerConfig, address, disabled)
}
