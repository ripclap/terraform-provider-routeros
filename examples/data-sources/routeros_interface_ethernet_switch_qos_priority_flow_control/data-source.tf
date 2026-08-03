# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_interface_ethernet_switch_qos_priority_flow_control" "example" {
  filter = {
    name = "example"
  }
}

output "interface_ethernet_switch_qos_priority_flow_control" {
  value = data.routeros_interface_ethernet_switch_qos_priority_flow_control.example.entries[*].name
}
