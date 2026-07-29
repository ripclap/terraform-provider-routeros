package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testResourceToolTrafficGeneratorPacketTemplate = "routeros_tool_traffic_generator_packet_template.test_tg_pkt_tmpl_x"

// `mac_dst` carries its mask because RouterOS reports it in `MAC/MASK` form; `compute_checksum_from_offset`,
// `data_byte`, `random_byte_offsets_and_masks` and `random_ranges` are set explicitly or the plan is non-empty.
func TestAccToolTrafficGeneratorPacketTemplateTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy: testCheckResourceDestroy("/tool/traffic-generator/packet-template",
					"routeros_tool_traffic_generator_packet_template"),
				Steps: []resource.TestStep{
					{
						Config: testAccToolTrafficGeneratorPacketTemplateConfig("5001"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testResourceToolTrafficGeneratorPacketTemplate),
							resource.TestCheckResourceAttr(testResourceToolTrafficGeneratorPacketTemplate, "name", "test_tg_pkt_tmpl_x"),
							resource.TestCheckResourceAttr(testResourceToolTrafficGeneratorPacketTemplate, "comment", "test_tg_pkt_tmpl_x"),
							resource.TestCheckResourceAttr(testResourceToolTrafficGeneratorPacketTemplate, "interface", "ether6"),
							resource.TestCheckResourceAttr(testResourceToolTrafficGeneratorPacketTemplate, "header_stack", "mac,ip,udp"),
							resource.TestCheckResourceAttr(testResourceToolTrafficGeneratorPacketTemplate, "mac_dst", "00:00:5E:00:53:01/FF:FF:FF:FF:FF:FF"),
							resource.TestCheckResourceAttr(testResourceToolTrafficGeneratorPacketTemplate, "ip_src", "198.51.100.1"),
							resource.TestCheckResourceAttr(testResourceToolTrafficGeneratorPacketTemplate, "ip_dst", "198.51.100.2"),
							resource.TestCheckResourceAttr(testResourceToolTrafficGeneratorPacketTemplate, "udp_src_port", "5000"),
							resource.TestCheckResourceAttr(testResourceToolTrafficGeneratorPacketTemplate, "udp_dst_port", "5001"),
							resource.TestCheckResourceAttr(testResourceToolTrafficGeneratorPacketTemplate, "data", "random"),
							resource.TestCheckResourceAttr(testResourceToolTrafficGeneratorPacketTemplate, "special_footer", "true"),
							resource.TestCheckResourceAttr(testResourceToolTrafficGeneratorPacketTemplate, "compute_checksum_from_offset", "no-checksum"),
						),
					},
					{
						Config: testAccToolTrafficGeneratorPacketTemplateConfig("5002"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testResourceToolTrafficGeneratorPacketTemplate),
							resource.TestCheckResourceAttr(testResourceToolTrafficGeneratorPacketTemplate, "udp_dst_port", "5002"),
						),
					},
				},
			})
		})
	}
}

func testAccToolTrafficGeneratorPacketTemplateConfig(udpDstPort string) string {
	return providerConfig + `
resource "routeros_tool_traffic_generator_packet_template" "test_tg_pkt_tmpl_x" {
	name                          = "test_tg_pkt_tmpl_x"
	comment                       = "test_tg_pkt_tmpl_x"
	interface                     = "ether6"
	header_stack                  = "mac,ip,udp"
	mac_dst                       = "00:00:5E:00:53:01/FF:FF:FF:FF:FF:FF"
	ip_src                        = "198.51.100.1"
	ip_dst                        = "198.51.100.2"
	udp_src_port                  = "5000"
	udp_dst_port                  = "` + udpDstPort + `"
	data                          = "random"
	data_byte                     = "0"
	random_byte_offsets_and_masks = ""
	random_ranges                 = ""
	special_footer                = true
	compute_checksum_from_offset  = "no-checksum"
}
`
}
