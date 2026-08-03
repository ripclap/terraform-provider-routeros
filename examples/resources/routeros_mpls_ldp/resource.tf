resource "routeros_mpls_ldp" "ldp" {
  comment                = "Managed by OpenTofu"
  disabled               = true
  distribute_for_default = true
}
