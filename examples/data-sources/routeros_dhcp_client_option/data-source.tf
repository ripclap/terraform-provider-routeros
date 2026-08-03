# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_dhcp_client_option" "example" {
  filter = {
    name = "example"
  }
}

output "dhcp_client_option" {
  value = data.routeros_dhcp_client_option.example.entries[*].name
}
