package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIotModbus = "routeros_iot_modbus.test_iot_modbus_x"

func TestAccIotModbusTest_basic(t *testing.T) {
	testCheckMenu(t, "/iot/modbus")
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
						Config: testAccIotModbusConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIotModbus),
							resource.TestCheckResourceAttr(testIotModbus, "disabled", "true"),
							resource.TestCheckResourceAttr(testIotModbus, "disable_security_rules", "false"),
							resource.TestCheckResourceAttr(testIotModbus, "interframe_gap", "5"),
							resource.TestCheckResourceAttr(testIotModbus, "rx_switch_offset", "10"),
							resource.TestCheckResourceAttr(testIotModbus, "tcp_port", "1502"),
							resource.TestCheckResourceAttr(testIotModbus, "timeout", "900"),
							resource.TestCheckResourceAttrSet(testIotModbus, "hardware_port"),
						),
					},
					{
						Config: testAccIotModbusRestoreConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIotModbus),
							resource.TestCheckResourceAttr(testIotModbus, "disabled", "true"),
							resource.TestCheckResourceAttr(testIotModbus, "disable_security_rules", "true"),
							resource.TestCheckResourceAttr(testIotModbus, "interframe_gap", "0"),
							resource.TestCheckResourceAttr(testIotModbus, "rx_switch_offset", "0"),
							resource.TestCheckResourceAttr(testIotModbus, "tcp_port", "502"),
							resource.TestCheckResourceAttr(testIotModbus, "timeout", "1000"),
						),
					},
				},
			})

		})
	}
}

func testAccIotModbusConfig() string {
	return providerConfig + `

resource "routeros_iot_modbus" "test_iot_modbus_x" {
	disabled               = true
	disable_security_rules = false
	interframe_gap         = 5
	rx_switch_offset       = 10
	tcp_port               = 1502
	timeout                = 900
}
`
}

func testAccIotModbusRestoreConfig() string {
	return providerConfig + `

resource "routeros_iot_modbus" "test_iot_modbus_x" {
	disabled               = true
	disable_security_rules = true
	interframe_gap         = 0
	rx_switch_offset       = 0
	tcp_port               = 502
	timeout                = 1000
}
`
}
