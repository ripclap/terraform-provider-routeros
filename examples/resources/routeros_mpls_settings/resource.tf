resource "routeros_mpls_settings" "settings" {
  allow_fast_path     = true
  dynamic_label_range = "1"
  propagate_ttl       = true
}
