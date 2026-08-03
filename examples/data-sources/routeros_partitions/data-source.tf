# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_partitions" "example" {
  filter = {
    name = "example"
  }
}

output "partitions" {
  value = data.routeros_partitions.example.entries[*].name
}
