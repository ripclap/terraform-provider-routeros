# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_interface_ethernet_switch_port_isolation" "example" {
  filter = {
    name = "example"
  }
}

output "interface_ethernet_switch_port_isolation" {
  value = data.routeros_interface_ethernet_switch_port_isolation.example.entries[*].name
}
