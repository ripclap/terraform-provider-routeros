resource "routeros_routing_rip_static_neighbor" "neighbor" {
  address  = "192.0.2.1"
  instance = "example"
}
