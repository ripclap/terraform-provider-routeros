# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_ip_ipsec_identity" "example" {
  filter = {
    comment = "managed by terraform"
  }
}

output "ip_ipsec_identity" {
  value = data.routeros_ip_ipsec_identity.example.entries[*].id
}
