# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_system_ntp_key" "example" {
  filter = {
    comment = "managed by terraform"
  }
}

output "system_ntp_key" {
  value = data.routeros_system_ntp_key.example.entries[*].id
}
