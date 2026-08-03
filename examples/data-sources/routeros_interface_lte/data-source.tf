# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_interface_lte" "example" {
  filter = {
    name = "example"
  }
}

output "interface_lte" {
  value = data.routeros_interface_lte.example.entries[*].name
}
