# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_interface_sstp_server_interface" "example" {
  filter = {
    name = "example"
  }
}

output "interface_sstp_server_interface" {
  value = data.routeros_interface_sstp_server_interface.example.entries[*].name
}
