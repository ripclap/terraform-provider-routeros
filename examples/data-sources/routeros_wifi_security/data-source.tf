# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_wifi_security" "example" {
  filter = {
    name = "example"
  }
}

output "wifi_security" {
  value = data.routeros_wifi_security.example.entries[*].name
}
