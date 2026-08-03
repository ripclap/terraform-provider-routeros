# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_tool_graphing_interface" "example" {
  filter = {
    interface = "ether1"
  }
}

output "tool_graphing_interface" {
  value = data.routeros_tool_graphing_interface.example.entries[*].id
}
