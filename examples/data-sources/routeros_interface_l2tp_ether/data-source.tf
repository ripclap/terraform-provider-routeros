# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_interface_l2tp_ether" "example" {
  filter = {
    name = "example"
  }
}

output "interface_l2tp_ether" {
  value = data.routeros_interface_l2tp_ether.example.entries[*].name
}
