# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_snmp_community" "example" {
  filter = {
    name = "example"
  }
}

output "snmp_community" {
  value = data.routeros_snmp_community.example.entries[*].name
}
