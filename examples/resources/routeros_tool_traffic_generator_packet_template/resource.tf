resource "routeros_tool_traffic_generator_packet_template" "template" {
  name                         = "example"
  comment                      = "Managed by OpenTofu"
  compute_checksum_from_offset = "example"
}
