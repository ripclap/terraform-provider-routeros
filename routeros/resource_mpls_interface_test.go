package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testMplsInterface = "routeros_mpls_interface.test_mpls_interface_x"

func TestAccMplsInterfaceTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/mpls/interface", "routeros_mpls_interface"),
				Steps: []resource.TestStep{
					{
						Config: testAccMplsInterfaceConfig(1508),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testMplsInterface),
							resource.TestCheckResourceAttr(testMplsInterface, "interface", "ether5"),
							resource.TestCheckResourceAttr(testMplsInterface, "input", "true"),
							resource.TestCheckResourceAttr(testMplsInterface, "mpls_mtu", "1508"),
							resource.TestCheckResourceAttr(testMplsInterface, "comment", "test_mpls_interface_x"),
						),
					},
					{
						Config: testAccMplsInterfaceConfig(1512),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testMplsInterface),
							resource.TestCheckResourceAttr(testMplsInterface, "mpls_mtu", "1512"),
						),
					},
				},
			})

		})
	}
}

func testAccMplsInterfaceConfig(mtu int) string {
	return fmt.Sprintf(`%v

resource "routeros_mpls_interface" "test_mpls_interface_x" {
  interface = "ether5"
  input     = true
  mpls_mtu  = %v
  comment   = "test_mpls_interface_x"
}
`, providerConfig, mtu)
}
