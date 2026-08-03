# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_interface_macsec_profile" "example" {
  filter = {
    name = "example"
  }
}

output "interface_macsec_profile" {
  value = data.routeros_interface_macsec_profile.example.entries[*].name
}
