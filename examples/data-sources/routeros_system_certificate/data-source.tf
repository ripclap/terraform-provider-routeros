# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_system_certificate" "example" {
  filter = {
    name = "example"
  }
}

output "system_certificate" {
  value = data.routeros_system_certificate.example.entries[*].name
}
