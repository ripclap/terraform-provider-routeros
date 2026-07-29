package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testMplsLdpAdvertiseFilter = "routeros_mpls_ldp_advertise_filter.test_mpls_ldp_advertise_filter_x"

func TestAccMplsLdpAdvertiseFilterTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy: testCheckResourceDestroy("/mpls/ldp/advertise-filter",
					"routeros_mpls_ldp_advertise_filter"),
				Steps: []resource.TestStep{
					{
						Config: testAccMplsLdpAdvertiseFilterConfig("true", "203.0.113.0/24"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testMplsLdpAdvertiseFilter),
							resource.TestCheckResourceAttr(testMplsLdpAdvertiseFilter, "advertise", "true"),
							resource.TestCheckResourceAttr(testMplsLdpAdvertiseFilter, "neighbor", "198.51.100.3"),
							resource.TestCheckResourceAttr(testMplsLdpAdvertiseFilter, "prefix", "203.0.113.0/24"),
							resource.TestCheckResourceAttr(testMplsLdpAdvertiseFilter, "vrf", "any"),
							resource.TestCheckResourceAttr(testMplsLdpAdvertiseFilter, "comment",
								"test_mpls_ldp_advertise_filter_x"),
						),
					},
					{
						Config: testAccMplsLdpAdvertiseFilterConfig("false", "192.0.2.0/24"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testMplsLdpAdvertiseFilter),
							resource.TestCheckResourceAttr(testMplsLdpAdvertiseFilter, "advertise", "false"),
							resource.TestCheckResourceAttr(testMplsLdpAdvertiseFilter, "prefix", "192.0.2.0/24"),
						),
					},
				},
			})

		})
	}
}

func testAccMplsLdpAdvertiseFilterConfig(advertise, prefix string) string {
	return fmt.Sprintf(`%v

resource "routeros_mpls_ldp_advertise_filter" "test_mpls_ldp_advertise_filter_x" {
  advertise = %v
  neighbor  = "198.51.100.3"
  prefix    = %q
  # RouterOS 7.23.2 only accepts "any" here: unlike /mpls/ldp/accept-filter, this menu
  # rejects a concrete VRF name with "input does not match any value of vrf".
  vrf     = "any"
  comment = "test_mpls_ldp_advertise_filter_x"
}
`, providerConfig, advertise, prefix)
}
