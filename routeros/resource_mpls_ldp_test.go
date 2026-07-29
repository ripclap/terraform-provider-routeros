package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testMplsLdp = "routeros_mpls_ldp.test_mpls_ldp_x"

func TestAccMplsLdpTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/mpls/ldp", "routeros_mpls_ldp"),
				Steps: []resource.TestStep{
					{
						Config: testAccMplsLdpConfig(255),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testMplsLdp),
							resource.TestCheckResourceAttr(testMplsLdp, "afi.#", "1"),
							resource.TestCheckTypeSetElemAttr(testMplsLdp, "afi.*", "ip"),
							resource.TestCheckResourceAttr(testMplsLdp, "lsr_id", "198.51.100.1"),
							resource.TestCheckResourceAttr(testMplsLdp, "transport_addresses.#", "1"),
							resource.TestCheckResourceAttr(testMplsLdp, "transport_addresses.0", "198.51.100.1"),
							resource.TestCheckResourceAttr(testMplsLdp, "preferred_afi", "ip"),
							resource.TestCheckResourceAttr(testMplsLdp, "hop_limit", "255"),
							resource.TestCheckResourceAttr(testMplsLdp, "path_vector_limit", "255"),
							resource.TestCheckResourceAttr(testMplsLdp, "loop_detect", "false"),
							resource.TestCheckResourceAttr(testMplsLdp, "use_explicit_null", "false"),
							resource.TestCheckResourceAttr(testMplsLdp, "distribute_for_default", "false"),
							resource.TestCheckResourceAttr(testMplsLdp, "vrf", "main"),
							resource.TestCheckResourceAttr(testMplsLdp, "disabled", "true"),
							resource.TestCheckResourceAttr(testMplsLdp, "inactive", "true"),
							resource.TestCheckResourceAttr(testMplsLdp, "comment", "test_mpls_ldp_x"),
						),
					},
					{
						Config: testAccMplsLdpConfig(64),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testMplsLdp),
							resource.TestCheckResourceAttr(testMplsLdp, "hop_limit", "64"),
						),
					},
				},
			})

		})
	}
}

func testAccMplsLdpConfig(hopLimit int) string {
	return fmt.Sprintf(`%v

resource "routeros_mpls_ldp" "test_mpls_ldp_x" {
  afi                    = ["ip"]
  lsr_id                 = "198.51.100.1"
  transport_addresses    = ["198.51.100.1"]
  preferred_afi          = "ip"
  hop_limit              = %v
  path_vector_limit      = 255
  loop_detect            = false
  use_explicit_null      = false
  distribute_for_default = false
  vrf                    = "main"
  disabled               = true
  comment                = "test_mpls_ldp_x"
}
`, providerConfig, hopLimit)
}
