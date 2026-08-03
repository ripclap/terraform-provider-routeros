resource "routeros_interface_sstp_server_interface" "interface" {
  name     = "example"
  comment  = "Managed by OpenTofu"
  disabled = true
}
