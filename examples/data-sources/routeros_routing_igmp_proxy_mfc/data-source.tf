# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_routing_igmp_proxy_mfc" "example" {
  filter = {
    disabled = "false"
  }
}

output "routing_igmp_proxy_mfc" {
  value = data.routeros_routing_igmp_proxy_mfc.example.entries[*].id
}
