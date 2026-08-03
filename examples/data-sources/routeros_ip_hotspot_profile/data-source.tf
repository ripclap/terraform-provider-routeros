# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_ip_hotspot_profile" "example" {
  filter = {
    name = "example"
  }
}

output "ip_hotspot_profile" {
  value = data.routeros_ip_hotspot_profile.example.entries[*].name
}
