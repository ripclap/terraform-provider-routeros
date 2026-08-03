# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_ip_smb_shares" "example" {
  filter = {
    name = "example"
  }
}

output "ip_smb_shares" {
  value = data.routeros_ip_smb_shares.example.entries[*].name
}
