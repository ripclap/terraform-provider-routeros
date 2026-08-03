# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_wifi_steering" "example" {
  filter = {
    name = "example"
  }
}

output "wifi_steering" {
  value = data.routeros_wifi_steering.example.entries[*].name
}
