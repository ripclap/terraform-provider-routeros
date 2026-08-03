# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_app" "example" {
  filter = {
    name = "example"
  }
}

output "app" {
  value = data.routeros_app.example.entries[*].name
}
