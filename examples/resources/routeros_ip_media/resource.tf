resource "routeros_ip_media" "media" {
  path             = "example"
  allowed_hostname = "example"
  allowed_ip       = "192.0.2.1"
}
