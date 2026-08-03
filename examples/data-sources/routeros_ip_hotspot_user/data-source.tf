# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_ip_hotspot_user" "example" {
  filter = {
    name = "example"
  }
}

output "ip_hotspot_user" {
  value = data.routeros_ip_hotspot_user.example.entries[*].name
}
