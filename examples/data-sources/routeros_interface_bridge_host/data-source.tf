# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_interface_bridge_host" "example" {
  filter = {
    interface = "ether1"
  }
}

output "interface_bridge_host" {
  value = data.routeros_interface_bridge_host.example.entries[*].id
}
