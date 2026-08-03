# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_certificate_scep_server_ra" "example" {
  filter = {
    name = "example"
  }
}

output "certificate_scep_server_ra" {
  value = data.routeros_certificate_scep_server_ra.example.entries[*].name
}
