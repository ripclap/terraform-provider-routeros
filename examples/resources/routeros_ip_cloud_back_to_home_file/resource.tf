resource "routeros_ip_cloud_back_to_home_file" "file" {
  path          = "example"
  allow_uploads = true
  comment       = "Managed by OpenTofu"
}
