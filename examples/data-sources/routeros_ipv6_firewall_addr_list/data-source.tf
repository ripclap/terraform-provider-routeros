# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_ipv6_firewall_addr_list" "example" {
  filter = {
    list = "example"
  }
}

output "ipv6_firewall_addr_list" {
  value = data.routeros_ipv6_firewall_addr_list.example.entries[*].id
}
