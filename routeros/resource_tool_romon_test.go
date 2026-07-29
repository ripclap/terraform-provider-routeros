package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testResourceToolRomon = "routeros_tool_romon.test_tool_romon"

// /tool/romon is a settings object that cannot be created or destroyed, so this test uses the
// singleton pattern (no CheckDestroy).
func TestAccToolRomonTest_basic(t *testing.T) {
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
						Config: testAccToolRomonConfig(true, "02:12:00:00:00:12", "test_tool_romon_secret"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testResourceToolRomon),
							resource.TestCheckResourceAttr(testResourceToolRomon, "enabled", "true"),
							resource.TestCheckResourceAttr(testResourceToolRomon, "romon_id", "02:12:00:00:00:12"),
							resource.TestCheckResourceAttr(testResourceToolRomon, "secrets", "test_tool_romon_secret"),
						),
					},
					{
						Config: testAccToolRomonConfig(false, "00:00:00:00:00:00", ""),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testResourceToolRomon),
							resource.TestCheckResourceAttr(testResourceToolRomon, "enabled", "false"),
							resource.TestCheckResourceAttr(testResourceToolRomon, "romon_id", "00:00:00:00:00:00"),
						),
					},
				},
			})
		})
	}
}

func testAccToolRomonConfig(enabled bool, romonId, secrets string) string {
	state := "false"
	if enabled {
		state = "true"
	}
	return providerConfig + `
resource "routeros_tool_romon" "test_tool_romon" {
	enabled  = ` + state + `
	romon_id = "` + romonId + `"
	secrets  = "` + secrets + `"
}
`
}
