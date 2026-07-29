package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testInterfaceWifiNetworkRadio = "routeros_interface_wifi_network_radio.test_wifi_network_radio_x"

func TestAccInterfaceWifiNetworkRadioTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy: testCheckResourceDestroy("/interface/wifi/network/radio",
					"routeros_interface_wifi_network_radio"),
				Steps: []resource.TestStep{
					{
						Config: testAccInterfaceWifiNetworkRadioConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfaceWifiNetworkRadio),
							resource.TestCheckResourceAttr(testInterfaceWifiNetworkRadio, "comment", "test_wifi_network_radio_x"),
							resource.TestCheckResourceAttr(testInterfaceWifiNetworkRadio, "disabled", "true"),
							resource.TestCheckResourceAttr(testInterfaceWifiNetworkRadio, "labels", "test_wnr_lbl_x"),
						),
					},
					{
						Config: testAccInterfaceWifiNetworkRadioUpdatedConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfaceWifiNetworkRadio),
							resource.TestCheckResourceAttr(testInterfaceWifiNetworkRadio, "comment", "updated"),
							resource.TestCheckResourceAttr(testInterfaceWifiNetworkRadio, "labels", "test_wnr_lbl_x"),
							resource.TestCheckResourceAttr(testInterfaceWifiNetworkRadio, "extra_labels", "test_wnr_extra_x"),
							resource.TestCheckResourceAttr(testInterfaceWifiNetworkRadio, "channel.band", "5ghz-ax"),
							resource.TestCheckResourceAttr(testInterfaceWifiNetworkRadio, "channel.frequency", "5180"),
							resource.TestCheckResourceAttr(testInterfaceWifiNetworkRadio, "channel.reselect_interval", "1h"),
							resource.TestCheckResourceAttr(testInterfaceWifiNetworkRadio, "channel.skip_dfs_channels", "all"),
							resource.TestCheckResourceAttr(testInterfaceWifiNetworkRadio, "channel.width", "20/40/80mhz"),
							resource.TestCheckResourceAttr(testInterfaceWifiNetworkRadio, "configuration.antenna_gain", "3"),
							resource.TestCheckResourceAttr(testInterfaceWifiNetworkRadio, "configuration.country", "Latvia"),
							resource.TestCheckResourceAttr(testInterfaceWifiNetworkRadio, "configuration.installation", "indoor"),
							resource.TestCheckResourceAttr(testInterfaceWifiNetworkRadio, "configuration.tx_power", "17"),
							resource.TestCheckResourceAttr(testInterfaceWifiNetworkRadio, "security.authentication_types", "wpa2-psk"),
						),
					},
				},
			})

		})
	}
}

func testAccInterfaceWifiNetworkRadioConfig() string {
	return providerConfig + `

resource "routeros_interface_wifi_network_radio" "test_wifi_network_radio_x" {
	comment  = "test_wifi_network_radio_x"
	disabled = true
	labels   = "test_wnr_lbl_x"
}
`
}

func testAccInterfaceWifiNetworkRadioUpdatedConfig() string {
	return providerConfig + `

resource "routeros_interface_wifi_network_radio" "test_wifi_network_radio_x" {
	comment      = "updated"
	disabled     = true
	labels       = "test_wnr_lbl_x"
	extra_labels = "test_wnr_extra_x"

	channel = {
		band              = "5ghz-ax"
		frequency         = "5180"
		reselect_interval = "1h"
		skip_dfs_channels = "all"
		width             = "20/40/80mhz"
	}

	configuration = {
		antenna_gain = "3"
		country      = "Latvia"
		installation = "indoor"
		tx_power     = "17"
	}

	security = {
		authentication_types = "wpa2-psk"
	}
}
`
}
