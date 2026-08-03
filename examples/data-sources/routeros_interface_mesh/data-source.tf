# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_interface_mesh" "example" {
  filter = {
    name = "example"
  }
}

output "interface_mesh" {
  value = data.routeros_interface_mesh.example.entries[*].name
}
