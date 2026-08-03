resource "routeros_routing_ospf_static_neighbor" "neighbor" {
  address = "192.0.2.1"
  area    = "example"
}
