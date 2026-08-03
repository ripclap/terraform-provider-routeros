# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_routing_filter_community_ext_list" "example" {
  filter = {
    list = "example"
  }
}

output "routing_filter_community_ext_list" {
  value = data.routeros_routing_filter_community_ext_list.example.entries[*].id
}
