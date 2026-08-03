resource "routeros_interface_ppp_client" "client" {
  name              = "example"
  add_default_route = true
  apn               = "example"
}
