resource "routeros_ipv6_neighbor" "neighbor" {
  address   = "192.0.2.1"
  interface = "ether1"
}
