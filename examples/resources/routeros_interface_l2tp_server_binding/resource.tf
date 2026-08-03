resource "routeros_interface_l2tp_server_binding" "binding" {
  name     = "example"
  comment  = "Managed by OpenTofu"
  disabled = true
}
