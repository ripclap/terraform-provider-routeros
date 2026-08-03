# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_system_scheduler" "example" {
  filter = {
    name = "example"
  }
}

output "system_scheduler" {
  value = data.routeros_system_scheduler.example.entries[*].name
}
