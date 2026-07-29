package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIpKidControlDevice = "routeros_ip_kid_control_device.test_kid_control_device_x"

func TestAccIpKidControlDeviceTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy: testCheckResourceDestroy("/ip/kid-control/device",
					"routeros_ip_kid_control_device"),
				Steps: []resource.TestStep{
					{
						Config: testAccIpKidControlDeviceConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpKidControlDevice),
							resource.TestCheckResourceAttr(testIpKidControlDevice, "name", "test_kid_control_device_x"),
							resource.TestCheckResourceAttr(testIpKidControlDevice, "mac_address", "00:11:22:33:44:70"),
							resource.TestCheckResourceAttr(testIpKidControlDevice, "user", "test_kid_control_device_user_x"),
							resource.TestCheckResourceAttr(testIpKidControlDevice, "disabled", "true"),
							resource.TestCheckResourceAttr(testIpKidControlDevice, "dynamic", "false"),
						),
					},
					{
						Config: testAccIpKidControlDeviceUpdatedConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpKidControlDevice),
							resource.TestCheckResourceAttr(testIpKidControlDevice, "name", "test_kid_control_device_x_upd"),
							resource.TestCheckResourceAttr(testIpKidControlDevice, "mac_address", "00:11:22:33:44:71"),
						),
					},
				},
			})

		})
	}
}

func testAccIpKidControlDeviceConfig() string {
	return providerConfig + `
resource "routeros_ip_kid_control" "test_kid_control_device_user_x" {
	name     = "test_kid_control_device_user_x"
	disabled = true
}

resource "routeros_ip_kid_control_device" "test_kid_control_device_x" {
	name        = "test_kid_control_device_x"
	mac_address = "00:11:22:33:44:70"
	user        = routeros_ip_kid_control.test_kid_control_device_user_x.name
	disabled    = true
}
`
}

func testAccIpKidControlDeviceUpdatedConfig() string {
	return providerConfig + `
resource "routeros_ip_kid_control" "test_kid_control_device_user_x" {
	name     = "test_kid_control_device_user_x"
	disabled = true
}

resource "routeros_ip_kid_control_device" "test_kid_control_device_x" {
	name        = "test_kid_control_device_x_upd"
	mac_address = "00:11:22:33:44:71"
	user        = routeros_ip_kid_control.test_kid_control_device_user_x.name
	disabled    = true
}
`
}
