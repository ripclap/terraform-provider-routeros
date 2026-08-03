# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_routing_bgp_vpn" "example" {
  filter = {
    name = "example"
  }
}

output "routing_bgp_vpn" {
  value = data.routeros_routing_bgp_vpn.example.entries[*].name
}
