resource "routeros_mpls_mangle" "mangle" {
  chain    = "forward"
  comment  = "Managed by OpenTofu"
  disabled = true
}
