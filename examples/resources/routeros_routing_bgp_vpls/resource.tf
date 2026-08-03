resource "routeros_routing_bgp_vpls" "vpls" {
  name        = "example"
  bridge      = "example"
  bridge_cost = 1
}
