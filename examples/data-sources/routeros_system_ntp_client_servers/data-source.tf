# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_system_ntp_client_servers" "example" {
  filter = {
    comment = "managed by terraform"
  }
}

output "system_ntp_client_servers" {
  value = data.routeros_system_ntp_client_servers.example.entries[*].id
}
