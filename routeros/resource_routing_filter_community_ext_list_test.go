package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testRoutingFilterCommunityExtList = "routeros_routing_filter_community_ext_list.test_routing_filter_community_ext_list"
const testRoutingFilterCommunityExtListRe = "routeros_routing_filter_community_ext_list.test_routing_filter_community_ext_list_re"

func TestAccRoutingFilterCommunityExtListTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy: testCheckResourceDestroy("/routing/filter/community-ext-list",
					"routeros_routing_filter_community_ext_list"),
				Steps: []resource.TestStep{
					{
						Config: testAccRoutingFilterCommunityExtListConfig("acc test ext community list"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testRoutingFilterCommunityExtList),
							resource.TestCheckResourceAttr(testRoutingFilterCommunityExtList, "list", "test_routing_filter_community_ext_list_x"),
							resource.TestCheckResourceAttr(testRoutingFilterCommunityExtList, "comment", "acc test ext community list"),
							resource.TestCheckResourceAttr(testRoutingFilterCommunityExtList, "communities.#", "2"),
							resource.TestCheckTypeSetElemAttr(testRoutingFilterCommunityExtList, "communities.*", "rt:1.1.1.1:222"),
							resource.TestCheckTypeSetElemAttr(testRoutingFilterCommunityExtList, "communities.*", "soo:65099:3"),

							testResourcePrimaryInstanceId(testRoutingFilterCommunityExtListRe),
							resource.TestCheckResourceAttr(testRoutingFilterCommunityExtListRe, "list", "test_routing_filter_community_ext_list_x"),
							resource.TestCheckResourceAttr(testRoutingFilterCommunityExtListRe, "regexp", "^rt:"),
							resource.TestCheckResourceAttr(testRoutingFilterCommunityExtListRe, "disabled", "true"),
						),
					},
					{
						Config: testAccRoutingFilterCommunityExtListConfig("acc test ext community list updated"),
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr(testRoutingFilterCommunityExtList, "comment", "acc test ext community list updated"),
						),
					},
				},
			})

		})
	}
}

func testAccRoutingFilterCommunityExtListConfig(comment string) string {
	return fmt.Sprintf(`%v

resource "routeros_routing_filter_community_ext_list" "test_routing_filter_community_ext_list" {
  list        = "test_routing_filter_community_ext_list_x"
  comment     = "%v"
  communities = ["rt:1.1.1.1:222", "soo:65099:3"]
}

resource "routeros_routing_filter_community_ext_list" "test_routing_filter_community_ext_list_re" {
  list     = "test_routing_filter_community_ext_list_x"
  regexp   = "^rt:"
  disabled = true
}
`, providerConfig, comment)
}
