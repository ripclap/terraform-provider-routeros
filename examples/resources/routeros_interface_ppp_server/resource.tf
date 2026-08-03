resource "routeros_interface_ppp_server" "server" {
  name         = "example"
  comment      = "Managed by OpenTofu"
  data_channel = "example"
}
