resource "routeros_ip_socksify" "socksify" {
  name               = "example"
  comment            = "Managed by OpenTofu"
  connection_timeout = 1
}
