# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_system_user" "example" {
  filter = {
    name = "example"
  }
}

output "system_user" {
  value = data.routeros_system_user.example.entries[*].name
}
