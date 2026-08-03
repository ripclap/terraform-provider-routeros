# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_ip_reverse_proxy" "example" {
  filter = {
    comment = "managed by terraform"
  }
}

output "ip_reverse_proxy" {
  value = data.routeros_ip_reverse_proxy.example.entries[*].id
}
