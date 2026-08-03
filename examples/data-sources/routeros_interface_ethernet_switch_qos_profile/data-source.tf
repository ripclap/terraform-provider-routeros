# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_interface_ethernet_switch_qos_profile" "example" {
  filter = {
    name = "example"
  }
}

output "interface_ethernet_switch_qos_profile" {
  value = data.routeros_interface_ethernet_switch_qos_profile.example.entries[*].name
}
