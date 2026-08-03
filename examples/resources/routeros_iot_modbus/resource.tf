resource "routeros_iot_modbus" "modbus" {
  disable_security_rules = true
  disabled               = true
  hardware_port          = "8728"
}
