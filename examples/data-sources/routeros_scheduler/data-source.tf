# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_scheduler" "example" {
  filter = {
    name = "example"
  }
}

output "scheduler" {
  value = data.routeros_scheduler.example.entries[*].name
}
