# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_disk" "example" {
  filter = {
    interface = "ether1"
  }
}

output "disk" {
  value = data.routeros_disk.example.entries[*].id
}
