resource "routeros_system_ptp_port" "port" {
  interface = "ether1"
  ptp       = "example"
}
