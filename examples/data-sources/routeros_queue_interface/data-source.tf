# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_queue_interface" "example" {
  filter = {
    interface = "ether1"
  }
}

output "queue_interface" {
  value = data.routeros_queue_interface.example.entries[*].id
}
