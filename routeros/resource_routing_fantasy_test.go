package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testRoutingFantasy = "routeros_routing_fantasy.test_routing_fantasy"

// An enabled entry immediately injects 'route_count' synthetic routes into the routing table, so the test
// keeps the generator disabled and only verifies that the configuration round-trips.
func TestAccRoutingFantasyTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/routing/fantasy", "routeros_routing_fantasy"),
				Steps: []resource.TestStep{
					{
						Config: testAccRoutingFantasyConfig("10", "acc test routing fantasy"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testRoutingFantasy),
							resource.TestCheckResourceAttr(testRoutingFantasy, "name", "test_routing_fantasy"),
							resource.TestCheckResourceAttr(testRoutingFantasy, "comment", "acc test routing fantasy"),
							resource.TestCheckResourceAttr(testRoutingFantasy, "disabled", "true"),
							resource.TestCheckResourceAttr(testRoutingFantasy, "route_count", "10"),
							resource.TestCheckResourceAttr(testRoutingFantasy, "dst_address", "192.0.2.0"),
							resource.TestCheckResourceAttr(testRoutingFantasy, "prefix_length", "24"),
							resource.TestCheckResourceAttr(testRoutingFantasy, "gateway", "127.0.0.1"),
							resource.TestCheckResourceAttr(testRoutingFantasy, "offset", "1"),
							resource.TestCheckResourceAttr(testRoutingFantasy, "scope", "30"),
							resource.TestCheckResourceAttr(testRoutingFantasy, "target_scope", "10"),
							resource.TestCheckResourceAttr(testRoutingFantasy, "seed", "7"),
							resource.TestCheckResourceAttr(testRoutingFantasy, "instance_id", "1"),
							resource.TestCheckResourceAttr(testRoutingFantasy, "dealer_id", "1"),
							resource.TestCheckResourceAttr(testRoutingFantasy, "priv_offs", "0"),
							resource.TestCheckResourceAttr(testRoutingFantasy, "priv_size", "0"),
							resource.TestCheckResourceAttr(testRoutingFantasy, "use_hold", "false"),
						),
					},
					{
						Config: testAccRoutingFantasyConfig("25", "acc test routing fantasy updated"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testRoutingFantasy),
							resource.TestCheckResourceAttr(testRoutingFantasy, "route_count", "25"),
							resource.TestCheckResourceAttr(testRoutingFantasy, "comment", "acc test routing fantasy updated"),
						),
					},
				},
			})

		})
	}
}

func testAccRoutingFantasyConfig(count, comment string) string {
	return fmt.Sprintf(`%v

resource "routeros_routing_fantasy" "test_routing_fantasy" {
  name          = "test_routing_fantasy"
  comment       = "%v"
  disabled      = true
  route_count   = "%v"
  dst_address   = "192.0.2.0"
  prefix_length = "24"
  gateway       = "127.0.0.1"
  offset        = "1"
  scope         = "30"
  target_scope  = "10"
  seed          = "7"
  instance_id   = "1"
  dealer_id     = "1"
  priv_offs     = "0"
  priv_size     = "0"
  use_hold      = false
}
`, providerConfig, comment, count)
}
