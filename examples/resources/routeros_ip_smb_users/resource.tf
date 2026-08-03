resource "routeros_ip_smb_users" "users" {
  name     = "example"
  comment  = "Managed by OpenTofu"
  disabled = true
}
