resource "routeros_interface_bridge_nat" "nat" {
  action = "accept"
  chain  = "forward"
}
