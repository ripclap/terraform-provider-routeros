package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testPPPL2tpSecretAddress = "routeros_ppp_l2tp_secret.test_ppp_l2tp_secret_x"

// RouterOS normalises the `address` property to `network/prefix`, so the configuration below uses
// that form directly.
func TestAccPPPL2tpSecretTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ppp/l2tp-secret", "routeros_ppp_l2tp_secret"),
				Steps: []resource.TestStep{
					{
						Config: testAccPPPL2tpSecretConfig("192.0.2.0/24", "test_ppp_l2tp_secret_x_key"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testPPPL2tpSecretAddress),
							resource.TestCheckResourceAttr(testPPPL2tpSecretAddress, "address", "192.0.2.0/24"),
							resource.TestCheckResourceAttr(testPPPL2tpSecretAddress, "secret",
								"test_ppp_l2tp_secret_x_key"),
							resource.TestCheckResourceAttr(testPPPL2tpSecretAddress, "comment",
								"test_ppp_l2tp_secret_x"),
						),
					},
					{
						Config: testAccPPPL2tpSecretConfig("192.0.2.8/32", "test_ppp_l2tp_secret_x_key2"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testPPPL2tpSecretAddress),
							resource.TestCheckResourceAttr(testPPPL2tpSecretAddress, "address", "192.0.2.8/32"),
							resource.TestCheckResourceAttr(testPPPL2tpSecretAddress, "secret",
								"test_ppp_l2tp_secret_x_key2"),
						),
					},
				},
			})

		})
	}
}

func testAccPPPL2tpSecretConfig(address, secret string) string {
	return providerConfig + `

resource "routeros_ppp_l2tp_secret" "test_ppp_l2tp_secret_x" {
	address = "` + address + `"
	secret  = "` + secret + `"
	comment = "test_ppp_l2tp_secret_x"
}

`
}
