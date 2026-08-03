# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_wireguard" "example" {
  filter = {
    name = "example"
  }
}

output "wireguard" {
  value = data.routeros_wireguard.example.entries[*].name
}
