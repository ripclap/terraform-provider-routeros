package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testRoutingRipInterfaceTemplate = "routeros_routing_rip_interface_template.test_rip_iftmpl_x"

func TestAccRoutingRipInterfaceTemplateTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy: resource.ComposeAggregateTestCheckFunc(
					testCheckResourceDestroy("/routing/rip/interface-template", "routeros_routing_rip_interface_template"),
					testCheckResourceDestroy("/routing/rip/keys", "routeros_routing_rip_keys"),
					testCheckResourceDestroy("/routing/rip/instance", "routeros_routing_rip_instance"),
				),
				Steps: []resource.TestStep{
					{
						Config: testAccRoutingRipInterfaceTemplateConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testRoutingRipInterfaceTemplate),
							resource.TestCheckResourceAttr(testRoutingRipInterfaceTemplate, "instance", "test_rip_iftmpl_x_inst"),
							resource.TestCheckResourceAttr(testRoutingRipInterfaceTemplate, "interfaces.#", "1"),
							resource.TestCheckTypeSetElemAttr(testRoutingRipInterfaceTemplate, "interfaces.*", "ether5"),
							resource.TestCheckResourceAttr(testRoutingRipInterfaceTemplate, "cost", "3"),
							resource.TestCheckResourceAttr(testRoutingRipInterfaceTemplate, "mode", "passive"),
							resource.TestCheckResourceAttr(testRoutingRipInterfaceTemplate, "key_chain", "test_rip_iftmpl_x_chain"),
							resource.TestCheckResourceAttr(testRoutingRipInterfaceTemplate, "poison_reverse", "true"),
							resource.TestCheckResourceAttr(testRoutingRipInterfaceTemplate, "split_horizon", "true"),
							resource.TestCheckResourceAttr(testRoutingRipInterfaceTemplate, "use_bfd", "false"),
						),
					},
					{
						Config: testAccRoutingRipInterfaceTemplateUpdatedConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testRoutingRipInterfaceTemplate),
							resource.TestCheckResourceAttr(testRoutingRipInterfaceTemplate, "interfaces.#", "2"),
							resource.TestCheckTypeSetElemAttr(testRoutingRipInterfaceTemplate, "interfaces.*", "ether6"),
							resource.TestCheckResourceAttr(testRoutingRipInterfaceTemplate, "cost", "4"),
							resource.TestCheckResourceAttr(testRoutingRipInterfaceTemplate, "poison_reverse", "false"),
							resource.TestCheckResourceAttr(testRoutingRipInterfaceTemplate, "disabled", "true"),
						),
					},
				},
			})

		})
	}
}

func testAccRoutingRipInterfaceTemplateConfig() string {
	return providerConfig + `

resource "routeros_routing_rip_instance" "test_rip_iftmpl_x_inst" {
	name = "test_rip_iftmpl_x_inst"
	afi  = "ip"
}

resource "routeros_routing_rip_keys" "test_rip_iftmpl_x_key" {
	chain  = "test_rip_iftmpl_x_chain"
	key    = "riptmplx_secret"
	key_id = 3
}

resource "routeros_routing_rip_interface_template" "test_rip_iftmpl_x" {
	instance       = routeros_routing_rip_instance.test_rip_iftmpl_x_inst.name
	interfaces     = ["ether5"]
	cost           = 3
	mode           = "passive"
	key_chain      = routeros_routing_rip_keys.test_rip_iftmpl_x_key.chain
	poison_reverse = true
	split_horizon  = true
	use_bfd        = false
}

`
}

func testAccRoutingRipInterfaceTemplateUpdatedConfig() string {
	return providerConfig + `

resource "routeros_routing_rip_instance" "test_rip_iftmpl_x_inst" {
	name = "test_rip_iftmpl_x_inst"
	afi  = "ip"
}

resource "routeros_routing_rip_keys" "test_rip_iftmpl_x_key" {
	chain  = "test_rip_iftmpl_x_chain"
	key    = "riptmplx_secret"
	key_id = 3
}

resource "routeros_routing_rip_interface_template" "test_rip_iftmpl_x" {
	instance       = routeros_routing_rip_instance.test_rip_iftmpl_x_inst.name
	interfaces     = ["ether5", "ether6"]
	disabled       = true
	cost           = 4
	mode           = "passive"
	key_chain      = routeros_routing_rip_keys.test_rip_iftmpl_x_key.chain
	poison_reverse = false
	split_horizon  = true
	use_bfd        = false
}

`
}
