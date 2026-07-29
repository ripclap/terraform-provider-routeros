package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testInterfaceMeshAddress = "routeros_interface_mesh.test_mesh_x"

func TestAccInterfaceMeshTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/interface/mesh", "routeros_interface_mesh"),
				Steps: []resource.TestStep{
					{
						Config: testAccInterfaceMeshConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfaceMeshAddress),
							resource.TestCheckResourceAttr(testInterfaceMeshAddress, "name", "test_mesh_x"),
							resource.TestCheckResourceAttr(testInterfaceMeshAddress, "comment", "test_mesh_x"),
							resource.TestCheckResourceAttr(testInterfaceMeshAddress, "mtu", "1400"),
							resource.TestCheckResourceAttr(testInterfaceMeshAddress, "arp", "proxy-arp"),
							resource.TestCheckResourceAttr(testInterfaceMeshAddress, "arp_timeout", "1m"),
							resource.TestCheckResourceAttr(testInterfaceMeshAddress, "auto_mac", "false"),
							resource.TestCheckResourceAttr(testInterfaceMeshAddress, "admin_mac", "02:11:22:33:44:55"),
							resource.TestCheckResourceAttr(testInterfaceMeshAddress, "mac_address", "02:11:22:33:44:55"),
							resource.TestCheckResourceAttr(testInterfaceMeshAddress, "hwmp_default_hoplimit", "20"),
							resource.TestCheckResourceAttr(testInterfaceMeshAddress, "hwmp_prep_lifetime", "4m"),
							resource.TestCheckResourceAttr(testInterfaceMeshAddress, "hwmp_preq_destination_only", "false"),
							resource.TestCheckResourceAttr(testInterfaceMeshAddress, "hwmp_preq_reply_and_forward", "false"),
							resource.TestCheckResourceAttr(testInterfaceMeshAddress, "hwmp_preq_retries", "3"),
							resource.TestCheckResourceAttr(testInterfaceMeshAddress, "hwmp_preq_waiting_time", "3s"),
							resource.TestCheckResourceAttr(testInterfaceMeshAddress, "hwmp_rann_interval", "12s"),
							resource.TestCheckResourceAttr(testInterfaceMeshAddress, "hwmp_rann_lifetime", "25s"),
							resource.TestCheckResourceAttr(testInterfaceMeshAddress, "hwmp_rann_propagation_delay", "1"),
							resource.TestCheckResourceAttr(testInterfaceMeshAddress, "mesh_portal", "true"),
							resource.TestCheckResourceAttr(testInterfaceMeshAddress, "reoptimize_paths", "true"),
							resource.TestCheckResourceAttr(testInterfaceMeshAddress, "disabled", "false"),
						),
					},
					{
						Config: testAccInterfaceMeshConfigUpdated(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfaceMeshAddress),
							resource.TestCheckResourceAttr(testInterfaceMeshAddress, "comment", "test_mesh_x updated"),
							resource.TestCheckResourceAttr(testInterfaceMeshAddress, "mtu", "1500"),
							resource.TestCheckResourceAttr(testInterfaceMeshAddress, "arp", "enabled"),
							resource.TestCheckResourceAttr(testInterfaceMeshAddress, "auto_mac", "true"),
							resource.TestCheckResourceAttr(testInterfaceMeshAddress, "hwmp_default_hoplimit", "32"),
							resource.TestCheckResourceAttr(testInterfaceMeshAddress, "mesh_portal", "false"),
							resource.TestCheckResourceAttr(testInterfaceMeshAddress, "reoptimize_paths", "false"),
							resource.TestCheckResourceAttr(testInterfaceMeshAddress, "disabled", "true"),
						),
					},
				},
			})

		})
	}
}

func testAccInterfaceMeshConfig() string {
	return providerConfig + `
resource "routeros_interface_mesh" "test_mesh_x" {
	name                        = "test_mesh_x"
	comment                     = "test_mesh_x"
	mtu                         = 1400
	arp                         = "proxy-arp"
	arp_timeout                 = "1m"
	auto_mac                    = false
	admin_mac                   = "02:11:22:33:44:55"
	hwmp_default_hoplimit       = 20
	hwmp_prep_lifetime          = "4m"
	hwmp_preq_destination_only  = false
	hwmp_preq_reply_and_forward = false
	hwmp_preq_retries           = 3
	hwmp_preq_waiting_time      = "3s"
	hwmp_rann_interval          = "12s"
	hwmp_rann_lifetime          = "25s"
	hwmp_rann_propagation_delay = "1"
	mesh_portal                 = true
	reoptimize_paths            = true
	disabled                    = false
}
`
}

func testAccInterfaceMeshConfigUpdated() string {
	return providerConfig + `
resource "routeros_interface_mesh" "test_mesh_x" {
	name                        = "test_mesh_x"
	comment                     = "test_mesh_x updated"
	mtu                         = 1500
	arp                         = "enabled"
	arp_timeout                 = "auto"
	auto_mac                    = true
	hwmp_default_hoplimit       = 32
	hwmp_prep_lifetime          = "5m"
	hwmp_preq_destination_only  = true
	hwmp_preq_reply_and_forward = true
	hwmp_preq_retries           = 2
	hwmp_preq_waiting_time      = "4s"
	hwmp_rann_interval          = "10s"
	hwmp_rann_lifetime          = "22s"
	mesh_portal                 = false
	reoptimize_paths            = false
	disabled                    = true
}
`
}
