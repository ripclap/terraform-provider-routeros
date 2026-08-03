# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_routing_bfd_configuration" "example" {
  filter = {
    disabled = "false"
  }
}

output "routing_bfd_configuration" {
  value = data.routeros_routing_bfd_configuration.example.entries[*].id
}
