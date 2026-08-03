resource "routeros_interface_pptp_client" "client" {
  name              = "example"
  add_default_route = true
  comment           = "Managed by OpenTofu"
}
