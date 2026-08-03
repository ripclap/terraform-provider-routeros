# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_ip_media" "example" {
  filter = {
    interface = "ether1"
  }
}

output "ip_media" {
  value = data.routeros_ip_media.example.entries[*].id
}
