# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_interface_ovpn_server" "example" {
  filter = {
    name = "example"
  }
}

output "interface_ovpn_server" {
  value = data.routeros_interface_ovpn_server.example.entries[*].name
}
