# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_ip_hotspot_walled_garden" "example" {
  filter = {
    comment = "managed by terraform"
  }
}

output "ip_hotspot_walled_garden" {
  value = data.routeros_ip_hotspot_walled_garden.example.entries[*].id
}
