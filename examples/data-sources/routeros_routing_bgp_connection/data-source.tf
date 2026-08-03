# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_routing_bgp_connection" "example" {
  filter = {
    name = "example"
  }
}

output "routing_bgp_connection" {
  value = data.routeros_routing_bgp_connection.example.entries[*].name
}
