# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_user_manager_router" "example" {
  filter = {
    name = "example"
  }
}

output "user_manager_router" {
  value = data.routeros_user_manager_router.example.entries[*].name
}
