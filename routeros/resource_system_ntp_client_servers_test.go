package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testSystemNtpClientServers = "routeros_system_ntp_client_servers.test_system_ntp_client_servers"

func TestAccSystemNtpClientServersTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy: testCheckResourceDestroy("/system/ntp/client/servers",
					"routeros_system_ntp_client_servers"),
				Steps: []resource.TestStep{
					{
						Config: testAccSystemNtpClientServersConfig("6", "10", "true"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testSystemNtpClientServers),
							resource.TestCheckResourceAttr(testSystemNtpClientServers, "address", "198.51.100.10"),
							resource.TestCheckResourceAttr(testSystemNtpClientServers, "comment",
								"test_system_ntp_client_servers"),
							resource.TestCheckResourceAttr(testSystemNtpClientServers, "min_poll", "6"),
							resource.TestCheckResourceAttr(testSystemNtpClientServers, "max_poll", "10"),
							resource.TestCheckResourceAttr(testSystemNtpClientServers, "iburst", "true"),
							resource.TestCheckResourceAttr(testSystemNtpClientServers, "auth_key", "none"),
							resource.TestCheckResourceAttr(testSystemNtpClientServers, "disabled", "false"),
							resource.TestCheckResourceAttr(testSystemNtpClientServers, "dynamic", "false"),
						),
					},
					{
						Config: testAccSystemNtpClientServersConfig("7", "12", "false"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testSystemNtpClientServers),
							resource.TestCheckResourceAttr(testSystemNtpClientServers, "min_poll", "7"),
							resource.TestCheckResourceAttr(testSystemNtpClientServers, "max_poll", "12"),
							resource.TestCheckResourceAttr(testSystemNtpClientServers, "iburst", "false"),
						),
					},
				},
			})

		})
	}
}

func testAccSystemNtpClientServersConfig(minPoll, maxPoll, iburst string) string {
	return fmt.Sprintf(`%v

resource "routeros_system_ntp_client_servers" "test_system_ntp_client_servers" {
	address  = "198.51.100.10"
	comment  = "test_system_ntp_client_servers"
	min_poll = %v
	max_poll = %v
	iburst   = %v
	auth_key = "none"
	disabled = false
}
`, providerConfig, minPoll, maxPoll, iburst)
}
