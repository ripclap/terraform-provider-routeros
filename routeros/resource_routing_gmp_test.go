package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testRoutingGmp = "routeros_routing_gmp.test_routing_gmp"
const testRoutingGmpExclude = "routeros_routing_gmp.test_routing_gmp_exclude"

// 'exclude' is a RouterOS flag property: the console takes it without a value and the REST output
// reports it as an empty string, so it is exercised in both states here.
// The menu has no 'comment' property, so the update step changes the source list instead.
func TestAccRoutingGmpTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/routing/gmp", "routeros_routing_gmp"),
				Steps: []resource.TestStep{
					{
						Config: testAccRoutingGmpConfig("192.0.2.1"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testRoutingGmp),
							resource.TestCheckResourceAttr(testRoutingGmp, "sources.#", "1"),
							resource.TestCheckResourceAttr(testRoutingGmp, "exclude", "false"),
							resource.TestCheckResourceAttr(testRoutingGmp, "groups.#", "2"),
							resource.TestCheckTypeSetElemAttr(testRoutingGmp, "groups.*", "239.9.9.1"),
							resource.TestCheckTypeSetElemAttr(testRoutingGmp, "groups.*", "239.9.9.2"),
							resource.TestCheckTypeSetElemAttr(testRoutingGmp, "interfaces.*", "ether5"),
							resource.TestCheckTypeSetElemAttr(testRoutingGmp, "sources.*", "192.0.2.1"),

							testResourcePrimaryInstanceId(testRoutingGmpExclude),
							resource.TestCheckResourceAttr(testRoutingGmpExclude, "exclude", "true"),
							resource.TestCheckResourceAttr(testRoutingGmpExclude, "disabled", "true"),
							resource.TestCheckTypeSetElemAttr(testRoutingGmpExclude, "groups.*", "239.9.9.3"),
							resource.TestCheckTypeSetElemAttr(testRoutingGmpExclude, "sources.*", "192.0.2.2"),
						),
					},
					{
						Config: testAccRoutingGmpConfig("192.0.2.11"),
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckTypeSetElemAttr(testRoutingGmp, "sources.*", "192.0.2.11"),
						),
					},
				},
			})

		})
	}
}

func testAccRoutingGmpConfig(source string) string {
	return fmt.Sprintf(`%v

resource "routeros_routing_gmp" "test_routing_gmp" {
  groups     = ["239.9.9.1", "239.9.9.2"]
  interfaces = ["ether5"]
  sources    = ["%v"]
}

resource "routeros_routing_gmp" "test_routing_gmp_exclude" {
  disabled   = true
  exclude    = true
  groups     = ["239.9.9.3"]
  interfaces = ["ether6"]
  sources    = ["192.0.2.2"]

  depends_on = [routeros_routing_gmp.test_routing_gmp]
}
`, providerConfig, source)
}
