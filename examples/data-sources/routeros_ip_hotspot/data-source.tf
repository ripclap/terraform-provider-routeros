# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_ip_hotspot" "example" {
  filter = {
    name = "example"
  }
}

output "ip_hotspot" {
  value = data.routeros_ip_hotspot.example.entries[*].name
}
