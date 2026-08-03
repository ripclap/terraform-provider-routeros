resource "routeros_ip_reverse_proxy" "proxy" {
  certificate = "example"
  comment     = "Managed by OpenTofu"
  disabled    = true
}
