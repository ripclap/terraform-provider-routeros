resource "routeros_ipv6_dhcp_server_binding" "binding" {
  duid                   = "example"
  address                = "192.0.2.1"
  allow_dual_stack_queue = true
}
