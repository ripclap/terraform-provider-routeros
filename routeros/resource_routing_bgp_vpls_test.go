package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testRoutingBgpVpls = "routeros_routing_bgp_vpls.test_routing_bgp_vpls"

func TestAccRoutingBgpVplsTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/routing/bgp/vpls", "routeros_routing_bgp_vpls"),
				Steps: []resource.TestStep{
					{
						Config: testAccRoutingBgpVplsConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testRoutingBgpVpls),
							resource.TestCheckResourceAttr(testRoutingBgpVpls, "name", "test_routing_bgp_vpls"),
							resource.TestCheckResourceAttr(testRoutingBgpVpls, "rd", "65599:9"),
							resource.TestCheckResourceAttr(testRoutingBgpVpls, "site_id", "9"),
							resource.TestCheckResourceAttr(testRoutingBgpVpls, "pw_type", "vpls"),
							resource.TestCheckResourceAttr(testRoutingBgpVpls, "bridge", "bridge"),
							resource.TestCheckTypeSetElemAttr(testRoutingBgpVpls, "import_route_targets.*", "65599:9"),
							resource.TestCheckTypeSetElemAttr(testRoutingBgpVpls, "export_route_targets.*", "65599:9"),
						),
					},
					{
						Config: testAccRoutingBgpVplsUpdatedConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testRoutingBgpVpls),
							resource.TestCheckResourceAttr(testRoutingBgpVpls, "comment", "acc test routing bgp vpls"),
							resource.TestCheckResourceAttr(testRoutingBgpVpls, "disabled", "true"),
							resource.TestCheckResourceAttr(testRoutingBgpVpls, "bridge_cost", "50"),
							resource.TestCheckResourceAttr(testRoutingBgpVpls, "bridge_horizon", "1"),
							resource.TestCheckResourceAttr(testRoutingBgpVpls, "bridge_pvid", "99"),
							resource.TestCheckResourceAttr(testRoutingBgpVpls, "pw_control_word", "enabled"),
							resource.TestCheckResourceAttr(testRoutingBgpVpls, "pw_l2mtu", "1500"),
							resource.TestCheckResourceAttr(testRoutingBgpVpls, "pw_type", "raw-ethernet"),
							resource.TestCheckResourceAttr(testRoutingBgpVpls, "vrf", "main"),
							resource.TestCheckTypeSetElemAttr(testRoutingBgpVpls, "import_route_targets.*", "65599:19"),
						),
					},
				},
			})

		})
	}
}

func testAccRoutingBgpVplsConfig() string {
	return fmt.Sprintf(`%v

resource "routeros_routing_bgp_vpls" "test_routing_bgp_vpls" {
  name                 = "test_routing_bgp_vpls"
  bridge               = "bridge"
  rd                   = "65599:9"
  site_id              = 9
  import_route_targets = ["65599:9"]
  export_route_targets = ["65599:9"]
  pw_type              = "vpls"
}
`, providerConfig)
}

func testAccRoutingBgpVplsUpdatedConfig() string {
	return fmt.Sprintf(`%v

resource "routeros_routing_bgp_vpls" "test_routing_bgp_vpls" {
  name                 = "test_routing_bgp_vpls"
  comment              = "acc test routing bgp vpls"
  disabled             = true
  bridge               = "bridge"
  bridge_cost          = 50
  bridge_horizon       = "1"
  bridge_pvid          = 99
  rd                   = "65599:9"
  site_id              = 9
  import_route_targets = ["65599:9", "65599:19"]
  export_route_targets = ["65599:9"]
  pw_control_word      = "enabled"
  pw_l2mtu             = 1500
  pw_type              = "raw-ethernet"
  vrf                  = "main"
}
`, providerConfig)
}
