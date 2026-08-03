# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_ip_proxy_direct" "example" {
  filter = {
    comment = "managed by terraform"
  }
}

output "ip_proxy_direct" {
  value = data.routeros_ip_proxy_direct.example.entries[*].id
}
