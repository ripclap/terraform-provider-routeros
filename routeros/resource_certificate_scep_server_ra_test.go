package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testResourceCertificateScepServerRa = "routeros_certificate_scep_server_ra.test_certificate_scep_ra_x"

func TestAccCertificateScepServerRaTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy: resource.ComposeAggregateTestCheckFunc(
					testCheckResourceDestroy("/certificate/scep-server/ra", "routeros_certificate_scep_server_ra"),
					testCheckResourceDestroy("/certificate", "routeros_system_certificate"),
				),
				Steps: []resource.TestStep{
					{
						Config: testAccCertificateScepServerRaConfig,
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testResourceCertificateScepServerRa),
							resource.TestCheckResourceAttr(testResourceCertificateScepServerRa, "name",
								"test_certificate_scep_ra_x"),
							resource.TestCheckResourceAttr(testResourceCertificateScepServerRa, "server_url",
								"http://192.0.2.40/scep"),
							resource.TestCheckResourceAttr(testResourceCertificateScepServerRa, "ca_identity",
								"test-certificate-scep-ra-x-ca"),
							resource.TestCheckResourceAttr(testResourceCertificateScepServerRa, "ra_path",
								"/scep/test_certificate_scep_ra_x"),
							resource.TestCheckResourceAttr(testResourceCertificateScepServerRa,
								"fingerprint_algorithm", "sha256"),
							resource.TestCheckResourceAttr(testResourceCertificateScepServerRa,
								"ra_transaction_lifetime", "1d"),
							resource.TestCheckResourceAttr(testResourceCertificateScepServerRa, "template",
								"test_certificate_scep_ra_x_tpl"),
							resource.TestCheckResourceAttr(testResourceCertificateScepServerRa, "disabled", "true"),
							resource.TestCheckResourceAttr(testResourceCertificateScepServerRa, "on_smart_card", "false"),
							resource.TestCheckResourceAttrSet(testResourceCertificateScepServerRa, "status"),
						),
					},
					{
						Config: testAccCertificateScepServerRaConfigUpdated,
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testResourceCertificateScepServerRa),
							resource.TestCheckResourceAttr(testResourceCertificateScepServerRa, "server_url",
								"http://192.0.2.41/scep"),
							resource.TestCheckResourceAttr(testResourceCertificateScepServerRa, "ca_identity",
								"test-certificate-scep-ra-x-ca-2"),
							resource.TestCheckResourceAttr(testResourceCertificateScepServerRa,
								"fingerprint_algorithm", "sha1"),
							resource.TestCheckResourceAttr(testResourceCertificateScepServerRa,
								"ra_transaction_lifetime", "2h"),
						),
					},
				},
			})

		})
	}
}

// The `template` of a registration authority has to reference a certificate that has no private
// key, i.e. a certificate template. A certificate created without a `sign` block is exactly that.
var testAccCertificateScepServerRaConfigDeps = providerConfig + `
resource "routeros_system_certificate" "test_certificate_scep_ra_x_tpl" {
	name        = "test_certificate_scep_ra_x_tpl"
	common_name = "test-certificate-scep-ra-x-tpl"
	key_size    = "1024"
	key_usage   = ["digital-signature", "key-encipherment", "tls-client"]
}
`

var testAccCertificateScepServerRaConfig = testAccCertificateScepServerRaConfigDeps + `
resource "routeros_certificate_scep_server_ra" "test_certificate_scep_ra_x" {
	name                    = "test_certificate_scep_ra_x"
	server_url              = "http://192.0.2.40/scep"
	ca_identity             = "test-certificate-scep-ra-x-ca"
	ra_path                 = "/scep/test_certificate_scep_ra_x"
	fingerprint_algorithm   = "sha256"
	ra_transaction_lifetime = "1d"
	template                = routeros_system_certificate.test_certificate_scep_ra_x_tpl.name
	disabled                = true
}
`

var testAccCertificateScepServerRaConfigUpdated = testAccCertificateScepServerRaConfigDeps + `
resource "routeros_certificate_scep_server_ra" "test_certificate_scep_ra_x" {
	name                    = "test_certificate_scep_ra_x"
	server_url              = "http://192.0.2.41/scep"
	ca_identity             = "test-certificate-scep-ra-x-ca-2"
	ra_path                 = "/scep/test_certificate_scep_ra_x"
	fingerprint_algorithm   = "sha1"
	ra_transaction_lifetime = "2h"
	template                = routeros_system_certificate.test_certificate_scep_ra_x_tpl.name
	disabled                = true
}
`
