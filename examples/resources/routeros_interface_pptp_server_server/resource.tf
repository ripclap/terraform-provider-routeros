resource "routeros_interface_pptp_server_server" "server" {
  default_profile   = "example"
  enabled           = true
  keepalive_timeout = "10s"
}
