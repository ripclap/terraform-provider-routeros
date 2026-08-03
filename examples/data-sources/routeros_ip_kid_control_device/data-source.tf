# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_ip_kid_control_device" "example" {
  filter = {
    name = "example"
  }
}

output "ip_kid_control_device" {
  value = data.routeros_ip_kid_control_device.example.entries[*].name
}
