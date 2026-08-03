# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_interface_ovpn_client" "example" {
  filter = {
    name = "example"
  }
}

output "interface_ovpn_client" {
  value = data.routeros_interface_ovpn_client.example.entries[*].name
}
