package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testSystemNtpKey = "routeros_system_ntp_key.test_system_ntp_key"

func TestAccSystemNtpKeyTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/system/ntp/key", "routeros_system_ntp_key"),
				Steps: []resource.TestStep{
					{
						Config: testAccSystemNtpKeyConfig("77", "test_system_ntp_key_val"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testSystemNtpKey),
							resource.TestCheckResourceAttr(testSystemNtpKey, "key_id", "77"),
							resource.TestCheckResourceAttr(testSystemNtpKey, "key_val",
								"test_system_ntp_key_val"),
						),
					},
					{
						Config: testAccSystemNtpKeyConfig("78", "test_system_ntp_key_val2"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testSystemNtpKey),
							resource.TestCheckResourceAttr(testSystemNtpKey, "key_id", "78"),
							resource.TestCheckResourceAttr(testSystemNtpKey, "key_val",
								"test_system_ntp_key_val2"),
						),
					},
				},
			})

		})
	}
}

func testAccSystemNtpKeyConfig(keyId, keyVal string) string {
	return fmt.Sprintf(`%v

resource "routeros_system_ntp_key" "test_system_ntp_key" {
	key_id  = %v
	key_val = "%v"
}
`, providerConfig, keyId, keyVal)
}
