resource "routeros_ip_firewall_service_port" "port" {
  name     = "example"
  disabled = true
  ports    = "example"
}
