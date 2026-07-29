package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testMplsTrafficEngPathAddress = "routeros_mpls_traffic_eng_path.test_mpls_te_path_x"

func TestAccMplsTrafficEngPathTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy: testCheckResourceDestroy("/mpls/traffic-eng/path",
					"routeros_mpls_traffic_eng_path"),
				Steps: []resource.TestStep{
					{
						Config: testAccMplsTrafficEngPathConfig("6"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testMplsTrafficEngPathAddress),
							resource.TestCheckResourceAttr(testMplsTrafficEngPathAddress, "name", "test_mpls_te_path_x"),
							resource.TestCheckResourceAttr(testMplsTrafficEngPathAddress, "hops.#", "2"),
							resource.TestCheckResourceAttr(testMplsTrafficEngPathAddress, "hops.0", "198.51.100.2/strict"),
							resource.TestCheckResourceAttr(testMplsTrafficEngPathAddress, "hops.1", "198.51.100.3/loose"),
							resource.TestCheckResourceAttr(testMplsTrafficEngPathAddress, "use_cspf", "true"),
							resource.TestCheckResourceAttr(testMplsTrafficEngPathAddress, "record_route", "true"),
							resource.TestCheckResourceAttr(testMplsTrafficEngPathAddress, "setup_priority", "6"),
							resource.TestCheckResourceAttr(testMplsTrafficEngPathAddress, "holding_priority", "6"),
							resource.TestCheckResourceAttr(testMplsTrafficEngPathAddress, "reoptimize_interval", "5m"),
						),
					},
					{
						Config: testAccMplsTrafficEngPathConfig("4"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testMplsTrafficEngPathAddress),
							resource.TestCheckResourceAttr(testMplsTrafficEngPathAddress, "setup_priority", "4"),
						),
					},
				},
			})

		})
	}
}

func testAccMplsTrafficEngPathConfig(setupPriority string) string {
	return providerConfig + `

resource "routeros_mpls_traffic_eng_path" "test_mpls_te_path_x" {
	name                 = "test_mpls_te_path_x"
	hops                 = ["198.51.100.2/strict", "198.51.100.3/loose"]
	use_cspf             = true
	record_route         = true
	setup_priority       = ` + setupPriority + `
	holding_priority     = 6
	affinity_include_any = "0x00000003"
	affinity_include_all = "0x00000001"
	affinity_exclude     = "0x00000008"
	reoptimize_interval  = "5m"
	comment              = "test_mpls_te_path_x"
}

`
}
