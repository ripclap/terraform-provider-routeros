# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_mpls_ldp_neighbor" "example" {
  filter = {
    comment = "managed by terraform"
  }
}

output "mpls_ldp_neighbor" {
  value = data.routeros_mpls_ldp_neighbor.example.entries[*].id
}
