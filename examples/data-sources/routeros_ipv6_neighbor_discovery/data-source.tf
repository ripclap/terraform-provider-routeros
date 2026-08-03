# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_ipv6_neighbor_discovery" "example" {
  filter = {
    interface = "ether1"
  }
}

output "ipv6_neighbor_discovery" {
  value = data.routeros_ipv6_neighbor_discovery.example.entries[*].id
}
