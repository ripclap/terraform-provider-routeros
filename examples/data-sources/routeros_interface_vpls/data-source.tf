# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_interface_vpls" "example" {
  filter = {
    name = "example"
  }
}

output "interface_vpls" {
  value = data.routeros_interface_vpls.example.entries[*].name
}
