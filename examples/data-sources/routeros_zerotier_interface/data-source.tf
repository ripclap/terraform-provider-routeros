# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_zerotier_interface" "example" {
  filter = {
    name = "example"
  }
}

output "zerotier_interface" {
  value = data.routeros_zerotier_interface.example.entries[*].name
}
