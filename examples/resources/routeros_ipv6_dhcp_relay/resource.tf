resource "routeros_ipv6_dhcp_relay" "relay" {
  dhcp_server = ["192.0.2.1"]
  interface   = "ether1"
  name        = "example"
}
