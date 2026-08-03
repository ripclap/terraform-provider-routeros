# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_interface_ethernet_switch_qos_map_ip" "example" {
  filter = {
    comment = "managed by terraform"
  }
}

output "interface_ethernet_switch_qos_map_ip" {
  value = data.routeros_interface_ethernet_switch_qos_map_ip.example.entries[*].id
}
