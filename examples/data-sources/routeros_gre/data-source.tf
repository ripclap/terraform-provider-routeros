# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_gre" "example" {
  filter = {
    name = "example"
  }
}

output "gre" {
  value = data.routeros_gre.example.entries[*].name
}
