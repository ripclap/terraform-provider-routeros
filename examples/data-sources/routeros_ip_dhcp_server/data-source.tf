# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_ip_dhcp_server" "example" {
  filter = {
    name = "example"
  }
}

output "ip_dhcp_server" {
  value = data.routeros_ip_dhcp_server.example.entries[*].name
}
