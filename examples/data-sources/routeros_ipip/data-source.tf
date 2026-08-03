# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_ipip" "example" {
  filter = {
    name = "example"
  }
}

output "ipip" {
  value = data.routeros_ipip.example.entries[*].name
}
