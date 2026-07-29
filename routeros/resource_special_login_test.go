package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testSpecialLogin = "routeros_special_login.test_special_login"

func TestAccSpecialLoginTest_basic(t *testing.T) {
	testCheckMenu(t, "/port")
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/special-login", "routeros_special_login"),
				Steps: []resource.TestStep{
					{
						Config: testAccSpecialLoginConfig("true"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testSpecialLogin),
							resource.TestCheckResourceAttr(testSpecialLogin, "port", "serial0"),
							resource.TestCheckResourceAttr(testSpecialLogin, "user", "admin"),
							resource.TestCheckResourceAttr(testSpecialLogin, "channel", "0"),
							resource.TestCheckResourceAttr(testSpecialLogin, "disabled", "true"),
						),
					},
					{
						Config: testAccSpecialLoginConfig("false"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testSpecialLogin),
							resource.TestCheckResourceAttr(testSpecialLogin, "disabled", "false"),
						),
					},
				},
			})

		})
	}
}

func testAccSpecialLoginConfig(disabled string) string {
	return fmt.Sprintf(`%v

resource "routeros_special_login" "test_special_login" {
	port     = "serial0"
	user     = "admin"
	channel  = 0
	disabled = %v
}
`, providerConfig, disabled)
}
