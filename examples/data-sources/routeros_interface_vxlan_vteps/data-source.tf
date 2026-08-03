# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_interface_vxlan_vteps" "example" {
  filter = {
    interface = "ether1"
  }
}

output "interface_vxlan_vteps" {
  value = data.routeros_interface_vxlan_vteps.example.entries[*].id
}
