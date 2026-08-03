# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_ip_dhcp_server_alert" "example" {
  filter = {
    interface = "ether1"
  }
}

output "ip_dhcp_server_alert" {
  value = data.routeros_ip_dhcp_server_alert.example.entries[*].id
}
