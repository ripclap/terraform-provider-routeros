package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testRoutingPimsmInterfaceTemplate = "routeros_routing_pimsm_interface_template.test_pimsm_iftmpl_x"

func TestAccRoutingPimsmInterfaceTemplateTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy: resource.ComposeAggregateTestCheckFunc(
					testCheckResourceDestroy("/routing/pimsm/interface-template", "routeros_routing_pimsm_interface_template"),
					testCheckResourceDestroy("/routing/pimsm/instance", "routeros_routing_pimsm_instance"),
				),
				Steps: []resource.TestStep{
					{
						Config: testAccRoutingPimsmInterfaceTemplateConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testRoutingPimsmInterfaceTemplate),
							resource.TestCheckResourceAttr(testRoutingPimsmInterfaceTemplate, "instance", "test_pimsm_iftmpl_x_inst"),
							resource.TestCheckResourceAttr(testRoutingPimsmInterfaceTemplate, "interfaces.#", "1"),
							resource.TestCheckTypeSetElemAttr(testRoutingPimsmInterfaceTemplate, "interfaces.*", "ether5"),
							resource.TestCheckResourceAttr(testRoutingPimsmInterfaceTemplate, "hello_delay", "6s"),
							resource.TestCheckResourceAttr(testRoutingPimsmInterfaceTemplate, "hello_period", "31s"),
							resource.TestCheckResourceAttr(testRoutingPimsmInterfaceTemplate, "join_prune_period", "1m1s"),
							resource.TestCheckResourceAttr(testRoutingPimsmInterfaceTemplate, "join_tracking_support", "true"),
							resource.TestCheckResourceAttr(testRoutingPimsmInterfaceTemplate, "override_interval", "2s500ms"),
							resource.TestCheckResourceAttr(testRoutingPimsmInterfaceTemplate, "propagation_delay", "600ms"),
							resource.TestCheckResourceAttr(testRoutingPimsmInterfaceTemplate, "priority", "3"),
						),
					},
					{
						Config: testAccRoutingPimsmInterfaceTemplateUpdatedConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testRoutingPimsmInterfaceTemplate),
							resource.TestCheckResourceAttr(testRoutingPimsmInterfaceTemplate, "interfaces.#", "2"),
							resource.TestCheckTypeSetElemAttr(testRoutingPimsmInterfaceTemplate, "interfaces.*", "ether6"),
							resource.TestCheckResourceAttr(testRoutingPimsmInterfaceTemplate, "join_tracking_support", "false"),
							resource.TestCheckResourceAttr(testRoutingPimsmInterfaceTemplate, "priority", "4"),
							resource.TestCheckResourceAttr(testRoutingPimsmInterfaceTemplate, "disabled", "true"),
						),
					},
				},
			})

		})
	}
}

func testAccRoutingPimsmInterfaceTemplateConfig() string {
	return providerConfig + `

resource "routeros_routing_pimsm_instance" "test_pimsm_iftmpl_x_inst" {
	name = "test_pimsm_iftmpl_x_inst"
	afi  = "ip"
}

resource "routeros_routing_pimsm_interface_template" "test_pimsm_iftmpl_x" {
	instance              = routeros_routing_pimsm_instance.test_pimsm_iftmpl_x_inst.name
	interfaces            = ["ether5"]
	hello_delay           = "6s"
	hello_period          = "31s"
	join_prune_period     = "1m1s"
	join_tracking_support = true
	override_interval     = "2s500ms"
	propagation_delay     = "600ms"
	priority              = 3
}

`
}

func testAccRoutingPimsmInterfaceTemplateUpdatedConfig() string {
	return providerConfig + `

resource "routeros_routing_pimsm_instance" "test_pimsm_iftmpl_x_inst" {
	name = "test_pimsm_iftmpl_x_inst"
	afi  = "ip"
}

resource "routeros_routing_pimsm_interface_template" "test_pimsm_iftmpl_x" {
	instance              = routeros_routing_pimsm_instance.test_pimsm_iftmpl_x_inst.name
	interfaces            = ["ether5", "ether6"]
	disabled              = true
	hello_delay           = "6s"
	hello_period          = "31s"
	join_prune_period     = "1m1s"
	join_tracking_support = false
	override_interval     = "2s500ms"
	propagation_delay     = "600ms"
	priority              = 4
}

`
}
