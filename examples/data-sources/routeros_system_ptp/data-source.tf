# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_system_ptp" "example" {
  filter = {
    name = "example"
  }
}

output "system_ptp" {
  value = data.routeros_system_ptp.example.entries[*].name
}
