# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_interface_w60g_station" "example" {
  filter = {
    name = "example"
  }
}

output "interface_w60g_station" {
  value = data.routeros_interface_w60g_station.example.entries[*].name
}
