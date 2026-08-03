# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_interface_macvlan" "example" {
  filter = {
    name = "example"
  }
}

output "interface_macvlan" {
  value = data.routeros_interface_macvlan.example.entries[*].name
}
