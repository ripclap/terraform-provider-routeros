# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_routing_filter_rule" "example" {
  filter = {
    chain = "forward"
  }
}

output "routing_filter_rule" {
  value = data.routeros_routing_filter_rule.example.entries[*].id
}
