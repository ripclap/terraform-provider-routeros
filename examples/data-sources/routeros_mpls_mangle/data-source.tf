# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_mpls_mangle" "example" {
  filter = {
    chain = "forward"
  }
}

output "mpls_mangle" {
  value = data.routeros_mpls_mangle.example.entries[*].id
}
