resource "routeros_mpls_ldp_advertise_filter" "filter" {
  advertise = true
  comment   = "Managed by OpenTofu"
  disabled  = true
}
