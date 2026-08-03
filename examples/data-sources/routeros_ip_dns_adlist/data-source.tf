# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_ip_dns_adlist" "example" {
  filter = {
    comment = "managed by terraform"
  }
}

output "ip_dns_adlist" {
  value = data.routeros_ip_dns_adlist.example.entries[*].id
}
