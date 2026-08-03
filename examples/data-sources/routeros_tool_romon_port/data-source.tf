# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_tool_romon_port" "example" {
  filter = {
    interface = "ether1"
  }
}

output "tool_romon_port" {
  value = data.routeros_tool_romon_port.example.entries[*].id
}
