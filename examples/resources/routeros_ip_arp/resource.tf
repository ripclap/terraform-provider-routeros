resource "routeros_ip_arp" "arp" {
  interface = "ether1"
  address   = "192.0.2.1"
  comment   = "Managed by OpenTofu"
}
