resource "routeros_routing_filter_community_ext_list" "list" {
  list     = "example"
  comment  = "Managed by OpenTofu"
  disabled = true
}
