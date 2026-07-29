package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testInterfaceWifiNetwork = "routeros_interface_wifi_network.test_wifi_network_x"

func TestAccInterfaceWifiNetworkTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy: testCheckResourceDestroy("/interface/wifi/network",
					"routeros_interface_wifi_network"),
				Steps: []resource.TestStep{
					{
						Config: testAccInterfaceWifiNetworkConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfaceWifiNetwork),
							resource.TestCheckResourceAttr(testInterfaceWifiNetwork, "ssid", "test_wifi_network_x"),
							resource.TestCheckResourceAttr(testInterfaceWifiNetwork, "comment", "test_wifi_network_x"),
							resource.TestCheckResourceAttr(testInterfaceWifiNetwork, "disabled", "true"),
						),
					},
					{
						Config: testAccInterfaceWifiNetworkUpdatedConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfaceWifiNetwork),
							resource.TestCheckResourceAttr(testInterfaceWifiNetwork, "ssid", "test_wifi_network_x"),
							resource.TestCheckResourceAttr(testInterfaceWifiNetwork, "comment", "updated"),
							resource.TestCheckResourceAttr(testInterfaceWifiNetwork, "beacon_interval", "100ms"),
							resource.TestCheckResourceAttr(testInterfaceWifiNetwork, "dtim_period", "2"),
							resource.TestCheckResourceAttr(testInterfaceWifiNetwork, "hide_ssid", "true"),
							resource.TestCheckResourceAttr(testInterfaceWifiNetwork, "hw_protection_mode", "rts-cts"),
							resource.TestCheckResourceAttr(testInterfaceWifiNetwork, "labels", "test_wifi_network_lbl_x"),
							resource.TestCheckResourceAttr(testInterfaceWifiNetwork, "max_clients", "32"),
							resource.TestCheckResourceAttr(testInterfaceWifiNetwork, "mode", "ap"),
							resource.TestCheckResourceAttr(testInterfaceWifiNetwork, "multicast_enhance", "enabled"),
							resource.TestCheckResourceAttr(testInterfaceWifiNetwork, "qos_classifier", "priority"),
							resource.TestCheckResourceAttr(testInterfaceWifiNetwork, "station_roaming", "true"),
							resource.TestCheckResourceAttr(testInterfaceWifiNetwork, "datapath.client_isolation", "true"),
							resource.TestCheckResourceAttr(testInterfaceWifiNetwork, "datapath.vlan_id", "101"),
							resource.TestCheckResourceAttr(testInterfaceWifiNetwork, "security.authentication_types", "wpa2-psk"),
							resource.TestCheckResourceAttr(testInterfaceWifiNetwork, "security.ft", "true"),
							resource.TestCheckResourceAttr(testInterfaceWifiNetwork, "security.passphrase", "TestWifiNetX123"),
						),
					},
				},
			})

		})
	}
}

func testAccInterfaceWifiNetworkConfig() string {
	return providerConfig + `

resource "routeros_interface_wifi_network" "test_wifi_network_x" {
	ssid     = "test_wifi_network_x"
	comment  = "test_wifi_network_x"
	disabled = true
}
`
}

func testAccInterfaceWifiNetworkUpdatedConfig() string {
	return providerConfig + `

resource "routeros_interface_wifi_network" "test_wifi_network_x" {
	ssid               = "test_wifi_network_x"
	comment            = "updated"
	disabled           = true
	beacon_interval    = "100ms"
	dtim_period        = 2
	hide_ssid          = true
	hw_protection_mode = "rts-cts"
	labels             = "test_wifi_network_lbl_x"
	max_clients        = 32
	mode               = "ap"
	multicast_enhance  = "enabled"
	qos_classifier     = "priority"
	station_roaming    = true

	datapath = {
		client_isolation = "true"
		vlan_id          = "101"
	}

	security = {
		authentication_types = "wpa2-psk"
		ft                   = "true"
		passphrase           = "TestWifiNetX123"
	}
}
`
}
