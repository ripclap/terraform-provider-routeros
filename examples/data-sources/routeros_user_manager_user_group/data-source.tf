# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_user_manager_user_group" "example" {
  filter = {
    name = "example"
  }
}

output "user_manager_user_group" {
  value = data.routeros_user_manager_user_group.example.entries[*].name
}
