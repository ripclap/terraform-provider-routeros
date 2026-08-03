resource "routeros_ip_proxy_access" "access" {
  action      = "allow"
  action_data = "example"
  comment     = "Managed by OpenTofu"
}
