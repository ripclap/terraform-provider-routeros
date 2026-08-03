# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_container_mounts" "example" {
  filter = {
    name = "example"
  }
}

output "container_mounts" {
  value = data.routeros_container_mounts.example.entries[*].name
}
