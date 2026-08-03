# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_port" "example" {
  filter = {
    name = "example"
  }
}

output "port" {
  value = data.routeros_port.example.entries[*].name
}
