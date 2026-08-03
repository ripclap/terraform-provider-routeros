# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_capsman_datapath" "example" {
  filter = {
    name = "example"
  }
}

output "capsman_datapath" {
  value = data.routeros_capsman_datapath.example.entries[*].name
}
