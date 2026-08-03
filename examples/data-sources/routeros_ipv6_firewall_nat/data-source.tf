# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_ipv6_firewall_nat" "example" {
  filter = {
    chain = "forward"
  }
}

output "ipv6_firewall_nat" {
  value = data.routeros_ipv6_firewall_nat.example.entries[*].id
}
