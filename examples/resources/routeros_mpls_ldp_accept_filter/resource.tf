resource "routeros_mpls_ldp_accept_filter" "filter" {
  accept   = true
  comment  = "Managed by OpenTofu"
  disabled = true
}
