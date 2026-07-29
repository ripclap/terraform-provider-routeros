package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIpKidControl = "routeros_ip_kid_control.test_kid_control_x"

// RouterOS normalizes the time-of-day ranges of this menu (`08:00:00-20:00:00` is reported back as
// `8h-20h`), so the configuration uses the very format the device reports.
func TestAccIpKidControlTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ip/kid-control", "routeros_ip_kid_control"),
				Steps: []resource.TestStep{
					{
						Config: testAccIpKidControlConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpKidControl),
							resource.TestCheckResourceAttr(testIpKidControl, "name", "test_kid_control_x"),
							resource.TestCheckResourceAttr(testIpKidControl, "mon", "8h-20h"),
							resource.TestCheckResourceAttr(testIpKidControl, "tue", "8h-20h"),
							resource.TestCheckResourceAttr(testIpKidControl, "wed", "8h-20h"),
							resource.TestCheckResourceAttr(testIpKidControl, "rate_limit", "3M"),
							resource.TestCheckResourceAttr(testIpKidControl, "tur_sat", "9h-11h"),
							resource.TestCheckResourceAttr(testIpKidControl, "disabled", "true"),
						),
					},
					{
						Config: testAccIpKidControlUpdatedConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpKidControl),
							resource.TestCheckResourceAttr(testIpKidControl, "mon", "9h-18h"),
							resource.TestCheckResourceAttr(testIpKidControl, "rate_limit", "10M"),
							resource.TestCheckResourceAttr(testIpKidControl, "tur_sat", "10h-12h"),
							resource.TestCheckResourceAttr(testIpKidControl, "tur_sun", "10h-12h"),
						),
					},
				},
			})

		})
	}
}

func testAccIpKidControlConfig() string {
	return providerConfig + `
resource "routeros_ip_kid_control" "test_kid_control_x" {
	name       = "test_kid_control_x"
	mon        = "8h-20h"
	tue        = "8h-20h"
	wed        = "8h-20h"
	rate_limit = "3M"
	tur_sat    = "9h-11h"
	disabled   = true
}
`
}

func testAccIpKidControlUpdatedConfig() string {
	return providerConfig + `
resource "routeros_ip_kid_control" "test_kid_control_x" {
	name       = "test_kid_control_x"
	mon        = "9h-18h"
	tue        = "8h-20h"
	wed        = "8h-20h"
	rate_limit = "10M"
	tur_sat    = "10h-12h"
	tur_sun    = "10h-12h"
	disabled   = true
}
`
}
