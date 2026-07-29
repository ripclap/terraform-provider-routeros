package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testResourceConsoleSettings = "routeros_console_settings.test_console_settings_x"

// `sanitize_names` is kept `false`: enabling it re-serializes every named object, slow enough to drop
// the transport ("Session closed"). A device-side property of the setting, not of the resource.

func TestAccConsoleSettingsTest_basic(t *testing.T) {
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
						Config: testAccConsoleSettingsConfig(6, false, false),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testResourceConsoleSettings),
							resource.TestCheckResourceAttr(testResourceConsoleSettings, "tab_width", "6"),
							resource.TestCheckResourceAttr(testResourceConsoleSettings, "sanitize_names", "false"),
							resource.TestCheckResourceAttr(testResourceConsoleSettings, "log_script_errors", "false"),
						),
					},
					{
						// Restore the RouterOS defaults.
						Config: testAccConsoleSettingsConfig(4, false, true),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testResourceConsoleSettings),
							resource.TestCheckResourceAttr(testResourceConsoleSettings, "tab_width", "4"),
							resource.TestCheckResourceAttr(testResourceConsoleSettings, "sanitize_names", "false"),
							resource.TestCheckResourceAttr(testResourceConsoleSettings, "log_script_errors", "true"),
						),
					},
				},
			})

		})
	}
}

func testAccConsoleSettingsConfig(tabWidth int, sanitizeNames, logScriptErrors bool) string {
	return providerConfig + fmt.Sprintf(`
resource "routeros_console_settings" "test_console_settings_x" {
	tab_width         = %v
	sanitize_names    = %v
	log_script_errors = %v
}
`, tabWidth, sanitizeNames, logScriptErrors)
}
