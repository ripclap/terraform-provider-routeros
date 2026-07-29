package routeros

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testResourceToolTrafficGeneratorRawPacketTemplate = "routeros_tool_traffic_generator_raw_packet_template.test_tg_raw_tmpl_x"
const testResourceToolTrafficGeneratorRawPacketTemplatePort = "routeros_tool_traffic_generator_port.test_tg_raw_port"

// A raw template binds to a port (not an interface), so the test creates its own port first.
// The always-reported fields are set explicitly to avoid a non-empty plan.
func TestAccToolTrafficGeneratorRawPacketTemplateTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy: testCheckResourceDestroy("/tool/traffic-generator/raw-packet-template",
					"routeros_tool_traffic_generator_raw_packet_template"),
				Steps: []resource.TestStep{
					{
						Config: testAccToolTrafficGeneratorRawPacketTemplatePortConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testResourceToolTrafficGeneratorRawPacketTemplatePort),
						),
					},
					{
						Config: testAccToolTrafficGeneratorRawPacketTemplateConfig("14"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testResourceToolTrafficGeneratorRawPacketTemplate),
							resource.TestCheckResourceAttr(testResourceToolTrafficGeneratorRawPacketTemplate, "name", "test_tg_raw_tmpl_x"),
							resource.TestCheckResourceAttr(testResourceToolTrafficGeneratorRawPacketTemplate, "comment", "test_tg_raw_tmpl_x"),
							resource.TestCheckResourceAttr(testResourceToolTrafficGeneratorRawPacketTemplate, "port", "test_tg_raw_port"),
							resource.TestCheckResourceAttr(testResourceToolTrafficGeneratorRawPacketTemplate, "header", "0000005E005301001122334455660800"),
							resource.TestCheckResourceAttr(testResourceToolTrafficGeneratorRawPacketTemplate, "header_length", "16"),
							resource.TestCheckResourceAttr(testResourceToolTrafficGeneratorRawPacketTemplate, "ip_header_offset", "14"),
							resource.TestCheckResourceAttr(testResourceToolTrafficGeneratorRawPacketTemplate, "data", "random"),
							resource.TestCheckResourceAttr(testResourceToolTrafficGeneratorRawPacketTemplate, "special_footer", "true"),
							resource.TestCheckResourceAttr(testResourceToolTrafficGeneratorRawPacketTemplate, "compute_checksum_from_offset", "no-checksum"),
							resource.TestCheckResourceAttr(testResourceToolTrafficGeneratorRawPacketTemplate, "dynamic", "false"),
						),
					},
					{
						Config: testAccToolTrafficGeneratorRawPacketTemplateConfig("12"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testResourceToolTrafficGeneratorRawPacketTemplate),
							resource.TestCheckResourceAttr(testResourceToolTrafficGeneratorRawPacketTemplate, "ip_header_offset", "12"),
						),
					},
				},
			})
		})
	}
}

func testAccToolTrafficGeneratorRawPacketTemplatePortConfig() string {
	return providerConfig + `
resource "routeros_tool_traffic_generator_port" "test_tg_raw_port" {
	name      = "test_tg_raw_port"
	interface = "ether6"
}
`
}

func testAccToolTrafficGeneratorRawPacketTemplateConfig(ipHeaderOffset string) string {
	return testAccToolTrafficGeneratorRawPacketTemplatePortConfig() + `
resource "routeros_tool_traffic_generator_raw_packet_template" "test_tg_raw_tmpl_x" {
	name                          = "test_tg_raw_tmpl_x"
	comment                       = "test_tg_raw_tmpl_x"
	port                          = routeros_tool_traffic_generator_port.test_tg_raw_port.name
	header                        = "0000005E005301001122334455660800"
	ip_header_offset              = "` + ipHeaderOffset + `"
	data                          = "random"
	data_byte                     = "0"
	random_byte_offsets_and_masks = ""
	random_ranges                 = ""
	special_footer                = true
	compute_checksum_from_offset  = "no-checksum"
}
`
}
