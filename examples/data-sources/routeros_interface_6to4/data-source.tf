# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_interface_6to4" "example" {
  filter = {
    name = "example"
  }
}

output "interface_6to4" {
  value = data.routeros_interface_6to4.example.entries[*].name
}
