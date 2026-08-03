resource "routeros_ipv6_nd_proxy" "proxy" {
  address   = "192.0.2.1"
  interface = "ether1"
}
