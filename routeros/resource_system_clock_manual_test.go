package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testSystemClockManual = "routeros_system_clock_manual.test_system_clock_manual"

func TestAccSystemClockManualTest_basic(t *testing.T) {
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
						Config: testAccSystemClockManualConfig("+02:00", "+01:00",
							"2026-03-29 02:00:00", "2026-10-25 03:00:00"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testSystemClockManual),
							resource.TestCheckResourceAttr(testSystemClockManual, "time_zone", "+02:00"),
							resource.TestCheckResourceAttr(testSystemClockManual, "dst_delta", "+01:00"),
							resource.TestCheckResourceAttr(testSystemClockManual, "dst_start",
								"2026-03-29 02:00:00"),
							resource.TestCheckResourceAttr(testSystemClockManual, "dst_end",
								"2026-10-25 03:00:00"),
						),
					},
					{
						Config: testAccSystemClockManualConfig("+00:00", "+00:00",
							"1970-01-01 00:00:00", "1970-01-01 00:00:00"),
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr(testSystemClockManual, "time_zone", "+00:00"),
							resource.TestCheckResourceAttr(testSystemClockManual, "dst_delta", "+00:00"),
							resource.TestCheckResourceAttr(testSystemClockManual, "dst_start",
								"1970-01-01 00:00:00"),
							resource.TestCheckResourceAttr(testSystemClockManual, "dst_end",
								"1970-01-01 00:00:00"),
						),
					},
				},
			})

		})
	}
}

func testAccSystemClockManualConfig(tz, delta, start, end string) string {
	return fmt.Sprintf(`%v

resource "routeros_system_clock_manual" "test_system_clock_manual" {
	time_zone = "%v"
	dst_delta = "%v"
	dst_start = "%v"
	dst_end   = "%v"
}
`, providerConfig, tz, delta, start, end)
}
