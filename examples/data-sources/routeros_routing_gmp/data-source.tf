# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_routing_gmp" "example" {
  filter = {
    disabled = "false"
  }
}

output "routing_gmp" {
  value = data.routeros_routing_gmp.example.entries[*].id
}
