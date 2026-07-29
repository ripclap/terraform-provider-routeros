package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testSystemLoggingAction = "routeros_system_logging_action.action"

// testSystemLoggingRemoteFormatVersion is the release where RouterOS replaced the boolean bsd-syslog parameter
// with the remote-log-format enum; sending bsd-syslog to a newer device fails with "unknown parameter
// bsd-syslog", so the test picks the attribute matching the device under test.
const testSystemLoggingRemoteFormatVersion = "7.18"

func TestAccSystemLoggingActionTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			useRemoteLogFormat := testCheckMinVersion(t, testSystemLoggingRemoteFormatVersion)

			formatChecks := resource.TestCheckResourceAttr(testSystemLoggingAction, "bsd_syslog", "true")
			if useRemoteLogFormat {
				formatChecks = resource.TestCheckResourceAttr(testSystemLoggingAction, "remote_log_format", "syslog")
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				Steps: []resource.TestStep{
					{
						Config: testAccSystemLoggingActionConfig(useRemoteLogFormat),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testSystemLoggingAction),
							resource.TestCheckResourceAttr(testSystemLoggingAction, "name", "action1"),
							resource.TestCheckResourceAttr(testSystemLoggingAction, "default", "false"),
							resource.TestCheckResourceAttr(testSystemLoggingAction, "target", "remote"),
							resource.TestCheckResourceAttr(testSystemLoggingAction, "remote", "192.168.1.1"),
							formatChecks,
							resource.TestCheckResourceAttr(testSystemLoggingAction, "syslog_facility", "user"),
							resource.TestCheckResourceAttr(testSystemLoggingAction, "syslog_severity", "notice"),
							resource.TestCheckResourceAttr(testSystemLoggingAction, "syslog_time_format", "iso8601"),
						),
					},
				},
			})

		})
	}
}

func testAccSystemLoggingActionConfig(useRemoteLogFormat bool) string {
	format := "bsd_syslog         = true"
	if useRemoteLogFormat {
		format = `remote_log_format  = "syslog"`
	}

	return providerConfig + `
resource "routeros_system_logging_action" "action" {
	name               = "action1"
	target             = "remote"
	remote             = "192.168.1.1"
	` + format + `
	syslog_facility    = "user"
	syslog_severity    = "notice"
	syslog_time_format = "iso8601"
}
`
}
