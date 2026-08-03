# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_interface_wireguard" "example" {
  filter = {
    name = "example"
  }
}

output "interface_wireguard" {
  value = data.routeros_interface_wireguard.example.entries[*].name
}
