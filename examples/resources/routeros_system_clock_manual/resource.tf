resource "routeros_system_clock_manual" "manual" {
  dst_delta = "192.0.2.1"
  dst_end   = "192.0.2.1"
  dst_start = "192.0.2.1"
}
