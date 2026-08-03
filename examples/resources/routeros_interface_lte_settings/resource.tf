resource "routeros_interface_lte_settings" "settings" {
  esim_channel        = "at"
  firmware_path       = "example"
  link_recovery_timer = "10s"
}
