resource "routeros_interface_ipipv6" "ipipv6" {
  name          = "example"
  clamp_tcp_mss = true
  comment       = "Managed by OpenTofu"
}
