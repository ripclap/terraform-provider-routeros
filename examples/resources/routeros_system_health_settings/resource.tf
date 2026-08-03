resource "routeros_system_health_settings" "settings" {
  cpu_overtemp_check         = true
  cpu_overtemp_startup_delay = "10s"
  cpu_overtemp_threshold     = 1
}
