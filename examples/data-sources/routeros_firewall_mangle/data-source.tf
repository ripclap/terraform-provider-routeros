# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_firewall_mangle" "example" {
  filter = {
    chain = "forward"
  }
}

output "firewall_mangle" {
  value = data.routeros_firewall_mangle.example.entries[*].id
}
