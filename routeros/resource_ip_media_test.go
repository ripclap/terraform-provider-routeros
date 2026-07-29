package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIpMedia = "routeros_ip_media.test_media_x"

func TestAccIpMediaTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ip/media", "routeros_ip_media"),
				Steps: []resource.TestStep{
					{
						Config: testAccIpMediaConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpMedia),
							resource.TestCheckResourceAttr(testIpMedia, "path", "test_media_x_path"),
							resource.TestCheckResourceAttr(testIpMedia, "interface", "ether3"),
							resource.TestCheckResourceAttr(testIpMedia, "friendly_name", "test_media_x"),
							resource.TestCheckResourceAttr(testIpMedia, "allowed_ip", "192.0.2.70"),
							resource.TestCheckResourceAttr(testIpMedia, "disabled", "true"),
						),
					},
					{
						Config: testAccIpMediaUpdatedConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpMedia),
							resource.TestCheckResourceAttr(testIpMedia, "path", "test_media_x_path2"),
							resource.TestCheckResourceAttr(testIpMedia, "interface", "ether4"),
							resource.TestCheckResourceAttr(testIpMedia, "friendly_name", "test_media_x_upd"),
							resource.TestCheckResourceAttr(testIpMedia, "allowed_ip", "192.0.2.71"),
							resource.TestCheckResourceAttr(testIpMedia, "allowed_hostname", "test-media-x.example.com"),
						),
					},
				},
			})

		})
	}
}

func testAccIpMediaConfig() string {
	return providerConfig + `
resource "routeros_ip_media" "test_media_x" {
	path          = "test_media_x_path"
	interface     = "ether3"
	friendly_name = "test_media_x"
	allowed_ip    = "192.0.2.70"
	disabled      = true
}
`
}

func testAccIpMediaUpdatedConfig() string {
	return providerConfig + `
resource "routeros_ip_media" "test_media_x" {
	path             = "test_media_x_path2"
	interface        = "ether4"
	friendly_name    = "test_media_x_upd"
	allowed_ip       = "192.0.2.71"
	allowed_hostname = "test-media-x.example.com"
	disabled         = true
}
`
}
