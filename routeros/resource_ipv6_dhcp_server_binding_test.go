package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIPv6DhcpServerBinding = "routeros_ipv6_dhcp_server_binding.test_ipv6_dhcp_server_binding_x"

func TestAccIPv6DhcpServerBindingTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy: testCheckResourceDestroy("/ipv6/dhcp-server/binding",
					"routeros_ipv6_dhcp_server_binding"),
				Steps: []resource.TestStep{
					{
						Config: testAccIPv6DhcpServerBindingConfig("1w"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIPv6DhcpServerBinding),
							resource.TestCheckResourceAttr(testIPv6DhcpServerBinding, "duid",
								"0x0003000102ca7f000071"),
							resource.TestCheckResourceAttr(testIPv6DhcpServerBinding, "address",
								"2001:db8:7f00:71::100/128"),
							resource.TestCheckResourceAttr(testIPv6DhcpServerBinding, "ia_type", "na"),
							resource.TestCheckResourceAttr(testIPv6DhcpServerBinding, "iaid", "71"),
							resource.TestCheckResourceAttr(testIPv6DhcpServerBinding, "life_time", "1w"),
							resource.TestCheckResourceAttr(testIPv6DhcpServerBinding, "comment",
								"test_ipv6_dhcp_server_binding_x"),
							resource.TestCheckResourceAttr(testIPv6DhcpServerBinding, "server", "all"),
							resource.TestCheckResourceAttr(testIPv6DhcpServerBinding, "status", "waiting"),
						),
					},
					{
						Config: testAccIPv6DhcpServerBindingConfig("2d"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIPv6DhcpServerBinding),
							resource.TestCheckResourceAttr(testIPv6DhcpServerBinding, "life_time", "2d"),
						),
					},
				},
			})

		})
	}
}

func testAccIPv6DhcpServerBindingConfig(lifeTime string) string {
	return fmt.Sprintf(`%v

resource "routeros_ipv6_dhcp_server_binding" "test_ipv6_dhcp_server_binding_x" {
  duid      = "0x0003000102ca7f000071"
  address   = "2001:db8:7f00:71::100/128"
  ia_type   = "na"
  iaid      = 71
  life_time = %q
  comment   = "test_ipv6_dhcp_server_binding_x"
}
`, providerConfig, lifeTime)
}
