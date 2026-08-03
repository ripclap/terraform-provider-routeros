# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_interface_ipipv6" "example" {
  filter = {
    name = "example"
  }
}

output "interface_ipipv6" {
  value = data.routeros_interface_ipipv6.example.entries[*].name
}
