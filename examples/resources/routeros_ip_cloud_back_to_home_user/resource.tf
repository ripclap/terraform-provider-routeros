resource "routeros_ip_cloud_back_to_home_user" "user" {
  name      = "example"
  allow_lan = true
  comment   = "Managed by OpenTofu"
}
