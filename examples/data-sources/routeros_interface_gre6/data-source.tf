# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_interface_gre6" "example" {
  filter = {
    name = "example"
  }
}

output "interface_gre6" {
  value = data.routeros_interface_gre6.example.entries[*].name
}
