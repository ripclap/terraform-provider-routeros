package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testRoutingFilterCommunityLargeList = "routeros_routing_filter_community_large_list.test_routing_filter_community_large_list"
const testRoutingFilterCommunityLargeListRe = "routeros_routing_filter_community_large_list.test_routing_filter_community_large_list_re"

func TestAccRoutingFilterCommunityLargeListTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy: testCheckResourceDestroy("/routing/filter/community-large-list",
					"routeros_routing_filter_community_large_list"),
				Steps: []resource.TestStep{
					{
						Config: testAccRoutingFilterCommunityLargeListConfig("acc test large community list"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testRoutingFilterCommunityLargeList),
							resource.TestCheckResourceAttr(testRoutingFilterCommunityLargeList, "list", "test_routing_filter_community_large_list_x"),
							resource.TestCheckResourceAttr(testRoutingFilterCommunityLargeList, "comment", "acc test large community list"),
							resource.TestCheckResourceAttr(testRoutingFilterCommunityLargeList, "communities.#", "2"),
							resource.TestCheckTypeSetElemAttr(testRoutingFilterCommunityLargeList, "communities.*", "65599:1:2"),
							resource.TestCheckTypeSetElemAttr(testRoutingFilterCommunityLargeList, "communities.*", "65599:3:4"),

							testResourcePrimaryInstanceId(testRoutingFilterCommunityLargeListRe),
							resource.TestCheckResourceAttr(testRoutingFilterCommunityLargeListRe, "list", "test_routing_filter_community_large_list_x"),
							resource.TestCheckResourceAttr(testRoutingFilterCommunityLargeListRe, "regexp", "^65599:"),
							resource.TestCheckResourceAttr(testRoutingFilterCommunityLargeListRe, "disabled", "true"),
						),
					},
					{
						Config: testAccRoutingFilterCommunityLargeListConfig("acc test large community list updated"),
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr(testRoutingFilterCommunityLargeList, "comment", "acc test large community list updated"),
						),
					},
				},
			})

		})
	}
}

func testAccRoutingFilterCommunityLargeListConfig(comment string) string {
	return fmt.Sprintf(`%v

resource "routeros_routing_filter_community_large_list" "test_routing_filter_community_large_list" {
  list        = "test_routing_filter_community_large_list_x"
  comment     = "%v"
  communities = ["65599:1:2", "65599:3:4"]
}

resource "routeros_routing_filter_community_large_list" "test_routing_filter_community_large_list_re" {
  list     = "test_routing_filter_community_large_list_x"
  regexp   = "^65599:"
  disabled = true
}
`, providerConfig, comment)
}
