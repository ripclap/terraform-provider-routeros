resource "routeros_ip_smb_shares" "shares" {
  name      = "example"
  comment   = "Managed by OpenTofu"
  directory = "example"
}
