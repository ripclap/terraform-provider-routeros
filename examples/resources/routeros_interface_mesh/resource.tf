resource "routeros_interface_mesh" "mesh" {
  name      = "example"
  admin_mac = "00:00:5E:00:53:01"
  arp       = "disabled"
}
