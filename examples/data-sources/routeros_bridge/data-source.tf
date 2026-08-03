# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_bridge" "example" {
  filter = {
    name = "example"
  }
}

output "bridge" {
  value = data.routeros_bridge.example.entries[*].name
}
