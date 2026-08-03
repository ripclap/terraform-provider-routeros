# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_routing_pimsm_instance" "example" {
  filter = {
    name = "example"
  }
}

output "routing_pimsm_instance" {
  value = data.routeros_routing_pimsm_instance.example.entries[*].name
}
