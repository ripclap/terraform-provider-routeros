# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_wifi_security_multi_passphrase" "example" {
  filter = {
    comment = "managed by terraform"
  }
}

output "wifi_security_multi_passphrase" {
  value = data.routeros_wifi_security_multi_passphrase.example.entries[*].id
}
