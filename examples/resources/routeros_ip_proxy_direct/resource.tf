resource "routeros_ip_proxy_direct" "direct" {
  action   = "allow"
  comment  = "Managed by OpenTofu"
  disabled = true
}
