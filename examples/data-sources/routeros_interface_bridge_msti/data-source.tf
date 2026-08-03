# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_interface_bridge_msti" "example" {
  filter = {
    comment = "managed by terraform"
  }
}

output "interface_bridge_msti" {
  value = data.routeros_interface_bridge_msti.example.entries[*].id
}
