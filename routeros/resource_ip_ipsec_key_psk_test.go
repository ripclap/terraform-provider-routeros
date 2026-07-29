package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIpIpsecKeyPsk = "routeros_ip_ipsec_key_psk.test_ipsec_key_psk_x"

// RouterOS requires `id` to be at least 16 characters long, `key` to be a hexadecimal string of at
// least 128 bytes (256 hex digits) and `peer` to be an existing IPsec peer. The device reports the
// key back in upper case, so the configuration has to use upper case as well.
const testIpIpsecKeyPskKey = "0123456789ABCDEF0123456789ABCDEF0123456789ABCDEF0123456789ABCDEF" +
	"0123456789ABCDEF0123456789ABCDEF0123456789ABCDEF0123456789ABCDEF" +
	"0123456789ABCDEF0123456789ABCDEF0123456789ABCDEF0123456789ABCDEF" +
	"0123456789ABCDEF0123456789ABCDEF0123456789ABCDEF0123456789ABCDEF"

func TestAccIpIpsecKeyPskTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ip/ipsec/key/psk", "routeros_ip_ipsec_key_psk"),
				Steps: []resource.TestStep{
					{
						Config: testAccIpIpsecKeyPskConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpIpsecKeyPsk),
							resource.TestCheckResourceAttr(testIpIpsecKeyPsk, "psk_id", "test_ipsec_key_psk_x_id"),
							resource.TestCheckResourceAttr(testIpIpsecKeyPsk, "peer", "test_ipsec_key_psk_peer_x"),
							resource.TestCheckResourceAttr(testIpIpsecKeyPsk, "key", testIpIpsecKeyPskKey),
						),
					},
					{
						// Every attribute is ForceNew: the entry is replaced, not modified.
						Config: testAccIpIpsecKeyPskUpdatedConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpIpsecKeyPsk),
							resource.TestCheckResourceAttr(testIpIpsecKeyPsk, "psk_id", "test_ipsec_key_psk_x_id2"),
							resource.TestCheckResourceAttr(testIpIpsecKeyPsk, "peer", "test_ipsec_key_psk_peer_x"),
						),
					},
				},
			})

		})
	}
}

func testAccIpIpsecKeyPskConfig() string {
	return providerConfig + `
resource "routeros_ip_ipsec_peer" "test_ipsec_key_psk_peer_x" {
	name          = "test_ipsec_key_psk_peer_x"
	address       = "192.0.2.70/32"
	exchange_mode = "ike2"
	disabled      = true
}

resource "routeros_ip_ipsec_key_psk" "test_ipsec_key_psk_x" {
	psk_id = "test_ipsec_key_psk_x_id"
	peer   = routeros_ip_ipsec_peer.test_ipsec_key_psk_peer_x.name
	key    = "` + testIpIpsecKeyPskKey + `"
}
`
}

func testAccIpIpsecKeyPskUpdatedConfig() string {
	return providerConfig + `
resource "routeros_ip_ipsec_peer" "test_ipsec_key_psk_peer_x" {
	name          = "test_ipsec_key_psk_peer_x"
	address       = "192.0.2.70/32"
	exchange_mode = "ike2"
	disabled      = true
}

resource "routeros_ip_ipsec_key_psk" "test_ipsec_key_psk_x" {
	psk_id = "test_ipsec_key_psk_x_id2"
	peer   = routeros_ip_ipsec_peer.test_ipsec_key_psk_peer_x.name
	key    = "` + testIpIpsecKeyPskKey + `"
}
`
}
