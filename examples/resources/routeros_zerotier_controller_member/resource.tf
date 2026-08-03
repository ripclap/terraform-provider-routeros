resource "routeros_zerotier_controller_member" "member" {
  network    = "192.0.2.0/24"
  zt_address = "192.0.2.1"
}
