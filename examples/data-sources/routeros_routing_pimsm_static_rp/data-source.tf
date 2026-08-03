# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_routing_pimsm_static_rp" "example" {
  filter = {
    comment = "managed by terraform"
  }
}

output "routing_pimsm_static_rp" {
  value = data.routeros_routing_pimsm_static_rp.example.entries[*].id
}
