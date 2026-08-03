# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_ipv6_dhcp_server_binding" "example" {
  filter = {
    comment = "managed by terraform"
  }
}

output "ipv6_dhcp_server_binding" {
  value = data.routeros_ipv6_dhcp_server_binding.example.entries[*].id
}
