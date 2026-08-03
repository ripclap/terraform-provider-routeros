# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_port_remote_access" "example" {
  filter = {
    comment = "managed by terraform"
  }
}

output "port_remote_access" {
  value = data.routeros_port_remote_access.example.entries[*].id
}
