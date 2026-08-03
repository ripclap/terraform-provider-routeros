# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_interface_ppp_client" "example" {
  filter = {
    name = "example"
  }
}

output "interface_ppp_client" {
  value = data.routeros_interface_ppp_client.example.entries[*].name
}
