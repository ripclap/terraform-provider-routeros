package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testRoutingRipInstance = "routeros_routing_rip_instance.test_rip_instance_x"

func TestAccRoutingRipInstanceTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/routing/rip/instance", "routeros_routing_rip_instance"),
				Steps: []resource.TestStep{
					{
						Config: testAccRoutingRipInstanceConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testRoutingRipInstance),
							resource.TestCheckResourceAttr(testRoutingRipInstance, "name", "test_rip_instance_x"),
							resource.TestCheckResourceAttr(testRoutingRipInstance, "afi", "ip"),
							resource.TestCheckResourceAttr(testRoutingRipInstance, "originate_default", "never"),
							resource.TestCheckResourceAttr(testRoutingRipInstance, "route_timeout", "180"),
							resource.TestCheckResourceAttr(testRoutingRipInstance, "route_gc_timeout", "120"),
							resource.TestCheckResourceAttr(testRoutingRipInstance, "update_interval", "45"),
							resource.TestCheckResourceAttr(testRoutingRipInstance, "routing_table", "main"),
							resource.TestCheckResourceAttr(testRoutingRipInstance, "redistribute.#", "2"),
							resource.TestCheckTypeSetElemAttr(testRoutingRipInstance, "redistribute.*", "connected"),
							resource.TestCheckTypeSetElemAttr(testRoutingRipInstance, "redistribute.*", "static"),
						),
					},
					{
						Config: testAccRoutingRipInstanceUpdatedConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testRoutingRipInstance),
							resource.TestCheckResourceAttr(testRoutingRipInstance, "originate_default", "if-installed"),
							resource.TestCheckResourceAttr(testRoutingRipInstance, "update_interval", "60"),
							resource.TestCheckResourceAttr(testRoutingRipInstance, "redistribute.#", "1"),
							resource.TestCheckTypeSetElemAttr(testRoutingRipInstance, "redistribute.*", "connected"),
							resource.TestCheckResourceAttr(testRoutingRipInstance, "disabled", "true"),
						),
					},
				},
			})

		})
	}
}

func testAccRoutingRipInstanceConfig() string {
	return providerConfig + `

resource "routeros_routing_rip_instance" "test_rip_instance_x" {
	name              = "test_rip_instance_x"
	afi               = "ip"
	originate_default = "never"
	redistribute      = ["connected", "static"]
	route_timeout     = "180"
	route_gc_timeout  = "120"
	update_interval   = "45"
	routing_table     = "main"
}

`
}

func testAccRoutingRipInstanceUpdatedConfig() string {
	return providerConfig + `

resource "routeros_routing_rip_instance" "test_rip_instance_x" {
	name              = "test_rip_instance_x"
	afi               = "ip"
	disabled          = true
	originate_default = "if-installed"
	redistribute      = ["connected"]
	route_timeout     = "180"
	route_gc_timeout  = "120"
	update_interval   = "60"
	routing_table     = "main"
}

`
}
