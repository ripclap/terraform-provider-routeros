# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_wifi_aaa" "example" {
  filter = {
    name = "example"
  }
}

output "wifi_aaa" {
  value = data.routeros_wifi_aaa.example.entries[*].name
}
