# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_tool_traffic_monitor" "example" {
  filter = {
    name = "example"
  }
}

output "tool_traffic_monitor" {
  value = data.routeros_tool_traffic_monitor.example.entries[*].name
}
