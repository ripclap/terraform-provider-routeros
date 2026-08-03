# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_ip_ipsec_policy_group" "example" {
  filter = {
    name = "example"
  }
}

output "ip_ipsec_policy_group" {
  value = data.routeros_ip_ipsec_policy_group.example.entries[*].name
}
