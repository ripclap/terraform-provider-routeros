package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testRoutingOspfInterface = "routeros_routing_ospf_interface.test_ospf_iface_x"

// This menu has no 'comment' property, so the update step is verified through cost, passive and disabled.

func TestAccRoutingOspfInterfaceTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy: resource.ComposeAggregateTestCheckFunc(
					testCheckResourceDestroy("/routing/ospf/interface", "routeros_routing_ospf_interface"),
					testCheckResourceDestroy("/routing/ospf/area", "routeros_routing_ospf_area"),
					testCheckResourceDestroy("/routing/ospf/instance", "routeros_routing_ospf_instance"),
				),
				Steps: []resource.TestStep{
					{
						Config: testAccRoutingOspfInterfaceConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testRoutingOspfInterface),
							resource.TestCheckResourceAttr(testRoutingOspfInterface, "area", "test_ospf_iface_x_area"),
							resource.TestCheckResourceAttr(testRoutingOspfInterface, "interface", "ether5"),
							resource.TestCheckResourceAttr(testRoutingOspfInterface, "cost", "44"),
							resource.TestCheckResourceAttr(testRoutingOspfInterface, "priority", "7"),
							resource.TestCheckResourceAttr(testRoutingOspfInterface, "type", "ptp"),
							resource.TestCheckResourceAttr(testRoutingOspfInterface, "hello_interval", "11s"),
							resource.TestCheckResourceAttr(testRoutingOspfInterface, "dead_interval", "44s"),
							resource.TestCheckResourceAttr(testRoutingOspfInterface, "instance_id", "3"),
							resource.TestCheckResourceAttr(testRoutingOspfInterface, "passive", "false"),
						),
					},
					{
						Config: testAccRoutingOspfInterfaceUpdatedConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testRoutingOspfInterface),
							resource.TestCheckResourceAttr(testRoutingOspfInterface, "cost", "45"),
							resource.TestCheckResourceAttr(testRoutingOspfInterface, "passive", "true"),
							resource.TestCheckResourceAttr(testRoutingOspfInterface, "disabled", "true"),
						),
					},
				},
			})

		})
	}
}

func testAccRoutingOspfInterfaceConfig() string {
	return providerConfig + `

resource "routeros_routing_ospf_instance" "test_ospf_iface_x_inst" {
	name      = "test_ospf_iface_x_inst"
	router_id = "192.0.2.110"
}

resource "routeros_routing_ospf_area" "test_ospf_iface_x_area" {
	name     = "test_ospf_iface_x_area"
	area_id  = "0.0.0.110"
	instance = routeros_routing_ospf_instance.test_ospf_iface_x_inst.name
}

resource "routeros_routing_ospf_interface" "test_ospf_iface_x" {
	area                = routeros_routing_ospf_area.test_ospf_iface_x_area.name
	interface           = "ether5"
	cost                = 44
	priority            = 7
	type                = "ptp"
	hello_interval      = "11s"
	dead_interval       = "44s"
	retransmit_interval = "6s"
	transmit_delay      = "2s"
	instance_id         = 3
}

`
}

func testAccRoutingOspfInterfaceUpdatedConfig() string {
	return providerConfig + `

resource "routeros_routing_ospf_instance" "test_ospf_iface_x_inst" {
	name      = "test_ospf_iface_x_inst"
	router_id = "192.0.2.110"
}

resource "routeros_routing_ospf_area" "test_ospf_iface_x_area" {
	name     = "test_ospf_iface_x_area"
	area_id  = "0.0.0.110"
	instance = routeros_routing_ospf_instance.test_ospf_iface_x_inst.name
}

resource "routeros_routing_ospf_interface" "test_ospf_iface_x" {
	area                = routeros_routing_ospf_area.test_ospf_iface_x_area.name
	interface           = "ether5"
	disabled            = true
	passive             = true
	cost                = 45
	priority            = 7
	type                = "ptp"
	hello_interval      = "11s"
	dead_interval       = "44s"
	retransmit_interval = "6s"
	transmit_delay      = "2s"
	instance_id         = 3
}

`
}
