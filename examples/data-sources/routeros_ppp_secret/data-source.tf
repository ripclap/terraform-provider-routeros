# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_ppp_secret" "example" {
  filter = {
    name = "example"
  }
}

output "ppp_secret" {
  value = data.routeros_ppp_secret.example.entries[*].name
}
