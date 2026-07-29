package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testMplsTrafficEngInterfaceAddress = "routeros_mpls_traffic_eng_interface.test_mpls_te_iface_x"

// The RSVP-TE interface is attached to one of the interface fixtures of the test router.
func TestAccMplsTrafficEngInterfaceTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy: testCheckResourceDestroy("/mpls/traffic-eng/interface",
					"routeros_mpls_traffic_eng_interface"),
				Steps: []resource.TestStep{
					{
						Config: testAccMplsTrafficEngInterfaceConfig("7"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testMplsTrafficEngInterfaceAddress),
							resource.TestCheckResourceAttr(testMplsTrafficEngInterfaceAddress, "interface", "ether6"),
							resource.TestCheckResourceAttr(testMplsTrafficEngInterfaceAddress, "bandwidth", "100000000"),
							resource.TestCheckResourceAttr(testMplsTrafficEngInterfaceAddress, "te_metric", "7"),
							resource.TestCheckResourceAttr(testMplsTrafficEngInterfaceAddress, "k_factor", "3"),
							resource.TestCheckResourceAttr(testMplsTrafficEngInterfaceAddress, "blockade_k_factor", "4"),
							resource.TestCheckResourceAttr(testMplsTrafficEngInterfaceAddress, "refresh_time", "30s"),
							resource.TestCheckResourceAttr(testMplsTrafficEngInterfaceAddress, "igp_flood_period", "3m"),
							resource.TestCheckResourceAttr(testMplsTrafficEngInterfaceAddress, "use_udp", "false"),
							resource.TestCheckResourceAttr(testMplsTrafficEngInterfaceAddress, "up_flood_thresholds.#", "3"),
							resource.TestCheckResourceAttr(testMplsTrafficEngInterfaceAddress, "up_flood_thresholds.0", "15"),
							resource.TestCheckResourceAttr(testMplsTrafficEngInterfaceAddress, "down_flood_thresholds.#", "3"),
							resource.TestCheckResourceAttr(testMplsTrafficEngInterfaceAddress, "down_flood_thresholds.0", "100"),
						),
					},
					{
						Config: testAccMplsTrafficEngInterfaceConfig("11"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testMplsTrafficEngInterfaceAddress),
							resource.TestCheckResourceAttr(testMplsTrafficEngInterfaceAddress, "te_metric", "11"),
						),
					},
				},
			})

		})
	}
}

func testAccMplsTrafficEngInterfaceConfig(teMetric string) string {
	return providerConfig + `

resource "routeros_mpls_traffic_eng_interface" "test_mpls_te_iface_x" {
	interface             = "ether6"
	bandwidth             = "100000000"
	te_metric             = ` + teMetric + `
	k_factor              = 3
	blockade_k_factor     = 4
	refresh_time          = "30s"
	igp_flood_period      = "3m"
	resource_class        = "0x00000003"
	use_udp               = false
	up_flood_thresholds   = ["15", "30", "45"]
	down_flood_thresholds = ["100", "90", "80"]
	comment               = "test_mpls_te_iface_x"
}

`
}
