# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_ip_dhcp_server_option_set" "example" {
  filter = {
    name = "example"
  }
}

output "ip_dhcp_server_option_set" {
  value = data.routeros_ip_dhcp_server_option_set.example.entries[*].name
}
