# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_ip_socksify" "example" {
  filter = {
    name = "example"
  }
}

output "ip_socksify" {
  value = data.routeros_ip_socksify.example.entries[*].name
}
