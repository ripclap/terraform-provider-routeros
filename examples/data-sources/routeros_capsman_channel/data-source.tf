# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_capsman_channel" "example" {
  filter = {
    name = "example"
  }
}

output "capsman_channel" {
  value = data.routeros_capsman_channel.example.entries[*].name
}
