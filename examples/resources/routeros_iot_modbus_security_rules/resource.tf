resource "routeros_iot_modbus_security_rules" "rules" {
  ip_range = "example"
  comment  = "Managed by OpenTofu"
  disabled = true
}
