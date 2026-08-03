# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_file_sync" "example" {
  filter = {
    comment = "managed by terraform"
  }
}

output "file_sync" {
  value = data.routeros_file_sync.example.entries[*].id
}
