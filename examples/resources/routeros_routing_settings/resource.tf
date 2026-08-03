resource "routeros_routing_settings" "settings" {
  check_gateway_ping_count    = 1
  check_gateway_ping_interval = "10s"
  check_gateway_ping_timeout  = "10s"
}
