# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_routing_rip_keys" "example" {
  filter = {
    chain = "forward"
  }
}

output "routing_rip_keys" {
  value = data.routeros_routing_rip_keys.example.entries[*].id
}
