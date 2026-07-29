package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testMplsLdpInterface = "routeros_mpls_ldp_interface.test_mpls_ldp_interface_x"

func TestAccMplsLdpInterfaceTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy: testCheckResourceDestroy("/mpls/ldp/interface",
					"routeros_mpls_ldp_interface"),
				Steps: []resource.TestStep{
					{
						Config: testAccMplsLdpInterfaceConfig("5s", "15s"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testMplsLdpInterface),
							resource.TestCheckResourceAttr(testMplsLdpInterface, "interface", "ether5"),
							resource.TestCheckResourceAttr(testMplsLdpInterface, "afi.#", "1"),
							resource.TestCheckTypeSetElemAttr(testMplsLdpInterface, "afi.*", "ip"),
							resource.TestCheckResourceAttr(testMplsLdpInterface,
								"accept_dynamic_neighbors", "true"),
							resource.TestCheckResourceAttr(testMplsLdpInterface, "hello_interval", "5s"),
							resource.TestCheckResourceAttr(testMplsLdpInterface, "hold_time", "15s"),
							resource.TestCheckResourceAttr(testMplsLdpInterface, "disabled", "true"),
							resource.TestCheckResourceAttr(testMplsLdpInterface, "comment",
								"test_mpls_ldp_interface_x"),
						),
					},
					{
						Config: testAccMplsLdpInterfaceConfig("10s", "30s"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testMplsLdpInterface),
							resource.TestCheckResourceAttr(testMplsLdpInterface, "hello_interval", "10s"),
							resource.TestCheckResourceAttr(testMplsLdpInterface, "hold_time", "30s"),
						),
					},
				},
			})

		})
	}
}

func testAccMplsLdpInterfaceConfig(helloInterval, holdTime string) string {
	return fmt.Sprintf(`%v

resource "routeros_mpls_ldp_interface" "test_mpls_ldp_interface_x" {
  interface                = "ether5"
  afi                      = ["ip"]
  accept_dynamic_neighbors = true
  hello_interval           = %q
  hold_time                = %q
  # Kept disabled so that the test never emits LDP hello messages.
  disabled = true
  comment  = "test_mpls_ldp_interface_x"
}
`, providerConfig, helloInterval, holdTime)
}
