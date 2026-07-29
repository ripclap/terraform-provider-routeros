package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testOpenflowPortAddress = "routeros_openflow_port.test_openflow_port_x"

func TestAccOpenflowPortTest_basic(t *testing.T) {
	testCheckMenu(t, "/openflow/port")
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/openflow/port", "routeros_openflow_port"),
				Steps: []resource.TestStep{
					{
						Config: testAccOpenflowPortConfig("11"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testOpenflowPortAddress),
							resource.TestCheckResourceAttr(testOpenflowPortAddress, "interface", "ether5"),
							resource.TestCheckResourceAttr(testOpenflowPortAddress, "switch", "test_openflow_sw_x"),
							resource.TestCheckResourceAttr(testOpenflowPortAddress, "port_id", "11"),
							resource.TestCheckResourceAttr(testOpenflowPortAddress, "disabled", "true"),
							resource.TestCheckResourceAttr(testOpenflowPortAddress, "dynamic", "false"),
							resource.TestCheckResourceAttr(testOpenflowPortAddress, "comment", "test_openflow_port_x"),
						),
					},
					{
						Config: testAccOpenflowPortConfig("12"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testOpenflowPortAddress),
							resource.TestCheckResourceAttr(testOpenflowPortAddress, "port_id", "12"),
						),
					},
				},
			})

		})
	}
}

func testAccOpenflowPortConfig(portId string) string {
	return providerConfig + `

resource "routeros_openflow" "test_openflow_sw_x" {
	name        = "test_openflow_sw_x"
	datapath_id = "9/02:00:00:00:00:09"
	certificate = "none"
	disabled    = true
	comment     = "test_openflow_sw_x"
}

resource "routeros_openflow_port" "test_openflow_port_x" {
	interface = "ether5"
	switch    = routeros_openflow.test_openflow_sw_x.name
	port_id   = ` + portId + `
	disabled  = true
	comment   = "test_openflow_port_x"
}

`
}
