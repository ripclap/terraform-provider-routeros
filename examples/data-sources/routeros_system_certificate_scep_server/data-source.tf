# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_system_certificate_scep_server" "example" {
  filter = {
    disabled = "false"
  }
}

output "system_certificate_scep_server" {
  value = data.routeros_system_certificate_scep_server.example.entries[*].id
}
