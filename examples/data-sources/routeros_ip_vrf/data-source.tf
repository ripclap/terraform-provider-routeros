# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_ip_vrf" "example" {
  filter = {
    name = "example"
  }
}

output "ip_vrf" {
  value = data.routeros_ip_vrf.example.entries[*].name
}
