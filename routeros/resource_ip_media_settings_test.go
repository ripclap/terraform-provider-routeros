package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIpMediaSettings = "routeros_ip_media_settings.test_media_settings_x"

// `/ip/media/settings` is a settings singleton: it cannot be created or destroyed, only modified,
// so the test follows the "settings" pattern and carries no CheckDestroy.
func TestAccIpMediaSettingsTest_basic(t *testing.T) {
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
						Config: testAccIpMediaSettingsConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpMediaSettings),
							resource.TestCheckResourceAttr(testIpMediaSettings, "thumbnails",
								"test_media_settings_x_thumbs"),
						),
					},
					{
						// Restores the RouterOS default (an empty location).
						Config: testAccIpMediaSettingsRestoreConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpMediaSettings),
							resource.TestCheckResourceAttr(testIpMediaSettings, "thumbnails", ""),
						),
					},
				},
			})

		})
	}
}

func testAccIpMediaSettingsConfig() string {
	return providerConfig + `
resource "routeros_ip_media_settings" "test_media_settings_x" {
	thumbnails = "test_media_settings_x_thumbs"
}
`
}

func testAccIpMediaSettingsRestoreConfig() string {
	return providerConfig + `
resource "routeros_ip_media_settings" "test_media_settings_x" {
	thumbnails = ""
}
`
}
