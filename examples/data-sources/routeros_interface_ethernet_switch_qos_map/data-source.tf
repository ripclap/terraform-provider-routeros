# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_interface_ethernet_switch_qos_map" "example" {
  filter = {
    name = "example"
  }
}

output "interface_ethernet_switch_qos_map" {
  value = data.routeros_interface_ethernet_switch_qos_map.example.entries[*].name
}
