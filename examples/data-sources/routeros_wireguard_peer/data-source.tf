# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_wireguard_peer" "example" {
  filter = {
    name = "example"
  }
}

output "wireguard_peer" {
  value = data.routeros_wireguard_peer.example.entries[*].name
}
