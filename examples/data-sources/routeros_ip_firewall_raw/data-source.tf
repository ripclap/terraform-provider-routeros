# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_ip_firewall_raw" "example" {
  filter = {
    chain = "forward"
  }
}

output "ip_firewall_raw" {
  value = data.routeros_ip_firewall_raw.example.entries[*].id
}
