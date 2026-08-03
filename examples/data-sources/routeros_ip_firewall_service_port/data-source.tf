# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_ip_firewall_service_port" "example" {
  filter = {
    name = "example"
  }
}

output "ip_firewall_service_port" {
  value = data.routeros_ip_firewall_service_port.example.entries[*].name
}
