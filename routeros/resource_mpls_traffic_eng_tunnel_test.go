package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testMplsTrafficEngTunnelAddress = "routeros_mpls_traffic_eng_tunnel.test_mpls_te_tun_x"

// RouterOS 7.23 accepts but never reports back `setup-priority`, `holding-priority`,
// `record-route`, `affinity-*`, `reoptimize-interval`, `from-address` and `vrf`, so they are
// omitted here; including them would leave every plan dirty.
func TestAccMplsTrafficEngTunnelTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy: testCheckResourceDestroy("/mpls/traffic-eng/tunnel",
					"routeros_mpls_traffic_eng_tunnel"),
				Steps: []resource.TestStep{
					{
						Config: testAccMplsTrafficEngTunnelConfig("198.51.100.9"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testMplsTrafficEngTunnelAddress),
							resource.TestCheckResourceAttr(testMplsTrafficEngTunnelAddress, "name", "test_mpls_te_tun_x"),
							resource.TestCheckResourceAttr(testMplsTrafficEngTunnelAddress, "to_address", "198.51.100.9"),
							resource.TestCheckResourceAttr(testMplsTrafficEngTunnelAddress, "bandwidth", "10000000"),
							resource.TestCheckResourceAttr(testMplsTrafficEngTunnelAddress, "primary_path",
								"test_mpls_te_tun_path_x"),
							resource.TestCheckResourceAttr(testMplsTrafficEngTunnelAddress, "secondary_paths.#", "1"),
							resource.TestCheckResourceAttr(testMplsTrafficEngTunnelAddress, "secondary_paths.0",
								"test_mpls_te_tun_path2_x"),
							resource.TestCheckResourceAttr(testMplsTrafficEngTunnelAddress, "primary_retry_interval", "1m"),
							resource.TestCheckResourceAttr(testMplsTrafficEngTunnelAddress, "secondary_standby", "false"),
							resource.TestCheckResourceAttr(testMplsTrafficEngTunnelAddress, "bandwidth_limit", "disabled"),
							resource.TestCheckResourceAttr(testMplsTrafficEngTunnelAddress, "auto_bandwidth_avg_interval", "5m"),
							resource.TestCheckResourceAttr(testMplsTrafficEngTunnelAddress, "auto_bandwidth_update_interval", "1h"),
						),
					},
					{
						Config: testAccMplsTrafficEngTunnelConfig("198.51.100.10"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testMplsTrafficEngTunnelAddress),
							resource.TestCheckResourceAttr(testMplsTrafficEngTunnelAddress, "to_address", "198.51.100.10"),
						),
					},
				},
			})

		})
	}
}

func testAccMplsTrafficEngTunnelConfig(toAddress string) string {
	return providerConfig + `

resource "routeros_mpls_traffic_eng_path" "test_mpls_te_tun_path_x" {
	name = "test_mpls_te_tun_path_x"
	hops = ["198.51.100.2/strict"]
}

resource "routeros_mpls_traffic_eng_path" "test_mpls_te_tun_path2_x" {
	name = "test_mpls_te_tun_path2_x"
	hops = ["198.51.100.4/loose"]

	# RouterOS closes the REST session when two entries of this menu are created at the same time.
	depends_on = [routeros_mpls_traffic_eng_path.test_mpls_te_tun_path_x]
}

resource "routeros_mpls_traffic_eng_tunnel" "test_mpls_te_tun_x" {
	name                           = "test_mpls_te_tun_x"
	to_address                     = "` + toAddress + `"
	bandwidth                      = "10000000"
	bandwidth_limit                = "disabled"
	primary_path                   = routeros_mpls_traffic_eng_path.test_mpls_te_tun_path_x.name
	secondary_paths                = [routeros_mpls_traffic_eng_path.test_mpls_te_tun_path2_x.name]
	primary_retry_interval         = "1m"
	secondary_standby              = false
	auto_bandwidth_avg_interval    = "5m"
	auto_bandwidth_update_interval = "1h"
	comment                        = "test_mpls_te_tun_x"
}

`
}
