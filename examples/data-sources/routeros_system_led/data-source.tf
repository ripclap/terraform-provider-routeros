# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_system_led" "example" {
  filter = {
    interface = "ether1"
  }
}

output "system_led" {
  value = data.routeros_system_led.example.entries[*].id
}
