resource "routeros_routing_rpki" "rpki" {
  address = "192.0.2.1"
  group   = "IPv4"
}
