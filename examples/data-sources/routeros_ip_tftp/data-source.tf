# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_ip_tftp" "example" {
  filter = {
    disabled = "false"
  }
}

output "ip_tftp" {
  value = data.routeros_ip_tftp.example.entries[*].id
}
