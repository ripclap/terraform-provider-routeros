# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_ipv6_pool" "example" {
  filter = {
    name = "example"
  }
}

output "ipv6_pool" {
  value = data.routeros_ipv6_pool.example.entries[*].name
}
