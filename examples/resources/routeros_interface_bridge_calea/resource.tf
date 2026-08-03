resource "routeros_interface_bridge_calea" "calea" {
  action = "sniff"
  chain  = "forward"
}
