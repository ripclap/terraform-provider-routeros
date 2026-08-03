# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_system_user_group" "example" {
  filter = {
    name = "example"
  }
}

output "system_user_group" {
  value = data.routeros_system_user_group.example.entries[*].name
}
