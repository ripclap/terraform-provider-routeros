# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_interface_ethernet_switch_qos_tx_manager" "example" {
  filter = {
    name = "example"
  }
}

output "interface_ethernet_switch_qos_tx_manager" {
  value = data.routeros_interface_ethernet_switch_qos_tx_manager.example.entries[*].name
}
