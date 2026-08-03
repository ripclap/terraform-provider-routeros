# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_interface_eoip" "example" {
  filter = {
    name = "example"
  }
}

output "interface_eoip" {
  value = data.routeros_interface_eoip.example.entries[*].name
}
