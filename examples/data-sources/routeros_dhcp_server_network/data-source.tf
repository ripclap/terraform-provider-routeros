# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_dhcp_server_network" "example" {
  filter = {
    comment = "managed by terraform"
  }
}

output "dhcp_server_network" {
  value = data.routeros_dhcp_server_network.example.entries[*].id
}
