resource "routeros_system_console" "console" {
  port     = "8728"
  channel  = 1
  disabled = true
}
