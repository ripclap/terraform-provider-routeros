# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_ip_dhcp_relay" "example" {
  filter = {
    name = "example"
  }
}

output "ip_dhcp_relay" {
  value = data.routeros_ip_dhcp_relay.example.entries[*].name
}
