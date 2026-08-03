# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_interface_dot1x_client" "example" {
  filter = {
    interface = "ether1"
  }
}

output "interface_dot1x_client" {
  value = data.routeros_interface_dot1x_client.example.entries[*].id
}
