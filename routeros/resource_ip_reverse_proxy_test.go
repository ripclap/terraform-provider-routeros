package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIpReverseProxyAddress = "routeros_ip_reverse_proxy.test_ip_reverse_proxy"

// The reverse proxy binds TCP/443 (owned by `www-ssl`, the provider's REST transport), so every entry
// is kept `disabled = true` to avoid cutting off the management plane.
func TestAccIpReverseProxyTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ip/reverse-proxy", "routeros_ip_reverse_proxy"),
				Steps: []resource.TestStep{
					{
						Config: testAccIpReverseProxyConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpReverseProxyAddress),
							resource.TestCheckResourceAttr(testIpReverseProxyAddress, "comment", "test_ip_reverse_proxy"),
							resource.TestCheckResourceAttr(testIpReverseProxyAddress, "disabled", "true"),
							resource.TestCheckResourceAttr(testIpReverseProxyAddress, "dynamic", "false"),
							resource.TestCheckResourceAttr(testIpReverseProxyAddress, "ip_address", "192.0.2.61"),
							resource.TestCheckResourceAttr(testIpReverseProxyAddress, "port", "8461"),
							resource.TestCheckResourceAttr(testIpReverseProxyAddress, "sni", "test-ip-reverse-proxy.example.com"),
							resource.TestCheckResourceAttr(testIpReverseProxyAddress, "certificate", "none"),
						),
					},
					{
						Config: testAccIpReverseProxyUpdatedConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpReverseProxyAddress),
							resource.TestCheckResourceAttr(testIpReverseProxyAddress, "comment", "test_ip_reverse_proxy updated"),
							resource.TestCheckResourceAttr(testIpReverseProxyAddress, "ip_address", "192.0.2.62"),
							resource.TestCheckResourceAttr(testIpReverseProxyAddress, "port", "8462"),
							resource.TestCheckResourceAttr(testIpReverseProxyAddress, "sni", "updated.test-ip-reverse-proxy.example.com"),
						),
					},
				},
			})

		})
	}
}

func testAccIpReverseProxyConfig() string {
	return providerConfig + `

resource "routeros_ip_reverse_proxy" "test_ip_reverse_proxy" {
	comment    = "test_ip_reverse_proxy"
	disabled   = true
	ip_address = "192.0.2.61"
	port       = 8461
	sni        = "test-ip-reverse-proxy.example.com"
}

`
}

func testAccIpReverseProxyUpdatedConfig() string {
	return providerConfig + `

resource "routeros_ip_reverse_proxy" "test_ip_reverse_proxy" {
	comment    = "test_ip_reverse_proxy updated"
	disabled   = true
	ip_address = "192.0.2.62"
	port       = 8462
	sni        = "updated.test-ip-reverse-proxy.example.com"
}

`
}
