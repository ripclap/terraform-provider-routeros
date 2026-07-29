package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIpSMBUsersAddress = "routeros_ip_smb_users.test_ip_smb_users"

func TestAccIpSMBUsersTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ip/smb/users", "routeros_ip_smb_users"),
				Steps: []resource.TestStep{
					{
						Config: testAccIpSMBUsersConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpSMBUsersAddress),
							resource.TestCheckResourceAttr(testIpSMBUsersAddress, "name", "test_ip_smb_users"),
							resource.TestCheckResourceAttr(testIpSMBUsersAddress, "comment", "test_ip_smb_users"),
							resource.TestCheckResourceAttr(testIpSMBUsersAddress, "disabled", "true"),
							resource.TestCheckResourceAttr(testIpSMBUsersAddress, "password", "TestIpSmbUsers1"),
							resource.TestCheckResourceAttr(testIpSMBUsersAddress, "read_only", "true"),
							resource.TestCheckResourceAttr(testIpSMBUsersAddress, "default", "false"),
							resource.TestCheckResourceAttr(testIpSMBUsersAddress, "dynamic", "false"),
						),
					},
					{
						Config: testAccIpSMBUsersUpdatedConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpSMBUsersAddress),
							resource.TestCheckResourceAttr(testIpSMBUsersAddress, "comment", "test_ip_smb_users updated"),
							resource.TestCheckResourceAttr(testIpSMBUsersAddress, "password", "TestIpSmbUsers2"),
							resource.TestCheckResourceAttr(testIpSMBUsersAddress, "read_only", "false"),
						),
					},
				},
			})

		})
	}
}

func testAccIpSMBUsersConfig() string {
	return providerConfig + `

resource "routeros_ip_smb_users" "test_ip_smb_users" {
	name      = "test_ip_smb_users"
	comment   = "test_ip_smb_users"
	disabled  = true
	password  = "TestIpSmbUsers1"
	read_only = true
}

`
}

func testAccIpSMBUsersUpdatedConfig() string {
	return providerConfig + `

resource "routeros_ip_smb_users" "test_ip_smb_users" {
	name      = "test_ip_smb_users"
	comment   = "test_ip_smb_users updated"
	disabled  = true
	password  = "TestIpSmbUsers2"
	read_only = false
}

`
}
