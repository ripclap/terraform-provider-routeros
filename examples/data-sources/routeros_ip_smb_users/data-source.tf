# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_ip_smb_users" "example" {
  filter = {
    name = "example"
  }
}

output "ip_smb_users" {
  value = data.routeros_ip_smb_users.example.entries[*].name
}
