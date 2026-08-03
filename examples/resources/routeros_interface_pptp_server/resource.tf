resource "routeros_interface_pptp_server" "server" {
  name     = "example"
  comment  = "Managed by OpenTofu"
  disabled = true
}
