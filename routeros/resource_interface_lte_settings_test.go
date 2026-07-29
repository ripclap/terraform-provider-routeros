package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testInterfaceLteSettingsAddress = "routeros_interface_lte_settings.test_lte_settings_x"

func TestAccInterfaceLteSettingsTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				Steps: []resource.TestStep{
					{
						Config: testAccInterfaceLteSettingsConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfaceLteSettingsAddress),
							resource.TestCheckResourceAttr(testInterfaceLteSettingsAddress, "mode", "user"),
							resource.TestCheckResourceAttr(testInterfaceLteSettingsAddress, "esim_channel", "at"),
							resource.TestCheckResourceAttr(testInterfaceLteSettingsAddress, "firmware_path", "test_lte_settings_x_fw"),
							resource.TestCheckResourceAttr(testInterfaceLteSettingsAddress, "link_recovery_timer", "180"),
						),
					},
					{
						Config: testAccInterfaceLteSettingsConfigRestore(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfaceLteSettingsAddress),
							resource.TestCheckResourceAttr(testInterfaceLteSettingsAddress, "mode", "auto"),
							resource.TestCheckResourceAttr(testInterfaceLteSettingsAddress, "esim_channel", "auto"),
							resource.TestCheckResourceAttr(testInterfaceLteSettingsAddress, "firmware_path", "firmware"),
							resource.TestCheckResourceAttr(testInterfaceLteSettingsAddress, "link_recovery_timer", "120"),
						),
					},
				},
			})

		})
	}
}

func testAccInterfaceLteSettingsConfig() string {
	return providerConfig + `
resource "routeros_interface_lte_settings" "test_lte_settings_x" {
	mode                = "user"
	esim_channel        = "at"
	firmware_path       = "test_lte_settings_x_fw"
	link_recovery_timer = "180"
}
`
}

// testAccInterfaceLteSettingsConfigRestore returns the LTE settings to their factory values.
func testAccInterfaceLteSettingsConfigRestore() string {
	return providerConfig + `
resource "routeros_interface_lte_settings" "test_lte_settings_x" {
	mode                = "auto"
	esim_channel        = "auto"
	firmware_path       = "firmware"
	link_recovery_timer = "120"
}
`
}
