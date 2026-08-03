# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_interface_ethernet_switch_port" "example" {
  filter = {
    name = "example"
  }
}

output "interface_ethernet_switch_port" {
  value = data.routeros_interface_ethernet_switch_port.example.entries[*].name
}
