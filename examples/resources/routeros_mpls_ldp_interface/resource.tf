resource "routeros_mpls_ldp_interface" "interface" {
  interface                = "ether1"
  accept_dynamic_neighbors = true
  comment                  = "Managed by OpenTofu"
}
