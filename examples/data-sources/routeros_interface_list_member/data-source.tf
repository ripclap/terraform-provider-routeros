# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_interface_list_member" "example" {
  filter = {
    interface = "ether1"
  }
}

output "interface_list_member" {
  value = data.routeros_interface_list_member.example.entries[*].id
}
