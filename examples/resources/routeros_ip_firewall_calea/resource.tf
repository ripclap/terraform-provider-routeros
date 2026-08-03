resource "routeros_ip_firewall_calea" "calea" {
  action = "sniff"
  chain  = "forward"
}
