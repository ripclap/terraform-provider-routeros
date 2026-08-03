# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_ip_ipsec_policy" "example" {
  filter = {
    comment = "managed by terraform"
  }
}

output "ip_ipsec_policy" {
  value = data.routeros_ip_ipsec_policy.example.entries[*].id
}
