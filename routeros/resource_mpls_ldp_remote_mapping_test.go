package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testMplsLdpRemoteMappingAddress = "routeros_mpls_ldp_remote_mapping.test_mpls_ldp_remote_mapping_x"

func TestAccMplsLdpRemoteMappingTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy: testCheckResourceDestroy("/mpls/ldp/remote-mapping",
					"routeros_mpls_ldp_remote_mapping"),
				Steps: []resource.TestStep{
					{
						Config: testAccMplsLdpRemoteMappingConfig("1608"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testMplsLdpRemoteMappingAddress),
							resource.TestCheckResourceAttr(testMplsLdpRemoteMappingAddress, "dst_address", "203.0.113.0/24"),
							resource.TestCheckResourceAttr(testMplsLdpRemoteMappingAddress, "label", "1608"),
							resource.TestCheckResourceAttr(testMplsLdpRemoteMappingAddress, "nexthop", "203.0.113.1"),
							resource.TestCheckResourceAttr(testMplsLdpRemoteMappingAddress, "comment",
								"test_mpls_ldp_remote_mapping_x"),
						),
					},
					{
						Config: testAccMplsLdpRemoteMappingConfig("1609"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testMplsLdpRemoteMappingAddress),
							resource.TestCheckResourceAttr(testMplsLdpRemoteMappingAddress, "label", "1609"),
						),
					},
				},
			})

		})
	}
}

func testAccMplsLdpRemoteMappingConfig(label string) string {
	return providerConfig + `

resource "routeros_mpls_ldp_remote_mapping" "test_mpls_ldp_remote_mapping_x" {
	dst_address = "203.0.113.0/24"
	label       = "` + label + `"
	nexthop     = "203.0.113.1"
	comment     = "test_mpls_ldp_remote_mapping_x"
}

`
}
