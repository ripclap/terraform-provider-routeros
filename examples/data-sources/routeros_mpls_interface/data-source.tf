# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_mpls_interface" "example" {
  filter = {
    interface = "ether1"
  }
}

output "mpls_interface" {
  value = data.routeros_mpls_interface.example.entries[*].id
}
