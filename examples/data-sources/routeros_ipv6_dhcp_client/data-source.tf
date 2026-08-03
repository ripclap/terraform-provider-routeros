# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_ipv6_dhcp_client" "example" {
  filter = {
    interface = "ether1"
  }
}

output "ipv6_dhcp_client" {
  value = data.routeros_ipv6_dhcp_client.example.entries[*].id
}
