# Every entry in the menu.
data "routeros_user_manager_profile_limitation" "example" {}

output "user_manager_profile_limitation" {
  value = data.routeros_user_manager_profile_limitation.example.entries[*].id
}
