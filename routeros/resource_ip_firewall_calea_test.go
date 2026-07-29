package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIPFirewallCalea = "routeros_ip_firewall_calea.test_calea_x"

func TestAccIPFirewallCaleaTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ip/firewall/calea", "routeros_ip_firewall_calea"),
				Steps: []resource.TestStep{
					{
						Config: testAccIPFirewallCaleaConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIPFirewallCalea),
							resource.TestCheckResourceAttr(testIPFirewallCalea, "chain", "forward"),
							resource.TestCheckResourceAttr(testIPFirewallCalea, "action", "sniff-pc"),
							resource.TestCheckResourceAttr(testIPFirewallCalea, "sniff_id", "70"),
							resource.TestCheckResourceAttr(testIPFirewallCalea, "sniff_target", "192.0.2.70"),
							resource.TestCheckResourceAttr(testIPFirewallCalea, "sniff_target_port", "3799"),
							resource.TestCheckResourceAttr(testIPFirewallCalea, "comment", "test_calea_x"),
							resource.TestCheckResourceAttr(testIPFirewallCalea, "disabled", "true"),
							resource.TestCheckResourceAttr(testIPFirewallCalea, "protocol", "tcp"),
							resource.TestCheckResourceAttr(testIPFirewallCalea, "dst_port", "8080"),
							resource.TestCheckResourceAttr(testIPFirewallCalea, "src_address", "198.51.100.70"),
						),
					},
					{
						Config: testAccIPFirewallCaleaUpdatedConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIPFirewallCalea),
							resource.TestCheckResourceAttr(testIPFirewallCalea, "action", "sniff"),
							resource.TestCheckResourceAttr(testIPFirewallCalea, "sniff_target", "192.0.2.71"),
							resource.TestCheckResourceAttr(testIPFirewallCalea, "sniff_target_port", "3800"),
							resource.TestCheckResourceAttr(testIPFirewallCalea, "comment", "test_calea_x_upd"),
							resource.TestCheckResourceAttr(testIPFirewallCalea, "dst_address", "203.0.113.0/24"),
						),
					},
				},
			})

		})
	}
}

func testAccIPFirewallCaleaConfig() string {
	return providerConfig + `
resource "routeros_ip_firewall_calea" "test_calea_x" {
	chain             = "forward"
	action            = "sniff-pc"
	sniff_id          = 70
	sniff_target      = "192.0.2.70"
	sniff_target_port = 3799
	protocol          = "tcp"
	dst_port          = "8080"
	src_address       = "198.51.100.70"
	comment           = "test_calea_x"
	disabled          = true
}
`
}

func testAccIPFirewallCaleaUpdatedConfig() string {
	return providerConfig + `
resource "routeros_ip_firewall_calea" "test_calea_x" {
	chain             = "forward"
	action            = "sniff"
	sniff_target      = "192.0.2.71"
	sniff_target_port = 3800
	protocol          = "tcp"
	dst_address       = "203.0.113.0/24"
	comment           = "test_calea_x_upd"
	disabled          = true
}
`
}
