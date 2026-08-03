resource "routeros_interface_macsec" "macsec" {
  interface = "ether1"
  name      = "example"
}
