# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_dns_record" "example" {
  filter = {
    name = "example"
  }
}

output "dns_record" {
  value = data.routeros_dns_record.example.entries[*].name
}
