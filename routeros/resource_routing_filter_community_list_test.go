package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testRoutingFilterCommunityList = "routeros_routing_filter_community_list.test_routing_filter_community_list"
const testRoutingFilterCommunityListRe = "routeros_routing_filter_community_list.test_routing_filter_community_list_re"

func TestAccRoutingFilterCommunityListTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy: testCheckResourceDestroy("/routing/filter/community-list",
					"routeros_routing_filter_community_list"),
				Steps: []resource.TestStep{
					{
						Config: testAccRoutingFilterCommunityListConfig("acc test community list"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testRoutingFilterCommunityList),
							resource.TestCheckResourceAttr(testRoutingFilterCommunityList, "list", "test_routing_filter_community_list_x"),
							resource.TestCheckResourceAttr(testRoutingFilterCommunityList, "comment", "acc test community list"),
							resource.TestCheckResourceAttr(testRoutingFilterCommunityList, "communities.#", "3"),
							resource.TestCheckTypeSetElemAttr(testRoutingFilterCommunityList, "communities.*", "65099:1"),
							resource.TestCheckTypeSetElemAttr(testRoutingFilterCommunityList, "communities.*", "65099:2"),
							resource.TestCheckTypeSetElemAttr(testRoutingFilterCommunityList, "communities.*", "no-export"),

							testResourcePrimaryInstanceId(testRoutingFilterCommunityListRe),
							resource.TestCheckResourceAttr(testRoutingFilterCommunityListRe, "list", "test_routing_filter_community_list_x"),
							resource.TestCheckResourceAttr(testRoutingFilterCommunityListRe, "regexp", "^65099:"),
							resource.TestCheckResourceAttr(testRoutingFilterCommunityListRe, "disabled", "true"),
						),
					},
					{
						Config: testAccRoutingFilterCommunityListConfig("acc test community list updated"),
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr(testRoutingFilterCommunityList, "comment", "acc test community list updated"),
						),
					},
				},
			})

		})
	}
}

func testAccRoutingFilterCommunityListConfig(comment string) string {
	return fmt.Sprintf(`%v

resource "routeros_routing_filter_community_list" "test_routing_filter_community_list" {
  list        = "test_routing_filter_community_list_x"
  comment     = "%v"
  communities = ["65099:1", "65099:2", "no-export"]
}

resource "routeros_routing_filter_community_list" "test_routing_filter_community_list_re" {
  list     = "test_routing_filter_community_list_x"
  regexp   = "^65099:"
  disabled = true
}
`, providerConfig, comment)
}
