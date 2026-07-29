package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testRoutingPimsmInstance = "routeros_routing_pimsm_instance.test_pimsm_instance_x"

func TestAccRoutingPimsmInstanceTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/routing/pimsm/instance", "routeros_routing_pimsm_instance"),
				Steps: []resource.TestStep{
					{
						Config: testAccRoutingPimsmInstanceConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testRoutingPimsmInstance),
							resource.TestCheckResourceAttr(testRoutingPimsmInstance, "name", "test_pimsm_instance_x"),
							resource.TestCheckResourceAttr(testRoutingPimsmInstance, "afi", "ip"),
							resource.TestCheckResourceAttr(testRoutingPimsmInstance, "ssm_range", "232.0.0.0/8"),
							resource.TestCheckResourceAttr(testRoutingPimsmInstance, "rp_hash_mask_length", "30"),
							resource.TestCheckResourceAttr(testRoutingPimsmInstance, "switch_to_spt", "true"),
							resource.TestCheckResourceAttr(testRoutingPimsmInstance, "switch_to_spt_bytes", "1000"),
							resource.TestCheckResourceAttr(testRoutingPimsmInstance, "switch_to_spt_interval", "1m"),
							resource.TestCheckResourceAttr(testRoutingPimsmInstance, "bsm_forward_back", "true"),
							resource.TestCheckResourceAttr(testRoutingPimsmInstance, "crp_advertise_contained", "true"),
							resource.TestCheckResourceAttr(testRoutingPimsmInstance, "rp_static_override", "true"),
						),
					},
					{
						Config: testAccRoutingPimsmInstanceUpdatedConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testRoutingPimsmInstance),
							resource.TestCheckResourceAttr(testRoutingPimsmInstance, "switch_to_spt", "false"),
							resource.TestCheckResourceAttr(testRoutingPimsmInstance, "rp_static_override", "false"),
							resource.TestCheckResourceAttr(testRoutingPimsmInstance, "rp_hash_mask_length", "24"),
							resource.TestCheckResourceAttr(testRoutingPimsmInstance, "disabled", "true"),
						),
					},
				},
			})

		})
	}
}

func testAccRoutingPimsmInstanceConfig() string {
	return providerConfig + `

resource "routeros_routing_pimsm_instance" "test_pimsm_instance_x" {
	name                    = "test_pimsm_instance_x"
	afi                     = "ip"
	ssm_range               = "232.0.0.0/8"
	rp_hash_mask_length     = 30
	switch_to_spt           = true
	switch_to_spt_bytes     = 1000
	switch_to_spt_interval  = "1m"
	bsm_forward_back        = true
	crp_advertise_contained = true
	rp_static_override      = true
}

`
}

func testAccRoutingPimsmInstanceUpdatedConfig() string {
	return providerConfig + `

resource "routeros_routing_pimsm_instance" "test_pimsm_instance_x" {
	name                    = "test_pimsm_instance_x"
	afi                     = "ip"
	disabled                = true
	ssm_range               = "232.0.0.0/8"
	rp_hash_mask_length     = 24
	switch_to_spt           = false
	switch_to_spt_bytes     = 1000
	switch_to_spt_interval  = "1m"
	bsm_forward_back        = true
	crp_advertise_contained = true
	rp_static_override      = false
}

`
}
