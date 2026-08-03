# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_routing_bgp_instance" "example" {
  filter = {
    name = "example"
  }
}

output "routing_bgp_instance" {
  value = data.routeros_routing_bgp_instance.example.entries[*].name
}
