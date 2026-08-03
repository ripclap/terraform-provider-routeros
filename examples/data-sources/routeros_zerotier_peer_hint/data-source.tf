# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_zerotier_peer_hint" "example" {
  filter = {
    comment = "managed by terraform"
  }
}

output "zerotier_peer_hint" {
  value = data.routeros_zerotier_peer_hint.example.entries[*].id
}
