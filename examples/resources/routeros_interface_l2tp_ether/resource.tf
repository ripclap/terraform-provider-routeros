resource "routeros_interface_l2tp_ether" "ether" {
  name            = "example"
  allow_fast_path = true
  circuit_id      = "example"
}
