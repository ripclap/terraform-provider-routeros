# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_firewall_nat" "example" {
  filter = {
    chain = "forward"
  }
}

output "firewall_nat" {
  value = data.routeros_firewall_nat.example.entries[*].id
}
