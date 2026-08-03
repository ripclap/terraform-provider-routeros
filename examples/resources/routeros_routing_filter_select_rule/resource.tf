resource "routeros_routing_filter_select_rule" "rule" {
  chain    = "forward"
  comment  = "Managed by OpenTofu"
  disabled = true
}
