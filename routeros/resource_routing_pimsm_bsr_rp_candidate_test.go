package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testRoutingPimsmBsrRpCandidate = "routeros_routing_pimsm_bsr_rp_candidate.test_pimsm_bsr_rp_cand_x"

func TestAccRoutingPimsmBsrRpCandidateTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy: resource.ComposeAggregateTestCheckFunc(
					testCheckResourceDestroy("/routing/pimsm/bsr/rp-candidate", "routeros_routing_pimsm_bsr_rp_candidate"),
					testCheckResourceDestroy("/routing/pimsm/instance", "routeros_routing_pimsm_instance"),
				),
				Steps: []resource.TestStep{
					{
						Config: testAccRoutingPimsmBsrRpCandidateConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testRoutingPimsmBsrRpCandidate),
							resource.TestCheckResourceAttr(testRoutingPimsmBsrRpCandidate, "instance", "test_pimsm_bsr_rp_cand_x_inst"),
							resource.TestCheckResourceAttr(testRoutingPimsmBsrRpCandidate, "address", "ether5"),
							resource.TestCheckResourceAttr(testRoutingPimsmBsrRpCandidate, "group", "239.210.0.0/16"),
							resource.TestCheckResourceAttr(testRoutingPimsmBsrRpCandidate, "holdtime", "150"),
							resource.TestCheckResourceAttr(testRoutingPimsmBsrRpCandidate, "priority", "20"),
						),
					},
					{
						Config: testAccRoutingPimsmBsrRpCandidateUpdatedConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testRoutingPimsmBsrRpCandidate),
							resource.TestCheckResourceAttr(testRoutingPimsmBsrRpCandidate, "group", "239.211.0.0/16"),
							resource.TestCheckResourceAttr(testRoutingPimsmBsrRpCandidate, "holdtime", "210"),
							resource.TestCheckResourceAttr(testRoutingPimsmBsrRpCandidate, "priority", "30"),
							resource.TestCheckResourceAttr(testRoutingPimsmBsrRpCandidate, "disabled", "true"),
						),
					},
				},
			})

		})
	}
}

func testAccRoutingPimsmBsrRpCandidateConfig() string {
	return providerConfig + `

resource "routeros_routing_pimsm_instance" "test_pimsm_bsr_rp_cand_x_inst" {
	name = "test_pimsm_bsr_rp_cand_x_inst"
	afi  = "ip"
}

resource "routeros_routing_pimsm_bsr_rp_candidate" "test_pimsm_bsr_rp_cand_x" {
	instance = routeros_routing_pimsm_instance.test_pimsm_bsr_rp_cand_x_inst.name
	address  = "ether5"
	group    = "239.210.0.0/16"
	holdtime = "150"
	priority = 20
}

`
}

func testAccRoutingPimsmBsrRpCandidateUpdatedConfig() string {
	return providerConfig + `

resource "routeros_routing_pimsm_instance" "test_pimsm_bsr_rp_cand_x_inst" {
	name = "test_pimsm_bsr_rp_cand_x_inst"
	afi  = "ip"
}

resource "routeros_routing_pimsm_bsr_rp_candidate" "test_pimsm_bsr_rp_cand_x" {
	instance = routeros_routing_pimsm_instance.test_pimsm_bsr_rp_cand_x_inst.name
	address  = "ether5"
	disabled = true
	group    = "239.211.0.0/16"
	holdtime = "210"
	priority = 30
}

`
}
