package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testRoutingPimsmBsrCandidate = "routeros_routing_pimsm_bsr_candidate.test_pimsm_bsr_cand_x"

func TestAccRoutingPimsmBsrCandidateTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy: resource.ComposeAggregateTestCheckFunc(
					testCheckResourceDestroy("/routing/pimsm/bsr/candidate", "routeros_routing_pimsm_bsr_candidate"),
					testCheckResourceDestroy("/routing/pimsm/instance", "routeros_routing_pimsm_instance"),
				),
				Steps: []resource.TestStep{
					{
						Config: testAccRoutingPimsmBsrCandidateConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testRoutingPimsmBsrCandidate),
							resource.TestCheckResourceAttr(testRoutingPimsmBsrCandidate, "instance", "test_pimsm_bsr_cand_x_inst"),
							resource.TestCheckResourceAttr(testRoutingPimsmBsrCandidate, "address", "ether5"),
							resource.TestCheckResourceAttr(testRoutingPimsmBsrCandidate, "hashmask_length", "30"),
							resource.TestCheckResourceAttr(testRoutingPimsmBsrCandidate, "priority", "100"),
						),
					},
					{
						Config: testAccRoutingPimsmBsrCandidateUpdatedConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testRoutingPimsmBsrCandidate),
							resource.TestCheckResourceAttr(testRoutingPimsmBsrCandidate, "address", "ether6"),
							resource.TestCheckResourceAttr(testRoutingPimsmBsrCandidate, "hashmask_length", "24"),
							resource.TestCheckResourceAttr(testRoutingPimsmBsrCandidate, "priority", "110"),
							resource.TestCheckResourceAttr(testRoutingPimsmBsrCandidate, "disabled", "true"),
						),
					},
				},
			})

		})
	}
}

func testAccRoutingPimsmBsrCandidateConfig() string {
	return providerConfig + `

resource "routeros_routing_pimsm_instance" "test_pimsm_bsr_cand_x_inst" {
	name = "test_pimsm_bsr_cand_x_inst"
	afi  = "ip"
}

resource "routeros_routing_pimsm_bsr_candidate" "test_pimsm_bsr_cand_x" {
	instance        = routeros_routing_pimsm_instance.test_pimsm_bsr_cand_x_inst.name
	address         = "ether5"
	hashmask_length = 30
	priority        = 100
}

`
}

func testAccRoutingPimsmBsrCandidateUpdatedConfig() string {
	return providerConfig + `

resource "routeros_routing_pimsm_instance" "test_pimsm_bsr_cand_x_inst" {
	name = "test_pimsm_bsr_cand_x_inst"
	afi  = "ip"
}

resource "routeros_routing_pimsm_bsr_candidate" "test_pimsm_bsr_cand_x" {
	instance        = routeros_routing_pimsm_instance.test_pimsm_bsr_cand_x_inst.name
	address         = "ether6"
	disabled        = true
	hashmask_length = 24
	priority        = 110
}

`
}
