# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_routing_ospf_area" "example" {
  filter = {
    name = "example"
  }
}

output "routing_ospf_area" {
  value = data.routeros_routing_ospf_area.example.entries[*].name
}
