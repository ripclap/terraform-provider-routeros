# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_routing_pimsm_bsr_rp_candidate" "example" {
  filter = {
    disabled = "false"
  }
}

output "routing_pimsm_bsr_rp_candidate" {
  value = data.routeros_routing_pimsm_bsr_rp_candidate.example.entries[*].id
}
