package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIpIpsecKeyQkd = "routeros_ip_ipsec_key_qkd.test_ipsec_key_qkd_x"

// `/ip/ipsec/key/qkd` is a settings singleton: it cannot be created or destroyed, only modified,
// so the test follows the "settings" pattern and carries no CheckDestroy.
func TestAccIpIpsecKeyQkdTest_basic(t *testing.T) {
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
						Config: testAccIpIpsecKeyQkdConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpIpsecKeyQkd),
							resource.TestCheckResourceAttr(testIpIpsecKeyQkd, "address", "192.0.2.75:8020"),
							resource.TestCheckResourceAttr(testIpIpsecKeyQkd, "kme_id", "test_ipsec_key_qkd_kme_x"),
							resource.TestCheckResourceAttr(testIpIpsecKeyQkd, "peer_sae_id", "test_ipsec_key_qkd_sae_x"),
							resource.TestCheckResourceAttr(testIpIpsecKeyQkd, "cache_size", "4"),
							resource.TestCheckResourceAttr(testIpIpsecKeyQkd, "key_size", "256"),
							resource.TestCheckResourceAttr(testIpIpsecKeyQkd, "enabled", "false"),
						),
					},
					{
						// Restores the RouterOS defaults of the menu.
						Config: testAccIpIpsecKeyQkdRestoreConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpIpsecKeyQkd),
							resource.TestCheckResourceAttr(testIpIpsecKeyQkd, "address", ""),
							resource.TestCheckResourceAttr(testIpIpsecKeyQkd, "kme_id", ""),
							resource.TestCheckResourceAttr(testIpIpsecKeyQkd, "peer_sae_id", ""),
							resource.TestCheckResourceAttr(testIpIpsecKeyQkd, "cache_size", "2"),
							resource.TestCheckResourceAttr(testIpIpsecKeyQkd, "key_size", "128"),
						),
					},
				},
			})

		})
	}
}

func testAccIpIpsecKeyQkdConfig() string {
	return providerConfig + `
resource "routeros_ip_ipsec_key_qkd" "test_ipsec_key_qkd_x" {
	address     = "192.0.2.75:8020"
	kme_id      = "test_ipsec_key_qkd_kme_x"
	peer_sae_id = "test_ipsec_key_qkd_sae_x"
	cache_size  = 4
	key_size    = 256
	enabled     = false
}
`
}

func testAccIpIpsecKeyQkdRestoreConfig() string {
	return providerConfig + `
resource "routeros_ip_ipsec_key_qkd" "test_ipsec_key_qkd_x" {
	address     = ""
	kme_id      = ""
	peer_sae_id = ""
	cache_size  = 2
	key_size    = 128
	enabled     = false
}
`
}
