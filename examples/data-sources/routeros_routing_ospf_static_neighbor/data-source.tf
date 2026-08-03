# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_routing_ospf_static_neighbor" "example" {
  filter = {
    comment = "managed by terraform"
  }
}

output "routing_ospf_static_neighbor" {
  value = data.routeros_routing_ospf_static_neighbor.example.entries[*].id
}
