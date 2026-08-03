# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_queue_type" "example" {
  filter = {
    name = "example"
  }
}

output "queue_type" {
  value = data.routeros_queue_type.example.entries[*].name
}
