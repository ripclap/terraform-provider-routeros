resource "routeros_routing_pimsm_instance" "instance" {
  name             = "example"
  afi              = "ip"
  bsm_forward_back = true
}
