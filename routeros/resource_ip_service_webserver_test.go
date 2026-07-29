package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIpServiceWebserverAddress = "routeros_ip_service_webserver.test_ip_service_webserver"

// `/ip/service/webserver` is a settings singleton (no CheckDestroy). Only `graphs_*` is toggled;
// writing `rest_*` would tear down the REST transport, and RouterOS refuses to disable `index_*` while WebFig/graphs are served.
func TestAccIpServiceWebserverTest_basic(t *testing.T) {
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
						Config: testAccIpServiceWebserverConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpServiceWebserverAddress),
							resource.TestCheckResourceAttr(testIpServiceWebserverAddress, "graphs_plain", "false"),
							resource.TestCheckResourceAttr(testIpServiceWebserverAddress, "graphs_secure", "false"),
							resource.TestCheckResourceAttr(testIpServiceWebserverAddress, "acme_plain", "true"),
							resource.TestCheckResourceAttr(testIpServiceWebserverAddress, "crl_plain", "true"),
							resource.TestCheckResourceAttr(testIpServiceWebserverAddress, "scep_plain", "true"),
						),
					},
					{
						Config: testAccIpServiceWebserverRestoredConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpServiceWebserverAddress),
							resource.TestCheckResourceAttr(testIpServiceWebserverAddress, "graphs_plain", "true"),
							resource.TestCheckResourceAttr(testIpServiceWebserverAddress, "graphs_secure", "true"),
						),
					},
				},
			})

		})
	}
}

func testAccIpServiceWebserverConfig() string {
	return providerConfig + `
resource "routeros_ip_service_webserver" "test_ip_service_webserver" {
	graphs_plain  = false
	graphs_secure = false
	acme_plain    = true
	crl_plain     = true
	scep_plain    = true
}`
}

func testAccIpServiceWebserverRestoredConfig() string {
	return providerConfig + `
resource "routeros_ip_service_webserver" "test_ip_service_webserver" {
	graphs_plain  = true
	graphs_secure = true
	acme_plain    = true
	crl_plain     = true
	scep_plain    = true
}`
}
