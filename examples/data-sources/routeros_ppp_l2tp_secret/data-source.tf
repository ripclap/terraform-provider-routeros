# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_ppp_l2tp_secret" "example" {
  filter = {
    comment = "managed by terraform"
  }
}

output "ppp_l2tp_secret" {
  value = data.routeros_ppp_l2tp_secret.example.entries[*].id
}
