# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_routing_ospf_instance" "example" {
  filter = {
    name = "example"
  }
}

output "routing_ospf_instance" {
  value = data.routeros_routing_ospf_instance.example.entries[*].name
}
