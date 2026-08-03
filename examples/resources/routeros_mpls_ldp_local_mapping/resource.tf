resource "routeros_mpls_ldp_local_mapping" "mapping" {
  comment     = "Managed by OpenTofu"
  disabled    = true
  dst_address = "192.0.2.1"
}
