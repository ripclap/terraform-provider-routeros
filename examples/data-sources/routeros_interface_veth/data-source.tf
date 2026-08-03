# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_interface_veth" "example" {
  filter = {
    name = "example"
  }
}

output "interface_veth" {
  value = data.routeros_interface_veth.example.entries[*].name
}
