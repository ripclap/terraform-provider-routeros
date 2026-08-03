# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_routing_igmp_proxy_interface" "example" {
  filter = {
    interface = "ether1"
  }
}

output "routing_igmp_proxy_interface" {
  value = data.routeros_routing_igmp_proxy_interface.example.entries[*].id
}
