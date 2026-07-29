package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIPv6NdPrefixDefault = "routeros_ipv6_nd_prefix_default.test_ipv6_nd_prefix_default_x"

func TestAccIPv6NdPrefixDefaultTest_basic(t *testing.T) {
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
						Config: testAccIPv6NdPrefixDefaultConfig("false", "true", "2d", "1w"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIPv6NdPrefixDefault),
							resource.TestCheckResourceAttr(testIPv6NdPrefixDefault, "autonomous", "false"),
							resource.TestCheckResourceAttr(testIPv6NdPrefixDefault, "dhcp6_pd_preferred", "true"),
							resource.TestCheckResourceAttr(testIPv6NdPrefixDefault, "preferred_lifetime", "2d"),
							resource.TestCheckResourceAttr(testIPv6NdPrefixDefault, "valid_lifetime", "1w"),
						),
					},
					// Restore the factory defaults of this settings object.
					{
						Config: testAccIPv6NdPrefixDefaultConfig("true", "false", "1w", "4w2d"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIPv6NdPrefixDefault),
							resource.TestCheckResourceAttr(testIPv6NdPrefixDefault, "autonomous", "true"),
							resource.TestCheckResourceAttr(testIPv6NdPrefixDefault, "dhcp6_pd_preferred", "false"),
							resource.TestCheckResourceAttr(testIPv6NdPrefixDefault, "preferred_lifetime", "1w"),
							resource.TestCheckResourceAttr(testIPv6NdPrefixDefault, "valid_lifetime", "4w2d"),
						),
					},
				},
			})

		})
	}
}

func testAccIPv6NdPrefixDefaultConfig(autonomous, pdPreferred, preferred, valid string) string {
	return fmt.Sprintf(`%v

resource "routeros_ipv6_nd_prefix_default" "test_ipv6_nd_prefix_default_x" {
  autonomous         = %v
  dhcp6_pd_preferred = %v
  preferred_lifetime = %q
  valid_lifetime     = %q
}
`, providerConfig, autonomous, pdPreferred, preferred, valid)
}
