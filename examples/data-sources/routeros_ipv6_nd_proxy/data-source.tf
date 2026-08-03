# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_ipv6_nd_proxy" "example" {
  filter = {
    interface = "ether1"
  }
}

output "ipv6_nd_proxy" {
  value = data.routeros_ipv6_nd_proxy.example.entries[*].id
}
