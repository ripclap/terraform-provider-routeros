package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testPortAddress = "routeros_port.test_port_serial_x"

// Serial ports are created by the hardware: the resource adopts the existing `/port` entry and
// destroying it only drops the entry from the Terraform state, so there is nothing to check with
// CheckDestroy. The last step restores the RouterOS defaults of the port.
func TestAccPortTest_basic(t *testing.T) {
	testCheckSpareSerialPort(t)
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
						Config: testAccPortConfig("115200", "7", "2", "odd"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testPortAddress),
							resource.TestCheckResourceAttr(testPortAddress, "name", "serial0"),
							resource.TestCheckResourceAttr(testPortAddress, "baud_rate", "115200"),
							resource.TestCheckResourceAttr(testPortAddress, "data_bits", "7"),
							resource.TestCheckResourceAttr(testPortAddress, "stop_bits", "2"),
							resource.TestCheckResourceAttr(testPortAddress, "parity", "odd"),
							resource.TestCheckResourceAttr(testPortAddress, "flow_control", "none"),
						),
					},
					{
						// Restore the RouterOS defaults of the serial port.
						Config: testAccPortConfig("9600", "8", "1", "none"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testPortAddress),
							resource.TestCheckResourceAttr(testPortAddress, "baud_rate", "9600"),
							resource.TestCheckResourceAttr(testPortAddress, "data_bits", "8"),
							resource.TestCheckResourceAttr(testPortAddress, "stop_bits", "1"),
							resource.TestCheckResourceAttr(testPortAddress, "parity", "none"),
						),
					},
				},
			})

		})
	}
}

func testAccPortConfig(baudRate, dataBits, stopBits, parity string) string {
	return providerConfig + `

resource "routeros_port" "test_port_serial_x" {
	name         = "serial0"
	baud_rate    = "` + baudRate + `"
	data_bits    = ` + dataBits + `
	stop_bits    = ` + stopBits + `
	parity       = "` + parity + `"
	flow_control = "none"
}

`
}
