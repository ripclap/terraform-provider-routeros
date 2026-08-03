# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_capsman_rates" "example" {
  filter = {
    name = "example"
  }
}

output "capsman_rates" {
  value = data.routeros_capsman_rates.example.entries[*].name
}
