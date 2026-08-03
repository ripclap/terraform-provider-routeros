resource "routeros_system_ntp_client_servers" "servers" {
  address  = "192.0.2.1"
  auth_key = "changeme"
  comment  = "Managed by OpenTofu"
}
