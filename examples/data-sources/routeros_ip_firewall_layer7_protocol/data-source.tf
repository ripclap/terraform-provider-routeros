# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_ip_firewall_layer7_protocol" "example" {
  filter = {
    name = "example"
  }
}

output "ip_firewall_layer7_protocol" {
  value = data.routeros_ip_firewall_layer7_protocol.example.entries[*].name
}
