# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_ipv6_route" "example" {
  filter = {
    comment = "managed by terraform"
  }
}

output "ipv6_route" {
  value = data.routeros_ipv6_route.example.entries[*].id
}
