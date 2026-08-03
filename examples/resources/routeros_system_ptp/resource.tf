resource "routeros_system_ptp" "ptp" {
  name       = "example"
  comment    = "Managed by OpenTofu"
  delay_mode = "auto"
}
