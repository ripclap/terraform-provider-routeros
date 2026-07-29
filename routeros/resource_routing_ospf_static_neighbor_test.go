package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testRoutingOspfStaticNeighbor = "routeros_routing_ospf_static_neighbor.test_ospf_neighbor_x"

func TestAccRoutingOspfStaticNeighborTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy: resource.ComposeAggregateTestCheckFunc(
					testCheckResourceDestroy("/routing/ospf/static-neighbor", "routeros_routing_ospf_static_neighbor"),
					testCheckResourceDestroy("/routing/ospf/area", "routeros_routing_ospf_area"),
					testCheckResourceDestroy("/routing/ospf/instance", "routeros_routing_ospf_instance"),
				),
				Steps: []resource.TestStep{
					{
						Config: testAccRoutingOspfStaticNeighborConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testRoutingOspfStaticNeighbor),
							resource.TestCheckResourceAttr(testRoutingOspfStaticNeighbor, "address", "192.0.2.211%ether5"),
							resource.TestCheckResourceAttr(testRoutingOspfStaticNeighbor, "area", "test_ospf_neighbor_x_area"),
							resource.TestCheckResourceAttr(testRoutingOspfStaticNeighbor, "instance_id", "4"),
							resource.TestCheckResourceAttr(testRoutingOspfStaticNeighbor, "poll_interval", "2m"),
						),
					},
					{
						Config: testAccRoutingOspfStaticNeighborUpdatedConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testRoutingOspfStaticNeighbor),
							resource.TestCheckResourceAttr(testRoutingOspfStaticNeighbor, "poll_interval", "3m"),
							resource.TestCheckResourceAttr(testRoutingOspfStaticNeighbor, "comment", "test_ospf_neighbor_x comment"),
							resource.TestCheckResourceAttr(testRoutingOspfStaticNeighbor, "disabled", "true"),
						),
					},
				},
			})

		})
	}
}

func testAccRoutingOspfStaticNeighborConfig() string {
	return providerConfig + `

resource "routeros_routing_ospf_instance" "test_ospf_neighbor_x_inst" {
	name      = "test_ospf_neighbor_x_inst"
	router_id = "192.0.2.110"
}

resource "routeros_routing_ospf_area" "test_ospf_neighbor_x_area" {
	name     = "test_ospf_neighbor_x_area"
	area_id  = "0.0.0.111"
	instance = routeros_routing_ospf_instance.test_ospf_neighbor_x_inst.name
}

resource "routeros_routing_ospf_static_neighbor" "test_ospf_neighbor_x" {
	area          = routeros_routing_ospf_area.test_ospf_neighbor_x_area.name
	address       = "192.0.2.211%ether5"
	instance_id   = 4
	poll_interval = "2m"
}

`
}

func testAccRoutingOspfStaticNeighborUpdatedConfig() string {
	return providerConfig + `

resource "routeros_routing_ospf_instance" "test_ospf_neighbor_x_inst" {
	name      = "test_ospf_neighbor_x_inst"
	router_id = "192.0.2.110"
}

resource "routeros_routing_ospf_area" "test_ospf_neighbor_x_area" {
	name     = "test_ospf_neighbor_x_area"
	area_id  = "0.0.0.111"
	instance = routeros_routing_ospf_instance.test_ospf_neighbor_x_inst.name
}

resource "routeros_routing_ospf_static_neighbor" "test_ospf_neighbor_x" {
	area          = routeros_routing_ospf_area.test_ospf_neighbor_x_area.name
	address       = "192.0.2.211%ether5"
	comment       = "test_ospf_neighbor_x comment"
	disabled      = true
	instance_id   = 4
	poll_interval = "3m"
}

`
}
