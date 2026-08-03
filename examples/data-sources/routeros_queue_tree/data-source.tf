# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_queue_tree" "example" {
  filter = {
    name = "example"
  }
}

output "queue_tree" {
  value = data.routeros_queue_tree.example.entries[*].name
}
