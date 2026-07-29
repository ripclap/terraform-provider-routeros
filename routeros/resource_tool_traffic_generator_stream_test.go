package routeros

import (
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testResourceToolTrafficGeneratorStream = "routeros_tool_traffic_generator_stream.test_tg_stream_x"
const testResourceToolTrafficGeneratorStreamPort = "routeros_tool_traffic_generator_port.test_tg_stream_port"
const testResourceToolTrafficGeneratorStreamTemplate = "routeros_tool_traffic_generator_packet_template.test_tg_stream_tmpl"

// A stream needs a port and a packet template, created in a separate step first.
// cpu_core is set explicitly to avoid a non-empty plan; default_port is only checked for presence.
func TestAccToolTrafficGeneratorStreamTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy: testCheckResourceDestroy("/tool/traffic-generator/stream",
					"routeros_tool_traffic_generator_stream"),
				Steps: []resource.TestStep{
					{
						Config: testAccToolTrafficGeneratorStreamDepsConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testResourceToolTrafficGeneratorStreamPort),
							testResourcePrimaryInstanceId(testResourceToolTrafficGeneratorStreamTemplate),
						),
					},
					{
						Config: testAccToolTrafficGeneratorStreamConfig(100),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testResourceToolTrafficGeneratorStream),
							resource.TestCheckResourceAttr(testResourceToolTrafficGeneratorStream, "name", "test_tg_stream_x"),
							resource.TestCheckResourceAttr(testResourceToolTrafficGeneratorStream, "port", "test_tg_stream_port"),
							resource.TestCheckResourceAttr(testResourceToolTrafficGeneratorStream, "tx_template", "test_tg_stream_tmpl"),
							resource.TestCheckResourceAttr(testResourceToolTrafficGeneratorStream, "stream_id", "12"),
							resource.TestCheckResourceAttr(testResourceToolTrafficGeneratorStream, "pps", "100"),
							resource.TestCheckResourceAttr(testResourceToolTrafficGeneratorStream, "packet_size", "128"),
							resource.TestCheckResourceAttr(testResourceToolTrafficGeneratorStream, "packet_count", "unlimited"),
							resource.TestCheckResourceAttr(testResourceToolTrafficGeneratorStream, "cpu_core", "0"),
							resource.TestCheckResourceAttr(testResourceToolTrafficGeneratorStream, "disabled", "false"),
							resource.TestCheckResourceAttr(testResourceToolTrafficGeneratorStream, "invalid", "false"),
							resource.TestCheckResourceAttrSet(testResourceToolTrafficGeneratorStream, "default_port"),
						),
					},
					{
						Config: testAccToolTrafficGeneratorStreamConfig(200),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testResourceToolTrafficGeneratorStream),
							resource.TestCheckResourceAttr(testResourceToolTrafficGeneratorStream, "pps", "200"),
						),
					},
				},
			})
		})
	}
}

func testAccToolTrafficGeneratorStreamDepsConfig() string {
	return providerConfig + `
resource "routeros_tool_traffic_generator_port" "test_tg_stream_port" {
	name      = "test_tg_stream_port"
	interface = "ether6"
}

resource "routeros_tool_traffic_generator_packet_template" "test_tg_stream_tmpl" {
	name                          = "test_tg_stream_tmpl"
	interface                     = "ether6"
	header_stack                  = "mac,ip,udp"
	mac_dst                       = "00:00:5E:00:53:01/FF:FF:FF:FF:FF:FF"
	ip_src                        = "198.51.100.1"
	ip_dst                        = "198.51.100.2"
	udp_src_port                  = "5000"
	udp_dst_port                  = "5001"
	data                          = "random"
	data_byte                     = "0"
	random_byte_offsets_and_masks = ""
	random_ranges                 = ""
	special_footer                = true
	compute_checksum_from_offset  = "no-checksum"
}
`
}

func testAccToolTrafficGeneratorStreamConfig(pps int) string {
	return testAccToolTrafficGeneratorStreamDepsConfig() + `
resource "routeros_tool_traffic_generator_stream" "test_tg_stream_x" {
	name         = "test_tg_stream_x"
	port         = routeros_tool_traffic_generator_port.test_tg_stream_port.name
	tx_template  = routeros_tool_traffic_generator_packet_template.test_tg_stream_tmpl.name
	stream_id    = 12
	pps          = ` + strconv.Itoa(pps) + `
	packet_size  = "128"
	packet_count = "unlimited"
	cpu_core     = "0"
	disabled     = false
}
`
}
