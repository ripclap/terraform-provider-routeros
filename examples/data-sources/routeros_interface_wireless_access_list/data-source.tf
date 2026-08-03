# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_interface_wireless_access_list" "example" {
  filter = {
    interface = "ether1"
  }
}

output "interface_wireless_access_list" {
  value = data.routeros_interface_wireless_access_list.example.entries[*].id
}
