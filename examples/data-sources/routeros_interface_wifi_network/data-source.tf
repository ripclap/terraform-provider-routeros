# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_interface_wifi_network" "example" {
  filter = {
    comment = "managed by terraform"
  }
}

output "interface_wifi_network" {
  value = data.routeros_interface_wifi_network.example.entries[*].id
}
