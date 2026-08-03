resource "routeros_ip_firewall_raw" "raw" {
  action = "accept"
  chain  = "forward"
}
