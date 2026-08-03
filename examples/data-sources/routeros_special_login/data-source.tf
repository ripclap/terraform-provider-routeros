# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_special_login" "example" {
  filter = {
    disabled = "false"
  }
}

output "special_login" {
  value = data.routeros_special_login.example.entries[*].id
}
