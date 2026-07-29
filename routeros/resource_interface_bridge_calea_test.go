package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testInterfaceBridgeCalea = "routeros_interface_bridge_calea.test_bridge_calea_x"

// src_mac_address is not exercised: the schema's ValidationMacAddress rejects the masked MAC form that
// RouterOS requires for src-mac-address; dst_mac_address (ValidationMacAddressWithMask) is tested instead.

func TestAccInterfaceBridgeCaleaTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy: testCheckResourceDestroy("/interface/bridge/calea",
					"routeros_interface_bridge_calea"),
				Steps: []resource.TestStep{
					{
						Config: testAccInterfaceBridgeCaleaConfig,
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfaceBridgeCalea),
							resource.TestCheckResourceAttr(testInterfaceBridgeCalea, "chain", "forward"),
							resource.TestCheckResourceAttr(testInterfaceBridgeCalea, "action", "sniff"),
							resource.TestCheckResourceAttr(testInterfaceBridgeCalea, "sniff_target", "192.0.2.77"),
							resource.TestCheckResourceAttr(testInterfaceBridgeCalea, "sniff_target_port", "37008"),
							resource.TestCheckResourceAttr(testInterfaceBridgeCalea, "mac_protocol", "ip"),
							resource.TestCheckResourceAttr(testInterfaceBridgeCalea, "ip_protocol", "tcp"),
							resource.TestCheckResourceAttr(testInterfaceBridgeCalea, "dst_port", "443"),
							resource.TestCheckResourceAttr(testInterfaceBridgeCalea, "in_bridge", "bridge"),
							resource.TestCheckResourceAttr(testInterfaceBridgeCalea, "comment", "test_bridge_calea_x"),
							resource.TestCheckResourceAttr(testInterfaceBridgeCalea, "disabled", "true"),
							resource.TestCheckResourceAttr(testInterfaceBridgeCalea, "dynamic", "false"),
						),
					},
					{
						Config: testAccInterfaceBridgeCaleaConfigUpdated,
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfaceBridgeCalea),
							resource.TestCheckResourceAttr(testInterfaceBridgeCalea, "action", "sniff-pc"),
							resource.TestCheckResourceAttr(testInterfaceBridgeCalea, "sniff_id", "11"),
							resource.TestCheckResourceAttr(testInterfaceBridgeCalea, "sniff_target", "192.0.2.78"),
							resource.TestCheckResourceAttr(testInterfaceBridgeCalea, "dst_mac_address",
								"02:00:00:CA:1E:A0/FF:FF:FF:FF:FF:FF"),
							resource.TestCheckResourceAttr(testInterfaceBridgeCalea, "mac_protocol", "vlan"),
							resource.TestCheckResourceAttr(testInterfaceBridgeCalea, "vlan_encap", "ip"),
							resource.TestCheckResourceAttr(testInterfaceBridgeCalea, "vlan_id", "512"),
							resource.TestCheckResourceAttr(testInterfaceBridgeCalea, "vlan_priority", "5"),
							resource.TestCheckResourceAttr(testInterfaceBridgeCalea, "comment",
								"test_bridge_calea_x updated"),
						),
					},
				},
			})

		})
	}
}

// in_bridge references the pre-existing bridge fixture; the rule only names it.
var testAccInterfaceBridgeCaleaConfigDeps = providerConfig

var testAccInterfaceBridgeCaleaConfig = testAccInterfaceBridgeCaleaConfigDeps + `
resource "routeros_interface_bridge_calea" "test_bridge_calea_x" {
	chain             = "forward"
	action            = "sniff"
	sniff_target      = "192.0.2.77"
	sniff_target_port = 37008
	mac_protocol      = "ip"
	ip_protocol       = "tcp"
	dst_port          = "443"
	in_bridge         = "bridge"
	comment           = "test_bridge_calea_x"
	disabled          = true
}
`

var testAccInterfaceBridgeCaleaConfigUpdated = testAccInterfaceBridgeCaleaConfigDeps + `
resource "routeros_interface_bridge_calea" "test_bridge_calea_x" {
	chain             = "forward"
	action            = "sniff-pc"
	sniff_id          = 11
	sniff_target      = "192.0.2.78"
	sniff_target_port = 37008
	mac_protocol      = "vlan"
	vlan_encap        = "ip"
	dst_mac_address   = "02:00:00:CA:1E:A0/FF:FF:FF:FF:FF:FF"
	vlan_id           = 512
	vlan_priority     = 5
	in_bridge         = "bridge"
	comment           = "test_bridge_calea_x updated"
	disabled          = true
}
`
