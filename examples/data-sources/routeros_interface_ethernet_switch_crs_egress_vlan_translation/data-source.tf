# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_interface_ethernet_switch_crs_egress_vlan_translation" "example" {
  filter = {
    comment = "managed by terraform"
  }
}

output "interface_ethernet_switch_crs_egress_vlan_translation" {
  value = data.routeros_interface_ethernet_switch_crs_egress_vlan_translation.example.entries[*].id
}
