# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_bridge_vlan" "example" {
  filter = {
    comment = "managed by terraform"
  }
}

output "bridge_vlan" {
  value = data.routeros_bridge_vlan.example.entries[*].id
}
