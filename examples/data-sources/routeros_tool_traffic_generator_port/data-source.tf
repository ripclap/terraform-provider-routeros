# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_tool_traffic_generator_port" "example" {
  filter = {
    name = "example"
  }
}

output "tool_traffic_generator_port" {
  value = data.routeros_tool_traffic_generator_port.example.entries[*].name
}
