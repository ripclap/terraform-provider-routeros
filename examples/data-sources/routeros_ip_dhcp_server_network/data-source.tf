# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_ip_dhcp_server_network" "example" {
  filter = {
    comment = "managed by terraform"
  }
}

output "ip_dhcp_server_network" {
  value = data.routeros_ip_dhcp_server_network.example.entries[*].id
}
