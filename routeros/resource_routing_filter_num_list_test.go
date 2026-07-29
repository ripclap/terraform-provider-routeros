package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testRoutingFilterNumList = "routeros_routing_filter_num_list.test_routing_filter_num_list"
const testRoutingFilterNumListSingle = "routeros_routing_filter_num_list.test_routing_filter_num_list_single"

func TestAccRoutingFilterNumListTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy: testCheckResourceDestroy("/routing/filter/num-list",
					"routeros_routing_filter_num_list"),
				Steps: []resource.TestStep{
					{
						Config: testAccRoutingFilterNumListConfig("64512-65534"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testRoutingFilterNumList),
							resource.TestCheckResourceAttr(testRoutingFilterNumList, "list", "test_routing_filter_num_list_x"),
							resource.TestCheckResourceAttr(testRoutingFilterNumList, "range", "64512-65534"),
							resource.TestCheckResourceAttr(testRoutingFilterNumList, "comment", "acc test num list"),

							testResourcePrimaryInstanceId(testRoutingFilterNumListSingle),
							resource.TestCheckResourceAttr(testRoutingFilterNumListSingle, "list", "test_routing_filter_num_list_x"),
							resource.TestCheckResourceAttr(testRoutingFilterNumListSingle, "range", "65535"),
							resource.TestCheckResourceAttr(testRoutingFilterNumListSingle, "disabled", "true"),
						),
					},
					{
						Config: testAccRoutingFilterNumListConfig("4200000000-4294967294"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testRoutingFilterNumList),
							resource.TestCheckResourceAttr(testRoutingFilterNumList, "range", "4200000000-4294967294"),
						),
					},
				},
			})

		})
	}
}

func testAccRoutingFilterNumListConfig(rng string) string {
	return fmt.Sprintf(`%v

resource "routeros_routing_filter_num_list" "test_routing_filter_num_list" {
  list    = "test_routing_filter_num_list_x"
  comment = "acc test num list"
  range   = "%v"
}

resource "routeros_routing_filter_num_list" "test_routing_filter_num_list_single" {
  list     = "test_routing_filter_num_list_x"
  range    = "65535"
  disabled = true
}
`, providerConfig, rng)
}
