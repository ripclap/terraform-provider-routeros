resource "routeros_ip_socks" "socks" {
  auth_method             = "none"
  connection_idle_timeout = "10s"
  enabled                 = true
}
