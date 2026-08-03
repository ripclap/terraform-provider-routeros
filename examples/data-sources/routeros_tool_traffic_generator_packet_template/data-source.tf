# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_tool_traffic_generator_packet_template" "example" {
  filter = {
    name = "example"
  }
}

output "tool_traffic_generator_packet_template" {
  value = data.routeros_tool_traffic_generator_packet_template.example.entries[*].name
}
