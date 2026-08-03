# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_mpls_ldp_remote_mapping" "example" {
  filter = {
    comment = "managed by terraform"
  }
}

output "mpls_ldp_remote_mapping" {
  value = data.routeros_mpls_ldp_remote_mapping.example.entries[*].id
}
