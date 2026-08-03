# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_interface_bridge_vlan" "example" {
  filter = {
    comment = "managed by terraform"
  }
}

output "interface_bridge_vlan" {
  value = data.routeros_interface_bridge_vlan.example.entries[*].id
}
