# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_ip_pool" "example" {
  filter = {
    name = "example"
  }
}

output "ip_pool" {
  value = data.routeros_ip_pool.example.entries[*].name
}
