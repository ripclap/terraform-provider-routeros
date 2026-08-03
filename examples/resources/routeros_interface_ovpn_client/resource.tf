resource "routeros_interface_ovpn_client" "client" {
  connect_to = "example"
  name       = "example"
  user       = "example"
}
