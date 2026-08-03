resource "routeros_mpls_ldp_neighbor" "neighbor" {
  comment       = "Managed by OpenTofu"
  disabled      = true
  send_targeted = true
}
