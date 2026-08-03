# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_routing_rule" "example" {
  filter = {
    interface = "ether1"
  }
}

output "routing_rule" {
  value = data.routeros_routing_rule.example.entries[*].id
}
