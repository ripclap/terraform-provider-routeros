# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_routing_filter_community_large_list" "example" {
  filter = {
    list = "example"
  }
}

output "routing_filter_community_large_list" {
  value = data.routeros_routing_filter_community_large_list.example.entries[*].id
}
