# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_ipv6_dhcp_server_option_sets" "example" {
  filter = {
    name = "example"
  }
}

output "ipv6_dhcp_server_option_sets" {
  value = data.routeros_ipv6_dhcp_server_option_sets.example.entries[*].name
}
