# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_zerotier_controller" "example" {
  filter = {
    name = "example"
  }
}

output "zerotier_controller" {
  value = data.routeros_zerotier_controller.example.entries[*].name
}
