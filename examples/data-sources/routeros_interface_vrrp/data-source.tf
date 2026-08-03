# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_interface_vrrp" "example" {
  filter = {
    name = "example"
  }
}

output "interface_vrrp" {
  value = data.routeros_interface_vrrp.example.entries[*].name
}
