# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_wifi_datapath" "example" {
  filter = {
    name = "example"
  }
}

output "wifi_datapath" {
  value = data.routeros_wifi_datapath.example.entries[*].name
}
