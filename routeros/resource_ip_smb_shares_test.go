package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIpSMBSharesAddress = "routeros_ip_smb_shares.test_ip_smb_shares"

func TestAccIpSMBSharesTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ip/smb/shares", "routeros_ip_smb_shares"),
				Steps: []resource.TestStep{
					{
						Config: testAccIpSMBSharesConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpSMBSharesAddress),
							resource.TestCheckResourceAttr(testIpSMBSharesAddress, "name", "test_ip_smb_shares"),
							resource.TestCheckResourceAttr(testIpSMBSharesAddress, "comment", "test_ip_smb_shares"),
							resource.TestCheckResourceAttr(testIpSMBSharesAddress, "directory", "/test_ip_smb_shares"),
							resource.TestCheckResourceAttr(testIpSMBSharesAddress, "disabled", "true"),
							resource.TestCheckResourceAttr(testIpSMBSharesAddress, "read_only", "true"),
							resource.TestCheckResourceAttr(testIpSMBSharesAddress, "require_encryption", "true"),
							resource.TestCheckResourceAttr(testIpSMBSharesAddress, "default", "false"),
							resource.TestCheckResourceAttr(testIpSMBSharesAddress, "dynamic", "false"),
							resource.TestCheckResourceAttr(testIpSMBSharesAddress, "valid_users.#", "1"),
							resource.TestCheckTypeSetElemAttr(testIpSMBSharesAddress, "valid_users.*", "test_ip_smb_shares_ok"),
							resource.TestCheckResourceAttr(testIpSMBSharesAddress, "invalid_users.#", "1"),
							resource.TestCheckTypeSetElemAttr(testIpSMBSharesAddress, "invalid_users.*", "test_ip_smb_shares_no"),
						),
					},
					{
						Config: testAccIpSMBSharesUpdatedConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpSMBSharesAddress),
							resource.TestCheckResourceAttr(testIpSMBSharesAddress, "comment", "test_ip_smb_shares updated"),
							resource.TestCheckResourceAttr(testIpSMBSharesAddress, "directory", "/test_ip_smb_shares_updated"),
							resource.TestCheckResourceAttr(testIpSMBSharesAddress, "read_only", "false"),
							resource.TestCheckResourceAttr(testIpSMBSharesAddress, "require_encryption", "false"),
							resource.TestCheckResourceAttr(testIpSMBSharesAddress, "valid_users.#", "2"),
							resource.TestCheckTypeSetElemAttr(testIpSMBSharesAddress, "valid_users.*", "test_ip_smb_shares_ok"),
							resource.TestCheckTypeSetElemAttr(testIpSMBSharesAddress, "valid_users.*", "test_ip_smb_shares_no"),
						),
					},
				},
			})

		})
	}
}

// The share references SMB users, so they are created by the very same configuration.
const testAccIpSMBSharesUsers = `

resource "routeros_ip_smb_users" "test_ip_smb_shares_ok" {
	name     = "test_ip_smb_shares_ok"
	password = "TestIpSmbShares1"
	disabled = true
}

resource "routeros_ip_smb_users" "test_ip_smb_shares_no" {
	name     = "test_ip_smb_shares_no"
	password = "TestIpSmbShares2"
	disabled = true
}

`

func testAccIpSMBSharesConfig() string {
	return providerConfig + testAccIpSMBSharesUsers + `

resource "routeros_ip_smb_shares" "test_ip_smb_shares" {
	name               = "test_ip_smb_shares"
	comment            = "test_ip_smb_shares"
	directory          = "/test_ip_smb_shares"
	disabled           = true
	read_only          = true
	require_encryption = true
	valid_users        = [routeros_ip_smb_users.test_ip_smb_shares_ok.name]
	invalid_users      = [routeros_ip_smb_users.test_ip_smb_shares_no.name]
}

`
}

func testAccIpSMBSharesUpdatedConfig() string {
	return providerConfig + testAccIpSMBSharesUsers + `

resource "routeros_ip_smb_shares" "test_ip_smb_shares" {
	name               = "test_ip_smb_shares"
	comment            = "test_ip_smb_shares updated"
	directory          = "/test_ip_smb_shares_updated"
	disabled           = true
	read_only          = false
	require_encryption = false
	valid_users = [
		routeros_ip_smb_users.test_ip_smb_shares_ok.name,
		routeros_ip_smb_users.test_ip_smb_shares_no.name,
	]
	invalid_users = []
}

`
}
