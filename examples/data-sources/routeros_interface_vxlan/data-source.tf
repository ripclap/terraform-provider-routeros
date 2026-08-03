# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_interface_vxlan" "example" {
  filter = {
    name = "example"
  }
}

output "interface_vxlan" {
  value = data.routeros_interface_vxlan.example.entries[*].name
}
