package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testMplsSettingsTask = "routeros_mpls_settings.test_mpls_settings_x"

func TestAccMplsSettingsTest_basic(t *testing.T) {
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
						Config: testAccMplsSettingsConfig("1000-1048575", "false"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testMplsSettingsTask),
							resource.TestCheckResourceAttr(testMplsSettingsTask, "dynamic_label_range", "1000-1048575"),
							resource.TestCheckResourceAttr(testMplsSettingsTask, "propagate_ttl", "false"),
							resource.TestCheckResourceAttr(testMplsSettingsTask, "allow_fast_path", "true"),
						),
					},
					{
						// Restore the RouterOS defaults.
						Config: testAccMplsSettingsConfig("16-1048575", "true"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testMplsSettingsTask),
							resource.TestCheckResourceAttr(testMplsSettingsTask, "dynamic_label_range", "16-1048575"),
							resource.TestCheckResourceAttr(testMplsSettingsTask, "propagate_ttl", "true"),
						),
					},
				},
			})

		})
	}
}

func testAccMplsSettingsConfig(labelRange, propagateTtl string) string {
	return providerConfig + `

resource "routeros_mpls_settings" "test_mpls_settings_x" {
	allow_fast_path     = true
	dynamic_label_range = "` + labelRange + `"
	propagate_ttl       = ` + propagateTtl + `
}

`
}
