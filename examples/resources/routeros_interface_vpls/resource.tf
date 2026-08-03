resource "routeros_interface_vpls" "vpls" {
  name        = "example"
  arp         = "disabled"
  arp_timeout = "10s"
}
