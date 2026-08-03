# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_system_ups" "example" {
  filter = {
    name = "example"
  }
}

output "system_ups" {
  value = data.routeros_system_ups.example.entries[*].name
}
