# Entries in the menu, narrowed by any attribute it returns.
# Omit `filter` to read all of them.
data "routeros_iot_modbus_security_rules" "example" {
  filter = {
    comment = "managed by terraform"
  }
}

output "iot_modbus_security_rules" {
  value = data.routeros_iot_modbus_security_rules.example.entries[*].id
}
