# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_wifi_provisioning" "example" {
  filter = {
    comment = "managed by terraform"
  }
}

output "wifi_provisioning" {
  value = data.routeros_wifi_provisioning.example.entries[*].id
}
