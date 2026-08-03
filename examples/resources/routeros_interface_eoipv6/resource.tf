resource "routeros_interface_eoipv6" "eoipv6" {
  name        = "example"
  arp         = "disabled"
  arp_timeout = "10s"
}
