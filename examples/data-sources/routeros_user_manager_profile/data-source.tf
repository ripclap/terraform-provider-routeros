# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_user_manager_profile" "example" {
  filter = {
    name = "example"
  }
}

output "user_manager_profile" {
  value = data.routeros_user_manager_profile.example.entries[*].name
}
