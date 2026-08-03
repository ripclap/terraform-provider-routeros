# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_mpls_traffic_eng_path" "example" {
  filter = {
    name = "example"
  }
}

output "mpls_traffic_eng_path" {
  value = data.routeros_mpls_traffic_eng_path.example.entries[*].name
}
