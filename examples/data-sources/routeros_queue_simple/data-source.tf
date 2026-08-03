# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_queue_simple" "example" {
  filter = {
    name = "example"
  }
}

output "queue_simple" {
  value = data.routeros_queue_simple.example.entries[*].name
}
