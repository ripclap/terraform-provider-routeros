package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testResourceCertificateSettings = "routeros_certificate_settings.test_certificate_settings_x"

func TestAccCertificateSettingsTest_basic(t *testing.T) {
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
						Config: testAccCertificateSettingsConfig("all", "system", true),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testResourceCertificateSettings),
							resource.TestCheckResourceAttr(testResourceCertificateSettings, "builtin_trust_store", "all"),
							resource.TestCheckResourceAttr(testResourceCertificateSettings, "crl_store", "system"),
							resource.TestCheckResourceAttr(testResourceCertificateSettings, "crl_download", "true"),
							resource.TestCheckResourceAttr(testResourceCertificateSettings, "crl_use", "false"),
							resource.TestCheckResourceAttr(testResourceCertificateSettings, "current_defaults",
								"fetch,mqtt,email,netwatch,container,lora,wiliot,dns,www,reverse-proxy"),
						),
					},
					{
						// Restore the RouterOS defaults.
						Config: testAccCertificateSettingsConfig("default", "ram", false),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testResourceCertificateSettings),
							resource.TestCheckResourceAttr(testResourceCertificateSettings, "builtin_trust_store", "default"),
							resource.TestCheckResourceAttr(testResourceCertificateSettings, "crl_store", "ram"),
							resource.TestCheckResourceAttr(testResourceCertificateSettings, "crl_download", "false"),
						),
					},
				},
			})

		})
	}
}

func testAccCertificateSettingsConfig(trustStore, crlStore string, crlDownload bool) string {
	return providerConfig + fmt.Sprintf(`
resource "routeros_certificate_settings" "test_certificate_settings_x" {
	builtin_trust_store = "%v"
	crl_download        = %v
	crl_store           = "%v"
	crl_use             = false
	current_defaults    = "fetch,mqtt,email,netwatch,container,lora,wiliot,dns,www,reverse-proxy"
}
`, trustStore, crlDownload, crlStore)
}
