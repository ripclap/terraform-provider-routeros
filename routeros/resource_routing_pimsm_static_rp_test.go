package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testRoutingPimsmStaticRp = "routeros_routing_pimsm_static_rp.test_pimsm_static_rp_x"

func TestAccRoutingPimsmStaticRpTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy: resource.ComposeAggregateTestCheckFunc(
					testCheckResourceDestroy("/routing/pimsm/static-rp", "routeros_routing_pimsm_static_rp"),
					testCheckResourceDestroy("/routing/pimsm/instance", "routeros_routing_pimsm_instance"),
				),
				Steps: []resource.TestStep{
					{
						Config: testAccRoutingPimsmStaticRpConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testRoutingPimsmStaticRp),
							resource.TestCheckResourceAttr(testRoutingPimsmStaticRp, "instance", "test_pimsm_static_rp_x_inst"),
							resource.TestCheckResourceAttr(testRoutingPimsmStaticRp, "address", "192.0.2.1"),
							resource.TestCheckResourceAttr(testRoutingPimsmStaticRp, "group", "239.212.0.0/16"),
						),
					},
					{
						Config: testAccRoutingPimsmStaticRpUpdatedConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testRoutingPimsmStaticRp),
							resource.TestCheckResourceAttr(testRoutingPimsmStaticRp, "address", "192.0.2.2"),
							resource.TestCheckResourceAttr(testRoutingPimsmStaticRp, "group", "239.213.0.0/16"),
							resource.TestCheckResourceAttr(testRoutingPimsmStaticRp, "comment", "test_pimsm_static_rp_x comment"),
							resource.TestCheckResourceAttr(testRoutingPimsmStaticRp, "disabled", "true"),
						),
					},
				},
			})

		})
	}
}

func testAccRoutingPimsmStaticRpConfig() string {
	return providerConfig + `

resource "routeros_routing_pimsm_instance" "test_pimsm_static_rp_x_inst" {
	name = "test_pimsm_static_rp_x_inst"
	afi  = "ip"
}

resource "routeros_routing_pimsm_static_rp" "test_pimsm_static_rp_x" {
	instance = routeros_routing_pimsm_instance.test_pimsm_static_rp_x_inst.name
	address  = "192.0.2.1"
	group    = "239.212.0.0/16"
}

`
}

func testAccRoutingPimsmStaticRpUpdatedConfig() string {
	return providerConfig + `

resource "routeros_routing_pimsm_instance" "test_pimsm_static_rp_x_inst" {
	name = "test_pimsm_static_rp_x_inst"
	afi  = "ip"
}

resource "routeros_routing_pimsm_static_rp" "test_pimsm_static_rp_x" {
	instance = routeros_routing_pimsm_instance.test_pimsm_static_rp_x_inst.name
	address  = "192.0.2.2"
	group    = "239.213.0.0/16"
	comment  = "test_pimsm_static_rp_x comment"
	disabled = true
}

`
}
