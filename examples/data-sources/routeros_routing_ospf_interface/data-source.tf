# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_routing_ospf_interface" "example" {
  filter = {
    interface = "ether1"
  }
}

output "routing_ospf_interface" {
  value = data.routeros_routing_ospf_interface.example.entries[*].id
}
