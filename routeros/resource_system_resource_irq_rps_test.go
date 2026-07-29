package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testResourceSystemResourceIrqRps = "routeros_system_resource_irq_rps.test_system_resource_irq_rps"

// RPS entries cannot be added or removed, so this test uses the singleton pattern without CheckDestroy.
// ether1 is the only interface with an RPS entry on the test VM; toggling the flag does not flap the link.
func TestAccSystemResourceIrqRpsTest_basic(t *testing.T) {
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
						Config: testAccSystemResourceIrqRpsConfig(true),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testResourceSystemResourceIrqRps),
							resource.TestCheckResourceAttr(testResourceSystemResourceIrqRps, "name", "ether1"),
							resource.TestCheckResourceAttr(testResourceSystemResourceIrqRps, "disabled", "true"),
						),
					},
					{
						Config: testAccSystemResourceIrqRpsConfig(false),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testResourceSystemResourceIrqRps),
							resource.TestCheckResourceAttr(testResourceSystemResourceIrqRps, "disabled", "false"),
						),
					},
				},
			})
		})
	}
}

func testAccSystemResourceIrqRpsConfig(disabled bool) string {
	state := "false"
	if disabled {
		state = "true"
	}
	return providerConfig + `
resource "routeros_system_resource_irq_rps" "test_system_resource_irq_rps" {
	name     = "ether1"
	disabled = ` + state + `
}
`
}
