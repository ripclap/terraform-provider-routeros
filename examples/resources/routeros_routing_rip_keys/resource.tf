resource "routeros_routing_rip_keys" "keys" {
  chain    = "forward"
  disabled = true
  key      = "changeme"
}
