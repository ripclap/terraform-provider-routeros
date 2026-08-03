resource "routeros_system_watchdog" "watchdog" {
  auto_send_supout      = true
  automatic_supout      = true
  ping_start_after_boot = "10s"
}
