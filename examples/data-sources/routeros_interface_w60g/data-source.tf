# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_interface_w60g" "example" {
  filter = {
    name = "example"
  }
}

output "interface_w60g" {
  value = data.routeros_interface_w60g.example.entries[*].name
}
