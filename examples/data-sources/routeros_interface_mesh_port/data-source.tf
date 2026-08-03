# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_interface_mesh_port" "example" {
  filter = {
    interface = "ether1"
  }
}

output "interface_mesh_port" {
  value = data.routeros_interface_mesh_port.example.entries[*].id
}
