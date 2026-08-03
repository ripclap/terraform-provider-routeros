# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_interface_bridge" "example" {
  filter = {
    name = "example"
  }
}

output "interface_bridge" {
  value = data.routeros_interface_bridge.example.entries[*].name
}
