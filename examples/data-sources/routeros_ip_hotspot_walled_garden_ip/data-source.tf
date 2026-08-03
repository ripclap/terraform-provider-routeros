# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_ip_hotspot_walled_garden_ip" "example" {
  filter = {
    comment = "managed by terraform"
  }
}

output "ip_hotspot_walled_garden_ip" {
  value = data.routeros_ip_hotspot_walled_garden_ip.example.entries[*].id
}
