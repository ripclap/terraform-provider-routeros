package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testSystemConsole = "routeros_system_console.test_system_console"

func TestAccSystemConsoleTest_basic(t *testing.T) {
	testCheckSpareSerialPort(t)
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/system/console", "routeros_system_console"),
				Steps: []resource.TestStep{
					{
						Config: testAccSystemConsoleConfig("vt100", "true"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testSystemConsole),
							resource.TestCheckResourceAttr(testSystemConsole, "port", "serial0"),
							resource.TestCheckResourceAttr(testSystemConsole, "term", "vt100"),
							resource.TestCheckResourceAttr(testSystemConsole, "channel", "0"),
							resource.TestCheckResourceAttr(testSystemConsole, "disabled", "true"),
							resource.TestCheckResourceAttr(testSystemConsole, "default", "false"),
						),
					},
					{
						Config: testAccSystemConsoleConfig("vt102", "false"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testSystemConsole),
							resource.TestCheckResourceAttr(testSystemConsole, "term", "vt102"),
							resource.TestCheckResourceAttr(testSystemConsole, "disabled", "false"),
						),
					},
				},
			})

		})
	}
}

func testAccSystemConsoleConfig(term, disabled string) string {
	return fmt.Sprintf(`%v

resource "routeros_system_console" "test_system_console" {
	port     = "serial0"
	term     = "%v"
	channel  = 0
	disabled = %v
}
`, providerConfig, term, disabled)
}
