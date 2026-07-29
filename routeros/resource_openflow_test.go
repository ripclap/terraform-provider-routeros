package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testOpenflowAddress = "routeros_openflow.test_openflow_x"

// The `controllers` value grammar on RouterOS 7.23.2 is `protocol/address/port`, the console
// rejects the `protocol/address:port` form. `datapath_id` and `certificate` are always reported
// back by the device, so both are pinned in the configuration; `datapath_id` only accepts the
// `number/MAC` form.
func TestAccOpenflowTest_basic(t *testing.T) {
	testCheckMenu(t, "/openflow")
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/openflow", "routeros_openflow"),
				Steps: []resource.TestStep{
					{
						Config: testAccOpenflowConfig("default"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testOpenflowAddress),
							resource.TestCheckResourceAttr(testOpenflowAddress, "name", "test_openflow_x"),
							resource.TestCheckResourceAttr(testOpenflowAddress, "controllers.#", "2"),
							resource.TestCheckResourceAttr(testOpenflowAddress, "controllers.0", "tcp/192.0.2.10/6653"),
							resource.TestCheckResourceAttr(testOpenflowAddress, "controllers.1", "tcp/192.0.2.11/6653"),
							resource.TestCheckResourceAttr(testOpenflowAddress, "datapath_id", "8/02:00:00:00:00:08"),
							resource.TestCheckResourceAttr(testOpenflowAddress, "certificate", "none"),
							resource.TestCheckResourceAttr(testOpenflowAddress, "verify_peer", "none"),
							resource.TestCheckResourceAttr(testOpenflowAddress, "version", "default"),
							resource.TestCheckResourceAttr(testOpenflowAddress, "passive_port", "6654"),
							resource.TestCheckResourceAttr(testOpenflowAddress, "isolate_controllers", "true"),
							resource.TestCheckResourceAttr(testOpenflowAddress, "disabled", "true"),
						),
					},
					{
						Config: testAccOpenflowConfig("1.3"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testOpenflowAddress),
							resource.TestCheckResourceAttr(testOpenflowAddress, "version", "1.3"),
						),
					},
				},
			})

		})
	}
}

func testAccOpenflowConfig(version string) string {
	return providerConfig + `

resource "routeros_openflow" "test_openflow_x" {
	name                = "test_openflow_x"
	controllers         = ["tcp/192.0.2.10/6653", "tcp/192.0.2.11/6653"]
	datapath_id         = "8/02:00:00:00:00:08"
	certificate         = "none"
	verify_peer         = "none"
	version             = "` + version + `"
	passive_port        = "6654"
	isolate_controllers = true
	disabled            = true
	comment             = "test_openflow_x"
}

`
}
