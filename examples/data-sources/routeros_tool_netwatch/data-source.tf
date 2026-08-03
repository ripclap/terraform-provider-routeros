# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_tool_netwatch" "example" {
  filter = {
    name = "example"
  }
}

output "tool_netwatch" {
  value = data.routeros_tool_netwatch.example.entries[*].name
}
