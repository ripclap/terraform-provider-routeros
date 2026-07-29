package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testResourceCertificateCrl = "routeros_certificate_crl.test_certificate_crl_x"

func TestAccCertificateCrlTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/certificate/crl", "routeros_certificate_crl"),
				Steps: []resource.TestStep{
					{
						Config: testAccCertificateCrlConfig("http://192.0.2.31/test_certificate_crl_x.crl"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testResourceCertificateCrl),
							resource.TestCheckResourceAttr(testResourceCertificateCrl, "url",
								"http://192.0.2.31/test_certificate_crl_x.crl"),
							resource.TestCheckResourceAttr(testResourceCertificateCrl, "dynamic", "false"),
							resource.TestCheckResourceAttr(testResourceCertificateCrl, "trust_store", "all"),
						),
					},
					{
						Config: testAccCertificateCrlConfig("http://192.0.2.32/test_certificate_crl_x2.crl"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testResourceCertificateCrl),
							resource.TestCheckResourceAttr(testResourceCertificateCrl, "url",
								"http://192.0.2.32/test_certificate_crl_x2.crl"),
						),
					},
				},
			})

		})
	}
}

func testAccCertificateCrlConfig(url string) string {
	return providerConfig + fmt.Sprintf(`
resource "routeros_certificate_crl" "test_certificate_crl_x" {
	url = "%v"
}
`, url)
}
