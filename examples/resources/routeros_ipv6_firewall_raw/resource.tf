resource "routeros_ipv6_firewall_raw" "raw" {
  action = "accept"
  chain  = "forward"
}
