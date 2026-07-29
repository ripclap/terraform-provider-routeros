package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testInterfaceL2tpEtherAddress = "routeros_interface_l2tp_ether.test_l2tp_ether_x"

func TestAccInterfaceL2tpEtherTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/interface/l2tp-ether", "routeros_interface_l2tp_ether"),
				Steps: []resource.TestStep{
					{
						Config: testAccInterfaceL2tpEtherConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfaceL2tpEtherAddress),
							resource.TestCheckResourceAttr(testInterfaceL2tpEtherAddress, "name", "test_l2tp_ether_x"),
							resource.TestCheckResourceAttr(testInterfaceL2tpEtherAddress, "connect_to", "192.0.2.55"),
							resource.TestCheckResourceAttr(testInterfaceL2tpEtherAddress, "comment", "test_l2tp_ether_x"),
							resource.TestCheckResourceAttr(testInterfaceL2tpEtherAddress, "mtu", "1400"),
							resource.TestCheckResourceAttr(testInterfaceL2tpEtherAddress, "circuit_id", "testl2tpetherx"),
							resource.TestCheckResourceAttr(testInterfaceL2tpEtherAddress, "allow_fast_path", "true"),
							resource.TestCheckResourceAttr(testInterfaceL2tpEtherAddress, "use_l2_specific_sublayer", "true"),
							resource.TestCheckResourceAttr(testInterfaceL2tpEtherAddress, "disabled", "false"),
						),
					},
					{
						Config: testAccInterfaceL2tpEtherConfigUpdated(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfaceL2tpEtherAddress),
							resource.TestCheckResourceAttr(testInterfaceL2tpEtherAddress, "connect_to", "192.0.2.56"),
							resource.TestCheckResourceAttr(testInterfaceL2tpEtherAddress, "comment", "test_l2tp_ether_x updated"),
							resource.TestCheckResourceAttr(testInterfaceL2tpEtherAddress, "mtu", "1300"),
							resource.TestCheckResourceAttr(testInterfaceL2tpEtherAddress, "allow_fast_path", "false"),
							resource.TestCheckResourceAttr(testInterfaceL2tpEtherAddress, "use_l2_specific_sublayer", "false"),
							resource.TestCheckResourceAttr(testInterfaceL2tpEtherAddress, "disabled", "true"),
						),
					},
				},
			})

		})
	}
}

func testAccInterfaceL2tpEtherConfig() string {
	return providerConfig + `
resource "routeros_interface_l2tp_ether" "test_l2tp_ether_x" {
	name                     = "test_l2tp_ether_x"
	connect_to               = "192.0.2.55"
	comment                  = "test_l2tp_ether_x"
	mtu                      = "1400"
	circuit_id               = "testl2tpetherx"
	allow_fast_path          = true
	use_l2_specific_sublayer = true
	disabled                 = false
}
`
}

func testAccInterfaceL2tpEtherConfigUpdated() string {
	return providerConfig + `
resource "routeros_interface_l2tp_ether" "test_l2tp_ether_x" {
	name                     = "test_l2tp_ether_x"
	connect_to               = "192.0.2.56"
	comment                  = "test_l2tp_ether_x updated"
	mtu                      = "1300"
	circuit_id               = "testl2tpetherx"
	allow_fast_path          = false
	use_l2_specific_sublayer = false
	disabled                 = true
}
`
}
