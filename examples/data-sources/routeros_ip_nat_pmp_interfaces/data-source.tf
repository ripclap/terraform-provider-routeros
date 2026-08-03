# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_ip_nat_pmp_interfaces" "example" {
  filter = {
    interface = "ether1"
  }
}

output "ip_nat_pmp_interfaces" {
  value = data.routeros_ip_nat_pmp_interfaces.example.entries[*].id
}
