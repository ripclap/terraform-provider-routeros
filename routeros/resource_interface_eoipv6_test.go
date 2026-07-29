package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testInterfaceEoipv6 = "routeros_interface_eoipv6.test_eoipv6_x"

func TestAccInterfaceEoipv6Test_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/interface/eoipv6", "routeros_interface_eoipv6"),
				Steps: []resource.TestStep{
					{
						Config: testAccInterfaceEoipv6Config("2001:db8:921::2", "3921"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfaceEoipv6),
							resource.TestCheckResourceAttr(testInterfaceEoipv6, "name", "test_eoipv6_x"),
							resource.TestCheckResourceAttr(testInterfaceEoipv6, "local_address", "2001:db8:921::1"),
							resource.TestCheckResourceAttr(testInterfaceEoipv6, "remote_address", "2001:db8:921::2"),
							resource.TestCheckResourceAttr(testInterfaceEoipv6, "tunnel_id", "3921"),
							resource.TestCheckResourceAttr(testInterfaceEoipv6, "mac_address", "02:00:5E:00:39:21"),
							resource.TestCheckResourceAttr(testInterfaceEoipv6, "mtu", "1400"),
							resource.TestCheckResourceAttr(testInterfaceEoipv6, "dont_fragment", "no"),
							resource.TestCheckResourceAttr(testInterfaceEoipv6, "dscp", "32"),
							resource.TestCheckResourceAttr(testInterfaceEoipv6, "clamp_tcp_mss", "true"),
							resource.TestCheckResourceAttr(testInterfaceEoipv6, "keepalive", "10s,5"),
							resource.TestCheckResourceAttr(testInterfaceEoipv6, "loop_protect", "on"),
							resource.TestCheckResourceAttr(testInterfaceEoipv6, "arp", "enabled"),
							resource.TestCheckResourceAttr(testInterfaceEoipv6, "comment", "test_eoipv6_x"),
							resource.TestCheckResourceAttr(testInterfaceEoipv6, "disabled", "false"),
							resource.TestCheckResourceAttr(testInterfaceEoipv6, "running", "false"),
							resource.TestCheckResourceAttrSet(testInterfaceEoipv6, "actual_mtu"),
							resource.TestCheckResourceAttrSet(testInterfaceEoipv6, "l2mtu"),
						),
					},
					{
						Config: testAccInterfaceEoipv6Config("2001:db8:921::3", "3922"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfaceEoipv6),
							resource.TestCheckResourceAttr(testInterfaceEoipv6, "remote_address", "2001:db8:921::3"),
							resource.TestCheckResourceAttr(testInterfaceEoipv6, "tunnel_id", "3922"),
						),
					},
				},
			})
		})
	}
}

// Endpoints use the 2001:db8::/32 documentation prefix, so the tunnel never comes up; RouterOS still
// stores every property, which exercises the whole schema.
func testAccInterfaceEoipv6Config(remote, tunnelId string) string {
	return providerConfig + `

resource "routeros_interface_eoipv6" "test_eoipv6_x" {
	name           = "test_eoipv6_x"
	local_address  = "2001:db8:921::1"
	remote_address = "` + remote + `"
	tunnel_id      = ` + tunnelId + `
	mac_address    = "02:00:5E:00:39:21"
	mtu            = "1400"
	dont_fragment  = "no"
	dscp           = "32"
	clamp_tcp_mss  = true
	keepalive      = "10s,5"
	loop_protect   = "on"
	arp            = "enabled"
	comment        = "test_eoipv6_x"
}
`
}
