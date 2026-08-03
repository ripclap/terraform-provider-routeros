# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_routing_ospf_area_range" "example" {
  filter = {
    comment = "managed by terraform"
  }
}

output "routing_ospf_area_range" {
  value = data.routeros_routing_ospf_area_range.example.entries[*].id
}
