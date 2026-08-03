# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_dhcp_client" "example" {
  filter = {
    interface = "ether1"
  }
}

output "dhcp_client" {
  value = data.routeros_dhcp_client.example.entries[*].id
}
