# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_interface_bridge_port_mst_override" "example" {
  filter = {
    interface = "ether1"
  }
}

output "interface_bridge_port_mst_override" {
  value = data.routeros_interface_bridge_port_mst_override.example.entries[*].id
}
