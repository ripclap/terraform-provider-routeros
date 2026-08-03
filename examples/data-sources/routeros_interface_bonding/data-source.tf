# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_interface_bonding" "example" {
  filter = {
    name = "example"
  }
}

output "interface_bonding" {
  value = data.routeros_interface_bonding.example.entries[*].name
}
