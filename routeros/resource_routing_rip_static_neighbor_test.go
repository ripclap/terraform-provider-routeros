package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testRoutingRipStaticNeighbor = "routeros_routing_rip_static_neighbor.test_rip_static_neighbor"

func TestAccRoutingRipStaticNeighborTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy: testCheckResourceDestroy("/routing/rip/static-neighbor",
					"routeros_routing_rip_static_neighbor"),
				Steps: []resource.TestStep{
					{
						Config: testAccRoutingRipStaticNeighborConfig("192.0.2.2", "false"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testRoutingRipStaticNeighbor),
							resource.TestCheckResourceAttr(testRoutingRipStaticNeighbor, "address", "192.0.2.2"),
							resource.TestCheckResourceAttr(testRoutingRipStaticNeighbor, "instance",
								"test_rip_static_neighbor_inst"),
							resource.TestCheckResourceAttr(testRoutingRipStaticNeighbor, "disabled", "false"),
						),
					},
					{
						Config: testAccRoutingRipStaticNeighborConfig("192.0.2.3", "true"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testRoutingRipStaticNeighbor),
							resource.TestCheckResourceAttr(testRoutingRipStaticNeighbor, "address", "192.0.2.3"),
							resource.TestCheckResourceAttr(testRoutingRipStaticNeighbor, "disabled", "true"),
						),
					},
				},
			})

		})
	}
}

func testAccRoutingRipStaticNeighborConfig(address, disabled string) string {
	return providerConfig + `

resource "routeros_routing_rip_instance" "test_rip_static_neighbor_inst" {
	name = "test_rip_static_neighbor_inst"
	afi  = "ip"
}

resource "routeros_routing_rip_static_neighbor" "test_rip_static_neighbor" {
	address  = "` + address + `"
	instance = routeros_routing_rip_instance.test_rip_static_neighbor_inst.name
	disabled = ` + disabled + `
}
`
}
