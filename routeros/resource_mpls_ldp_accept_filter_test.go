package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testMplsLdpAcceptFilter = "routeros_mpls_ldp_accept_filter.test_mpls_ldp_accept_filter_x"

func TestAccMplsLdpAcceptFilterTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy: testCheckResourceDestroy("/mpls/ldp/accept-filter",
					"routeros_mpls_ldp_accept_filter"),
				Steps: []resource.TestStep{
					{
						Config: testAccMplsLdpAcceptFilterConfig("true", "203.0.113.0/24"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testMplsLdpAcceptFilter),
							resource.TestCheckResourceAttr(testMplsLdpAcceptFilter, "accept", "true"),
							resource.TestCheckResourceAttr(testMplsLdpAcceptFilter, "neighbor", "198.51.100.2"),
							resource.TestCheckResourceAttr(testMplsLdpAcceptFilter, "prefix", "203.0.113.0/24"),
							resource.TestCheckResourceAttr(testMplsLdpAcceptFilter, "vrf", "main"),
							resource.TestCheckResourceAttr(testMplsLdpAcceptFilter, "comment",
								"test_mpls_ldp_accept_filter_x"),
						),
					},
					{
						Config: testAccMplsLdpAcceptFilterConfig("false", "192.0.2.0/24"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testMplsLdpAcceptFilter),
							resource.TestCheckResourceAttr(testMplsLdpAcceptFilter, "accept", "false"),
							resource.TestCheckResourceAttr(testMplsLdpAcceptFilter, "prefix", "192.0.2.0/24"),
						),
					},
				},
			})

		})
	}
}

func testAccMplsLdpAcceptFilterConfig(accept, prefix string) string {
	return fmt.Sprintf(`%v

resource "routeros_mpls_ldp_accept_filter" "test_mpls_ldp_accept_filter_x" {
  accept   = %v
  neighbor = "198.51.100.2"
  prefix   = %q
  vrf      = "main"
  comment  = "test_mpls_ldp_accept_filter_x"
}
`, providerConfig, accept, prefix)
}
