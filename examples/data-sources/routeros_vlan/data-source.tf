# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_vlan" "example" {
  filter = {
    name = "example"
  }
}

output "vlan" {
  value = data.routeros_vlan.example.entries[*].name
}
