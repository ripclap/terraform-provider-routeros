resource "routeros_ip_proxy_cache" "cache" {
  action   = "allow"
  comment  = "Managed by OpenTofu"
  disabled = true
}
