resource "routeros_ip_socks_access" "access" {
  action   = "allow"
  comment  = "Managed by OpenTofu"
  disabled = true
}
