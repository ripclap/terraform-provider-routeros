# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_container_envs" "example" {
  filter = {
    name = "example"
  }
}

output "container_envs" {
  value = data.routeros_container_envs.example.entries[*].name
}
