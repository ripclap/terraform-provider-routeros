# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_mpls_traffic_eng_tunnel" "example" {
  filter = {
    name = "example"
  }
}

output "mpls_traffic_eng_tunnel" {
  value = data.routeros_mpls_traffic_eng_tunnel.example.entries[*].name
}
