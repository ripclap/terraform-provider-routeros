# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_interface_lte_apn" "example" {
  filter = {
    name = "example"
  }
}

output "interface_lte_apn" {
  value = data.routeros_interface_lte_apn.example.entries[*].name
}
