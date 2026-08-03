# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_ip_dns_record" "example" {
  filter = {
    name = "example"
  }
}

output "ip_dns_record" {
  value = data.routeros_ip_dns_record.example.entries[*].name
}
