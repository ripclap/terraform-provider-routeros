# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_interface_bridge_nat" "example" {
  filter = {
    chain = "forward"
  }
}

output "interface_bridge_nat" {
  value = data.routeros_interface_bridge_nat.example.entries[*].id
}
