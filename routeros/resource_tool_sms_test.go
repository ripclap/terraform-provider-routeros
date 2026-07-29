package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testResourceToolSms = "routeros_tool_sms.test_tool_sms"

// /tool/sms is a settings object that cannot be created or destroyed, so this test uses the singleton
// pattern (no CheckDestroy); every property is accepted without a modem attached (`port = none`).
func TestAccToolSmsTest_basic(t *testing.T) {
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
						Config: testAccToolSmsConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testResourceToolSms),
							resource.TestCheckResourceAttr(testResourceToolSms, "allowed_number", "+15550100"),
							resource.TestCheckResourceAttr(testResourceToolSms, "channel", "1"),
							resource.TestCheckResourceAttr(testResourceToolSms, "polling", "true"),
							resource.TestCheckResourceAttr(testResourceToolSms, "port", "none"),
							resource.TestCheckResourceAttr(testResourceToolSms, "receive_enabled", "false"),
							resource.TestCheckResourceAttr(testResourceToolSms, "remove_sent_sms_after_send", "true"),
							resource.TestCheckResourceAttr(testResourceToolSms, "secret", "test_tool_sms_secret"),
							resource.TestCheckResourceAttr(testResourceToolSms, "sim_pin", "1234"),
							resource.TestCheckResourceAttr(testResourceToolSms, "sms_storage", "modem"),
						),
					},
					{
						Config: testAccToolSmsDefaultsConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testResourceToolSms),
							resource.TestCheckResourceAttr(testResourceToolSms, "allowed_number", ""),
							resource.TestCheckResourceAttr(testResourceToolSms, "channel", "0"),
							resource.TestCheckResourceAttr(testResourceToolSms, "polling", "false"),
							resource.TestCheckResourceAttr(testResourceToolSms, "remove_sent_sms_after_send", "false"),
							resource.TestCheckResourceAttr(testResourceToolSms, "sms_storage", "sim"),
						),
					},
				},
			})
		})
	}
}

func testAccToolSmsConfig() string {
	return providerConfig + `
resource "routeros_tool_sms" "test_tool_sms" {
	allowed_number             = "+15550100"
	channel                    = 1
	polling                    = true
	port                       = "none"
	receive_enabled            = false
	remove_sent_sms_after_send = true
	secret                     = "test_tool_sms_secret"
	sim_pin                    = "1234"
	sms_storage                = "modem"
}
`
}

func testAccToolSmsDefaultsConfig() string {
	return providerConfig + `
resource "routeros_tool_sms" "test_tool_sms" {
	allowed_number             = ""
	channel                    = 0
	polling                    = false
	port                       = "none"
	receive_enabled            = false
	remove_sent_sms_after_send = false
	secret                     = ""
	sim_pin                    = ""
	sms_storage                = "sim"
}
`
}
