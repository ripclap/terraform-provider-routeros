# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_interface_wireless_security_profiles" "example" {
  filter = {
    name = "example"
  }
}

output "interface_wireless_security_profiles" {
  value = data.routeros_interface_wireless_security_profiles.example.entries[*].name
}
