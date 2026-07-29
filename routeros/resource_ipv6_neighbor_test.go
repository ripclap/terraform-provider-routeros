package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIPv6Neighbor = "routeros_ipv6_neighbor.test_ipv6_neighbor_x"

func TestAccIPv6NeighborTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/ipv6/neighbor", "routeros_ipv6_neighbor"),
				Steps: []resource.TestStep{
					{
						Config: testAccIPv6NeighborConfig("00:00:5E:00:53:74"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIPv6Neighbor),
							resource.TestCheckResourceAttr(testIPv6Neighbor, "address", "2001:db8:7f00:74::1"),
							resource.TestCheckResourceAttr(testIPv6Neighbor, "interface", "test_ipv6_neighbor_x_br"),
							resource.TestCheckResourceAttr(testIPv6Neighbor, "mac_address", "00:00:5E:00:53:74"),
							resource.TestCheckResourceAttr(testIPv6Neighbor, "comment", "test_ipv6_neighbor_x"),
							resource.TestCheckResourceAttr(testIPv6Neighbor, "dynamic", "false"),
							resource.TestCheckResourceAttr(testIPv6Neighbor, "status", "permanent"),
							resource.TestCheckResourceAttr(testIPv6Neighbor, "vrf", "main"),
						),
					},
					{
						Config: testAccIPv6NeighborConfig("00:00:5E:00:53:75"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIPv6Neighbor),
							resource.TestCheckResourceAttr(testIPv6Neighbor, "mac_address", "00:00:5E:00:53:75"),
						),
					},
				},
			})

		})
	}
}

func testAccIPv6NeighborConfig(mac string) string {
	return fmt.Sprintf(`%v

// A neighbor entry only becomes "permanent" while its interface is running, so the test
// uses its own always-up bridge instead of a physical port that may be down.
resource "routeros_interface_bridge" "test_ipv6_neighbor_x_br" {
  name = "test_ipv6_neighbor_x_br"
}

resource "routeros_ipv6_neighbor" "test_ipv6_neighbor_x" {
  address     = "2001:db8:7f00:74::1"
  interface   = routeros_interface_bridge.test_ipv6_neighbor_x_br.name
  mac_address = %q
  comment     = "test_ipv6_neighbor_x"
}
`, providerConfig, mac)
}
