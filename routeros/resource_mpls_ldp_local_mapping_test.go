package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testMplsLdpLocalMapping = "routeros_mpls_ldp_local_mapping.test_mpls_ldp_local_mapping_x"

func TestAccMplsLdpLocalMappingTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy: testCheckResourceDestroy("/mpls/ldp/local-mapping",
					"routeros_mpls_ldp_local_mapping"),
				Steps: []resource.TestStep{
					{
						Config: testAccMplsLdpLocalMappingConfig("impl-null"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testMplsLdpLocalMapping),
							resource.TestCheckResourceAttr(testMplsLdpLocalMapping, "dst_address",
								"203.0.113.0/24"),
							resource.TestCheckResourceAttr(testMplsLdpLocalMapping, "label", "impl-null"),
							resource.TestCheckResourceAttr(testMplsLdpLocalMapping, "vrf", "main"),
							resource.TestCheckResourceAttr(testMplsLdpLocalMapping, "inactive", "true"),
							resource.TestCheckResourceAttr(testMplsLdpLocalMapping, "comment",
								"test_mpls_ldp_local_mapping_x"),
						),
					},
					{
						Config: testAccMplsLdpLocalMappingConfig("1077"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testMplsLdpLocalMapping),
							resource.TestCheckResourceAttr(testMplsLdpLocalMapping, "label", "1077"),
						),
					},
				},
			})

		})
	}
}

func testAccMplsLdpLocalMappingConfig(label string) string {
	return fmt.Sprintf(`%v

resource "routeros_mpls_ldp_local_mapping" "test_mpls_ldp_local_mapping_x" {
  dst_address = "203.0.113.0/24"
  label       = %q
  vrf         = "main"
  comment     = "test_mpls_ldp_local_mapping_x"
}
`, providerConfig, label)
}
