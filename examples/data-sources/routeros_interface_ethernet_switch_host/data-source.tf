# Every entry in the menu.
data "routeros_interface_ethernet_switch_host" "example" {}

output "interface_ethernet_switch_host" {
  value = data.routeros_interface_ethernet_switch_host.example.entries[*].id
}
