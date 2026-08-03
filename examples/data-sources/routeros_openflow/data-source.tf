# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_openflow" "example" {
  filter = {
    name = "example"
  }
}

output "openflow" {
  value = data.routeros_openflow.example.entries[*].name
}
