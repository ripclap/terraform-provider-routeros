resource "routeros_routing_pimsm_bsr_candidate" "candidate" {
  instance = "example"
  address  = "192.0.2.1"
  disabled = true
}
