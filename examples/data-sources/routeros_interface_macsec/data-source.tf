# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_interface_macsec" "example" {
  filter = {
    name = "example"
  }
}

output "interface_macsec" {
  value = data.routeros_interface_macsec.example.entries[*].name
}
