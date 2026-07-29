package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIPv6FirewallRaw = "routeros_ipv6_firewall_raw.test_ipv6_firewall_raw_x"

func TestAccIPv6FirewallRawTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ipv6/firewall/raw", "routeros_ipv6_firewall_raw"),
				Steps: []resource.TestStep{
					{
						Config: testAccIPv6FirewallRawConfig("443"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIPv6FirewallRaw),
							resource.TestCheckResourceAttr(testIPv6FirewallRaw, "action", "drop"),
							resource.TestCheckResourceAttr(testIPv6FirewallRaw, "chain", "prerouting"),
							resource.TestCheckResourceAttr(testIPv6FirewallRaw, "src_address", "2001:db8:7f00:72::1/128"),
							resource.TestCheckResourceAttr(testIPv6FirewallRaw, "dst_address", "2001:db8:7f00:72::2/128"),
							resource.TestCheckResourceAttr(testIPv6FirewallRaw, "protocol", "tcp"),
							resource.TestCheckResourceAttr(testIPv6FirewallRaw, "dst_port", "443"),
							resource.TestCheckResourceAttr(testIPv6FirewallRaw, "hop_limit", "equal:64"),
							resource.TestCheckResourceAttr(testIPv6FirewallRaw, "log", "true"),
							resource.TestCheckResourceAttr(testIPv6FirewallRaw, "log_prefix", "test_ipv6_raw_x"),
							resource.TestCheckResourceAttr(testIPv6FirewallRaw, "comment",
								"test_ipv6_firewall_raw_x"),
						),
					},
					{
						Config: testAccIPv6FirewallRawConfig("8443"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIPv6FirewallRaw),
							resource.TestCheckResourceAttr(testIPv6FirewallRaw, "dst_port", "8443"),
						),
					},
				},
			})

		})
	}
}

func testAccIPv6FirewallRawConfig(dstPort string) string {
	return fmt.Sprintf(`%v

resource "routeros_ipv6_firewall_raw" "test_ipv6_firewall_raw_x" {
  action      = "drop"
  chain       = "prerouting"
  src_address = "2001:db8:7f00:72::1/128"
  dst_address = "2001:db8:7f00:72::2/128"
  protocol    = "tcp"
  dst_port    = %q
  hop_limit   = "equal:64"
  log         = true
  log_prefix  = "test_ipv6_raw_x"
  comment     = "test_ipv6_firewall_raw_x"
}
`, providerConfig, dstPort)
}
