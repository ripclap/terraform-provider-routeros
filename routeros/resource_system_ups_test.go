package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testResourceSystemUps = "routeros_system_ups.test_system_ups_x"

// A /system/ups entry can be created without a UPS attached; RouterOS just flags it INVALID, so the
// settable properties are still testable on the VM.
func TestAccSystemUpsTest_basic(t *testing.T) {
	testCheckMenu(t, "/system/ups")
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/system/ups", "routeros_system_ups"),
				Steps: []resource.TestStep{
					{
						Config: testAccSystemUpsConfig("immediate", "never"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testResourceSystemUps),
							resource.TestCheckResourceAttr(testResourceSystemUps, "name", "test_system_ups_x"),
							resource.TestCheckResourceAttr(testResourceSystemUps, "port", "serial0"),
							resource.TestCheckResourceAttr(testResourceSystemUps, "alarm_setting", "immediate"),
							resource.TestCheckResourceAttr(testResourceSystemUps, "min_runtime", "never"),
							resource.TestCheckResourceAttr(testResourceSystemUps, "offline_time", "1m"),
							resource.TestCheckResourceAttr(testResourceSystemUps, "check_capabilities", "false"),
							resource.TestCheckResourceAttr(testResourceSystemUps, "comment", "test_system_ups_x"),
							resource.TestCheckResourceAttr(testResourceSystemUps, "disabled", "false"),
						),
					},
					{
						Config: testAccSystemUpsConfig("low-battery", "5m"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testResourceSystemUps),
							resource.TestCheckResourceAttr(testResourceSystemUps, "alarm_setting", "low-battery"),
							resource.TestCheckResourceAttr(testResourceSystemUps, "min_runtime", "5m"),
						),
					},
				},
			})
		})
	}
}

func testAccSystemUpsConfig(alarmSetting, minRuntime string) string {
	return providerConfig + `
resource "routeros_system_ups" "test_system_ups_x" {
	name               = "test_system_ups_x"
	comment            = "test_system_ups_x"
	port               = "serial0"
	alarm_setting      = "` + alarmSetting + `"
	check_capabilities = false
	min_runtime        = "` + minRuntime + `"
	offline_time       = "1m"
}
`
}
