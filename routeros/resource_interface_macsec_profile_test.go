package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testInterfaceMacsecProfileAddress = "routeros_interface_macsec_profile.test_macsec_profile_x"

func TestAccInterfaceMacsecProfileTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy: testCheckResourceDestroy("/interface/macsec/profile",
					"routeros_interface_macsec_profile"),
				Steps: []resource.TestStep{
					{
						Config: testAccInterfaceMacsecProfileConfig("aes-gcm-xpn-128", 42),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfaceMacsecProfileAddress),
							resource.TestCheckResourceAttr(testInterfaceMacsecProfileAddress, "name", "test_macsec_profile_x"),
							resource.TestCheckResourceAttr(testInterfaceMacsecProfileAddress, "ciphers", "aes-gcm-xpn-128"),
							resource.TestCheckResourceAttr(testInterfaceMacsecProfileAddress, "server_priority", "42"),
							resource.TestCheckResourceAttr(testInterfaceMacsecProfileAddress, "default", "false"),
						),
					},
					{
						Config: testAccInterfaceMacsecProfileConfig("aes-gcm-128", 200),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfaceMacsecProfileAddress),
							resource.TestCheckResourceAttr(testInterfaceMacsecProfileAddress, "ciphers", "aes-gcm-128"),
							resource.TestCheckResourceAttr(testInterfaceMacsecProfileAddress, "server_priority", "200"),
						),
					},
				},
			})

		})
	}
}

func testAccInterfaceMacsecProfileConfig(ciphers string, priority int) string {
	return providerConfig + fmt.Sprintf(`
resource "routeros_interface_macsec_profile" "test_macsec_profile_x" {
	name            = "test_macsec_profile_x"
	ciphers         = "%v"
	server_priority = %v
}
`, ciphers, priority)
}
