# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_ppp_profile" "example" {
  filter = {
    name = "example"
  }
}

output "ppp_profile" {
  value = data.routeros_ppp_profile.example.entries[*].name
}
