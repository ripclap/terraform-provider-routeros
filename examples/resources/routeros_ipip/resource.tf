resource "routeros_ipip" "ipip" {
  name            = "example"
  allow_fast_path = true
  clamp_tcp_mss   = true
}
