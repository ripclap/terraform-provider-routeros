# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_interface_bridge_mdb" "example" {
  filter = {
    comment = "managed by terraform"
  }
}

output "interface_bridge_mdb" {
  value = data.routeros_interface_bridge_mdb.example.entries[*].id
}
