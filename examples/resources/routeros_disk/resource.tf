resource "routeros_disk" "disk" {
  comment         = "Managed by OpenTofu"
  compress        = true
  crypted_backend = "example"
}
