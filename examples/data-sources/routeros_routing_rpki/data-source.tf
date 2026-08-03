# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_routing_rpki" "example" {
  filter = {
    comment = "managed by terraform"
  }
}

output "routing_rpki" {
  value = data.routeros_routing_rpki.example.entries[*].id
}
