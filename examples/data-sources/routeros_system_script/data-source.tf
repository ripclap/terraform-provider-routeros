# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_system_script" "example" {
  filter = {
    name = "example"
  }
}

output "system_script" {
  value = data.routeros_system_script.example.entries[*].name
}
