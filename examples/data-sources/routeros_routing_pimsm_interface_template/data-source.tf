# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_routing_pimsm_interface_template" "example" {
  filter = {
    disabled = "false"
  }
}

output "routing_pimsm_interface_template" {
  value = data.routeros_routing_pimsm_interface_template.example.entries[*].id
}
