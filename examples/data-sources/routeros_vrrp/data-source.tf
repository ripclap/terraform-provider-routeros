# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_vrrp" "example" {
  filter = {
    name = "example"
  }
}

output "vrrp" {
  value = data.routeros_vrrp.example.entries[*].name
}
