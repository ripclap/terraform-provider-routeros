# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_wifi_channel" "example" {
  filter = {
    name = "example"
  }
}

output "wifi_channel" {
  value = data.routeros_wifi_channel.example.entries[*].name
}
