# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_ip_dns_forwarders" "example" {
  filter = {
    name = "example"
  }
}

output "ip_dns_forwarders" {
  value = data.routeros_ip_dns_forwarders.example.entries[*].name
}
