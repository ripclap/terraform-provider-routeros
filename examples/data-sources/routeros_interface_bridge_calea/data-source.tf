# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_interface_bridge_calea" "example" {
  filter = {
    chain = "forward"
  }
}

output "interface_bridge_calea" {
  value = data.routeros_interface_bridge_calea.example.entries[*].id
}
