package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testInterfaceIpipv6 = "routeros_interface_ipipv6.test_ipipv6_x"

func TestAccInterfaceIpipv6Test_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/interface/ipipv6", "routeros_interface_ipipv6"),
				Steps: []resource.TestStep{
					{
						Config: testAccInterfaceIpipv6Config("2001:db8:922::2", "1400"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfaceIpipv6),
							resource.TestCheckResourceAttr(testInterfaceIpipv6, "name", "test_ipipv6_x"),
							resource.TestCheckResourceAttr(testInterfaceIpipv6, "local_address", "2001:db8:922::1"),
							resource.TestCheckResourceAttr(testInterfaceIpipv6, "remote_address", "2001:db8:922::2"),
							resource.TestCheckResourceAttr(testInterfaceIpipv6, "mtu", "1400"),
							resource.TestCheckResourceAttr(testInterfaceIpipv6, "dont_fragment", "no"),
							resource.TestCheckResourceAttr(testInterfaceIpipv6, "dscp", "32"),
							resource.TestCheckResourceAttr(testInterfaceIpipv6, "clamp_tcp_mss", "true"),
							resource.TestCheckResourceAttr(testInterfaceIpipv6, "keepalive", "10s,5"),
							resource.TestCheckResourceAttr(testInterfaceIpipv6, "comment", "test_ipipv6_x"),
							resource.TestCheckResourceAttr(testInterfaceIpipv6, "disabled", "false"),
							resource.TestCheckResourceAttr(testInterfaceIpipv6, "running", "false"),
							resource.TestCheckResourceAttrSet(testInterfaceIpipv6, "actual_mtu"),
						),
					},
					{
						Config: testAccInterfaceIpipv6Config("2001:db8:922::3", "1350"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testInterfaceIpipv6),
							resource.TestCheckResourceAttr(testInterfaceIpipv6, "remote_address", "2001:db8:922::3"),
							resource.TestCheckResourceAttr(testInterfaceIpipv6, "mtu", "1350"),
						),
					},
				},
			})
		})
	}
}

// Endpoints use the 2001:db8::/32 documentation prefix, so the interface never comes up. `ipsec_secret`
// is omitted on purpose: it would install a dynamic peer and policy in the shared global /ip/ipsec tree.
func testAccInterfaceIpipv6Config(remote, mtu string) string {
	return providerConfig + `

resource "routeros_interface_ipipv6" "test_ipipv6_x" {
	name           = "test_ipipv6_x"
	local_address  = "2001:db8:922::1"
	remote_address = "` + remote + `"
	mtu            = "` + mtu + `"
	dont_fragment  = "no"
	dscp           = "32"
	clamp_tcp_mss  = true
	keepalive      = "10s,5"
	comment        = "test_ipipv6_x"
}
`
}
