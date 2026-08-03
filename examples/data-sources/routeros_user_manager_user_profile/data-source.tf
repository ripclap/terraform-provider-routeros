# Every entry in the menu.
data "routeros_user_manager_user_profile" "example" {}

output "user_manager_user_profile" {
  value = data.routeros_user_manager_user_profile.example.entries[*].id
}
