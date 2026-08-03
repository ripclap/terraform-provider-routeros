# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_interface_l2tp_server_binding" "example" {
  filter = {
    name = "example"
  }
}

output "interface_l2tp_server_binding" {
  value = data.routeros_interface_l2tp_server_binding.example.entries[*].name
}
