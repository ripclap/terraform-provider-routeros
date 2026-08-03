# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_interface_wireguard_peer" "example" {
  filter = {
    name = "example"
  }
}

output "interface_wireguard_peer" {
  value = data.routeros_interface_wireguard_peer.example.entries[*].name
}
