package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIotModbusSecurityRules = "routeros_iot_modbus_security_rules.test_iot_modbus_sec_rule_x"

func TestAccIotModbusSecurityRulesTest_basic(t *testing.T) {
	testCheckMenu(t, "/iot/modbus/security-rules")
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy: testCheckResourceDestroy("/iot/modbus/security-rules",
					"routeros_iot_modbus_security_rules"),
				Steps: []resource.TestStep{
					{
						Config: testAccIotModbusSecurityRulesConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIotModbusSecurityRules),
							resource.TestCheckResourceAttr(testIotModbusSecurityRules, "ip_range", "198.51.100.0/24"),
							resource.TestCheckResourceAttr(testIotModbusSecurityRules, "disabled", "true"),
							resource.TestCheckResourceAttr(testIotModbusSecurityRules, "allowed_function_codes.#", "2"),
							resource.TestCheckResourceAttr(testIotModbusSecurityRules, "allowed_function_codes.0", "3"),
							resource.TestCheckResourceAttr(testIotModbusSecurityRules, "allowed_function_codes.1", "6"),
						),
					},
					{
						Config: testAccIotModbusSecurityRulesUpdatedConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIotModbusSecurityRules),
							resource.TestCheckResourceAttr(testIotModbusSecurityRules, "ip_range", "203.0.113.0/25"),
							resource.TestCheckResourceAttr(testIotModbusSecurityRules, "disabled", "false"),
							resource.TestCheckResourceAttr(testIotModbusSecurityRules, "allowed_function_codes.#", "3"),
							resource.TestCheckResourceAttr(testIotModbusSecurityRules, "allowed_function_codes.0", "1"),
							resource.TestCheckResourceAttr(testIotModbusSecurityRules, "allowed_function_codes.1", "3"),
							resource.TestCheckResourceAttr(testIotModbusSecurityRules, "allowed_function_codes.2", "16"),
						),
					},
				},
			})

		})
	}
}

func testAccIotModbusSecurityRulesConfig() string {
	return providerConfig + `

resource "routeros_iot_modbus_security_rules" "test_iot_modbus_sec_rule_x" {
	ip_range               = "198.51.100.0/24"
	allowed_function_codes = [3, 6]
	disabled               = true
}
`
}

func testAccIotModbusSecurityRulesUpdatedConfig() string {
	return providerConfig + `

resource "routeros_iot_modbus_security_rules" "test_iot_modbus_sec_rule_x" {
	ip_range               = "203.0.113.0/25"
	allowed_function_codes = [1, 3, 16]
	disabled               = false
}
`
}
