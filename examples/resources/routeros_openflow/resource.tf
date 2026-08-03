resource "routeros_openflow" "openflow" {
  name        = "example"
  certificate = "example"
  comment     = "Managed by OpenTofu"
}
