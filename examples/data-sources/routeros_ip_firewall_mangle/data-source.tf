# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_ip_firewall_mangle" "example" {
  filter = {
    chain = "forward"
  }
}

output "ip_firewall_mangle" {
  value = data.routeros_ip_firewall_mangle.example.entries[*].id
}
