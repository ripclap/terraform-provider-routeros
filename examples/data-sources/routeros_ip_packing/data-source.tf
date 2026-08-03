# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_ip_packing" "example" {
  filter = {
    interface = "ether1"
  }
}

output "ip_packing" {
  value = data.routeros_ip_packing.example.entries[*].id
}
