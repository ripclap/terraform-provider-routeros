# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_ip_traffic_flow_target" "example" {
  filter = {
    disabled = "false"
  }
}

output "ip_traffic_flow_target" {
  value = data.routeros_ip_traffic_flow_target.example.entries[*].id
}
