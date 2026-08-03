# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_ip_socks_users" "example" {
  filter = {
    name = "example"
  }
}

output "ip_socks_users" {
  value = data.routeros_ip_socks_users.example.entries[*].name
}
