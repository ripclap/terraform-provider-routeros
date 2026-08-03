# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_system_console" "example" {
  filter = {
    disabled = "false"
  }
}

output "system_console" {
  value = data.routeros_system_console.example.entries[*].id
}
