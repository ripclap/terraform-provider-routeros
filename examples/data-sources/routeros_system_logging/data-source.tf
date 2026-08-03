# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_system_logging" "example" {
  filter = {
    disabled = "false"
  }
}

output "system_logging" {
  value = data.routeros_system_logging.example.entries[*].id
}
