package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testSystemResourceHardwareUsbSettings = "routeros_system_resource_hardware_usb_settings." +
	"test_system_resource_hardware_usb_settings"

func TestAccSystemResourceHardwareUsbSettingsTest_basic(t *testing.T) {
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
						Config: testAccSystemResourceHardwareUsbSettingsConfig("true"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testSystemResourceHardwareUsbSettings),
							resource.TestCheckResourceAttr(testSystemResourceHardwareUsbSettings,
								"authorization", "true"),
						),
					},
					{
						Config: testAccSystemResourceHardwareUsbSettingsConfig("false"),
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr(testSystemResourceHardwareUsbSettings,
								"authorization", "false"),
						),
					},
				},
			})

		})
	}
}

func testAccSystemResourceHardwareUsbSettingsConfig(authorization string) string {
	return fmt.Sprintf(`%v

resource "routeros_system_resource_hardware_usb_settings" "test_system_resource_hardware_usb_settings" {
	authorization = %v
}
`, providerConfig, authorization)
}
