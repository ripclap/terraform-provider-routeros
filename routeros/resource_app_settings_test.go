package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testResourceAppSettings = "routeros_app_settings.test_app_settings_x"

func TestAccAppSettingsTest_basic(t *testing.T) {
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
						Config: testAccAppSettingsConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testResourceAppSettings),
							resource.TestCheckResourceAttr(testResourceAppSettings, "auto_update", "false"),
							resource.TestCheckResourceAttr(testResourceAppSettings, "show_in_webfig", "true"),
							resource.TestCheckResourceAttr(testResourceAppSettings, "lan_bridge", "bridge"),
							resource.TestCheckResourceAttr(testResourceAppSettings, "router_ip", "192.0.2.1"),
							resource.TestCheckResourceAttr(testResourceAppSettings, "download_path",
								"test_app_settings_x/downloads"),
							resource.TestCheckResourceAttr(testResourceAppSettings, "media_path",
								"test_app_settings_x/media"),
							resource.TestCheckResourceAttr(testResourceAppSettings, "registry_mirrors",
								"https://mirror.test-app-settings-x.invalid"),
							resource.TestCheckResourceAttr(testResourceAppSettings, "app_store_urls",
								"https://store.test-app-settings-x.invalid/apps.yaml"),
						),
					},
					{
						Config: testAccAppSettingsConfig(),
					},
				},
			})

		})
	}
}

func testAccAppSettingsConfig() string {
	return providerConfig + `
resource "routeros_app_settings" "test_app_settings_x" {
	app_store_urls   = "https://store.test-app-settings-x.invalid/apps.yaml"
	auto_update      = false
	download_path    = "test_app_settings_x/downloads"
	lan_bridge       = "bridge"
	media_path       = "test_app_settings_x/media"
	registry_mirrors = "https://mirror.test-app-settings-x.invalid"
	router_ip        = "192.0.2.1"
	show_in_webfig   = true
}
`
}
