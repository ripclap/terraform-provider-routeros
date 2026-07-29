package routeros

import (
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testResourceToolTrafficGenerator = "routeros_tool_traffic_generator.test_tool_traffic_generator"

// /tool/traffic-generator is a settings singleton, so the test uses no CheckDestroy.
// The second step restores factory values to leave the shared test device untouched.
func TestAccToolTrafficGeneratorTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				Steps: []resource.TestStep{
					{
						Config: testAccToolTrafficGeneratorConfig("200us", true, 200, 12),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testResourceToolTrafficGenerator),
							resource.TestCheckResourceAttr(testResourceToolTrafficGenerator, "latency_distribution_max", "200us"),
							resource.TestCheckResourceAttr(testResourceToolTrafficGenerator, "measure_out_of_order", "true"),
							resource.TestCheckResourceAttr(testResourceToolTrafficGenerator, "stats_samples_to_keep", "200"),
							resource.TestCheckResourceAttr(testResourceToolTrafficGenerator, "test_id", "12"),
						),
					},
					{
						Config: testAccToolTrafficGeneratorConfig("100us", false, 100, 0),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testResourceToolTrafficGenerator),
							resource.TestCheckResourceAttr(testResourceToolTrafficGenerator, "latency_distribution_max", "100us"),
							resource.TestCheckResourceAttr(testResourceToolTrafficGenerator, "measure_out_of_order", "false"),
							resource.TestCheckResourceAttr(testResourceToolTrafficGenerator, "stats_samples_to_keep", "100"),
							resource.TestCheckResourceAttr(testResourceToolTrafficGenerator, "test_id", "0"),
						),
					},
				},
			})
		})
	}
}

func testAccToolTrafficGeneratorConfig(latencyMax string, outOfOrder bool, samples, testId int) string {
	oooState := "false"
	if outOfOrder {
		oooState = "true"
	}
	return providerConfig + `
resource "routeros_tool_traffic_generator" "test_tool_traffic_generator" {
	latency_distribution_max = "` + latencyMax + `"
	measure_out_of_order     = ` + oooState + `
	stats_samples_to_keep    = ` + strconv.Itoa(samples) + `
	test_id                  = ` + strconv.Itoa(testId) + `
}
`
}
